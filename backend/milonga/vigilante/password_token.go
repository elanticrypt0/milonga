package vigilante

import (
	"crypto/aes"
	"crypto/cipher"
	cryptoRand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	mathRand "math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PasswordToken struct {
	ID         uuid.UUID `gorm:"type:varchar(50);primary_key;"`
	UserAuthID uuid.UUID `gorm:"type:varchar(50);not null"`
	Token      string    `gorm:"unique;not null"`
	IsUsed     bool      `gorm:"default:false"`
	TotalUses  int       `gorm:"default:0"`
	MaxUses    int       `gorm:"default:1"`
	ExpiresAt  time.Time `gorm:"not null"`
	User       UserAuth  `gorm:"foreignKey:UserAuthID"`
	gorm.Model
}

const (
	DefaultTokenValidity = 48 * time.Hour
	DefaultTokenLength   = 6
	TOKEN_NOT_VALID      = "token is not valid"
)

func NewPasswordToken() *PasswordToken {
	return &PasswordToken{}
}

// BeforeCreate será llamado por GORM antes de crear un nuevo usuario
func (me *PasswordToken) BeforeCreate(tx *gorm.DB) error {
	me.ID = uuid.New()
	return nil
}

func (me *PasswordToken) updateOTPUses() {
	me.TotalUses++
	me.IsUsed = me.checkIsUsed()
}

func (me *PasswordToken) Create(userID uuid.UUID, tx *gorm.DB) (string, error) {
	return me.CreateWithValidity(userID, DefaultTokenValidity, tx)
}

func (me *PasswordToken) DenyOTPOfUser(tx *gorm.DB, userID uuid.UUID) error {
	// si el usuario tiene otros token creados pero no usados
	err := tx.Model(&PasswordToken{}).Where("user_auth_id = ?", userID).Update("is_used", 1).Error
	if err != nil {
		return fmt.Errorf("error denying other OTP of user: %v", err)
	}
	return nil
}

// update max uses

func (me *PasswordToken) UpdateMaxUses(tx *gorm.DB, userID uuid.UUID, maxUses int) error {

	passmodel := &PasswordToken{}
	passmodel.UserAuthID = userID
	err := tx.Model(&PasswordToken{}).Where("user_auth_id = ?", userID).Last(&passmodel).Error
	if err != nil {
		return fmt.Errorf("error retriving user by id: %v", err)
	}
	// cambia el máximo de usos
	passmodel.MaxUses = maxUses
	// actualiza si habilita el logueo
	passmodel.IsUsed = passmodel.checkIsUsed()
	err = tx.Save(&passmodel).Error

	if err != nil {
		return fmt.Errorf("error updating max uses: %s", err)
	}

	return nil
}

func (me *PasswordToken) CreateWithValidity(userID uuid.UUID, validity time.Duration, tx *gorm.DB) (string, error) {
	plainToken := me.generateToken()
	encryptedToken, err := me.encryptToken(plainToken)
	if err != nil {
		return "", fmt.Errorf("error encrypting token: %w", err)
	}

	newPassToken := &PasswordToken{
		UserAuthID: userID,
		Token:      encryptedToken,
		ExpiresAt:  time.Now().Add(validity),
	}

	err = tx.Save(newPassToken).Error
	if err != nil {
		return "", fmt.Errorf("error creating password token: %w", err)
	}

	return plainToken, nil
}

func (me *PasswordToken) UpdateToken(tx *gorm.DB) (string, error) {
	return me.UpdateTokenWithValidity(DefaultTokenValidity, tx)
}

func (me *PasswordToken) UpdateTokenWithValidity(validity time.Duration, tx *gorm.DB) (string, error) {
	plainToken := me.generateToken()
	encryptedToken, err := me.encryptToken(plainToken)
	if err != nil {
		return "", fmt.Errorf("error encrypting token: %w", err)
	}

	sameToken := &PasswordToken{
		UserAuthID: me.UserAuthID,
		Token:      encryptedToken,
		IsUsed:     false,
		ExpiresAt:  time.Now().Add(validity),
	}

	if err := tx.Save(sameToken).Error; err != nil {
		return "", fmt.Errorf("error updating token: %w", err)
	}

	return plainToken, nil // Retornamos el token sin encriptar para el usuario
}

func (me *PasswordToken) RefreshToken(userID uuid.UUID, token string, tx *gorm.DB) (string, error) {
	return me.RefreshTokenWithValidity(userID, token, DefaultTokenValidity, tx)
}

func (me *PasswordToken) RefreshTokenWithValidity(userID uuid.UUID, token string, validity time.Duration, tx *gorm.DB) (string, error) {
	var instance PasswordToken
	if err := tx.First(&instance, "user_auth_id = ? and is_used=?", userID, false).Error; err != nil {
		return "", fmt.Errorf("user's token not found")
	}

	decrypted_token, err := me._decryptToken(instance.Token)
	if err != nil {
		return "", fmt.Errorf("token not found or already used")
	}

	if token != decrypted_token {
		return "", fmt.Errorf(TOKEN_NOT_VALID)
	}

	if time.Now().After(instance.ExpiresAt) {
		return "", fmt.Errorf("token has expired")
	}

	// Marcar el token actual como usado
	instance.updateOTPUses()
	if err := tx.Save(&instance).Error; err != nil {
		return "", fmt.Errorf("error marking token as used: %w", err)
	}

	// Generar y encriptar nuevo token
	newPlainToken := me.generateToken()
	newEncryptedToken, err := me.encryptToken(newPlainToken)
	if err != nil {
		return "", fmt.Errorf("error encrypting new token: %w", err)
	}

	newToken := &PasswordToken{
		UserAuthID: instance.UserAuthID,
		Token:      newEncryptedToken,
		ExpiresAt:  time.Now().Add(validity),
	}

	if err := tx.Save(newToken).Error; err != nil {
		return "", fmt.Errorf("error creating new token: %w", err)
	}

	return newPlainToken, nil
}

func (me *PasswordToken) CheckToken(userID uuid.UUID, token string, tx *gorm.DB) error {

	var instance PasswordToken
	if err := tx.First(&instance, "user_auth_id = ? and is_used=?", userID, false).Error; err != nil {
		return fmt.Errorf("user's token not found")
	}

	decrypted_token, err := me._decryptToken(instance.Token)
	if err != nil {
		return fmt.Errorf("token not found or already used")
	}

	// si el token es corto: es decir ABC-DFG
	if len(token) == 7 {
		if token != decrypted_token {
			return fmt.Errorf(TOKEN_NOT_VALID)
		}
	} else {
		// si el token es largo (es decir sin descifrar)
		inputToken, err := me._decryptToken(token)
		if err != nil {
			return fmt.Errorf(TOKEN_NOT_VALID)
		}

		if inputToken != decrypted_token {
			return fmt.Errorf(TOKEN_NOT_VALID)
		}
	}

	if time.Now().After(instance.ExpiresAt) {
		return fmt.Errorf("token has expired")
	}

	// actualiza si el token está usado
	instance.updateOTPUses()
	if err := tx.Save(&instance).Error; err != nil {
		return fmt.Errorf("error marking token as used: %w", err)
	}

	return nil
}

func (me *PasswordToken) checkIsUsed() bool {
	return me.TotalUses >= me.MaxUses
}

func (me *PasswordToken) generateToken() string {
	r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	const chars = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"

	tokenLength := DefaultTokenLength
	var token strings.Builder
	token.Grow(tokenLength)

	for i := 0; i < tokenLength; i++ {
		if i > 0 && i%3 == 0 {
			token.WriteByte('-')
		}
		token.WriteByte(chars[r.Intn(len(chars))])
	}

	return token.String()
}

// Función para encriptar el token usando AES-256-GCM
func (me *PasswordToken) encryptToken(token string) (string, error) {
	encryptionKey, err := GetPasswordTokenEncryptionKey()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	block, err := aes.NewCipher([]byte(encryptionKey))
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(cryptoRand.Reader, nonce); err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(token), nil)

	// Combinar nonce y ciphertext para almacenamiento
	encrypted := append(nonce, ciphertext...)

	// Codificar en base64 para almacenamiento seguro en BD
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// Función para desencriptar el token (si es necesario para debugging)
func (me *PasswordToken) _decryptToken(encryptedToken string) (string, error) {
	// Decodificar base64
	encrypted, err := base64.StdEncoding.DecodeString(encryptedToken)
	if err != nil {
		return "", err
	}

	encryptionKey, _ := GetPasswordTokenEncryptionKey()

	block, err := aes.NewCipher([]byte(encryptionKey))
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(encrypted) < 12 {
		return "", fmt.Errorf("encrypted token too short")
	}

	nonce := encrypted[:12]
	ciphertext := encrypted[12:]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
