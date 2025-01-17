package vigilante

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoginMethod string

const (
	LoginMethodPassword LoginMethod = "password"
	LoginMethodToken    LoginMethod = "token"
	LoginMethodOAuth    LoginMethod = "oauth"
	LoginStatusSuccess  string      = "success"
	LoginStatusFailed   string      = "failed"
	LoginStatusBlocked  string      = "blocked"
)

type LoginAudit struct {
	ID            uuid.UUID `gorm:"type:varchar(50);primary_key;"`
	UserID        uuid.UUID `gorm:"type:varchar(50);not null"`
	LoginTime     time.Time `gorm:"not null"`
	LogoutTime    *time.Time
	IPAddress     string      `gorm:"size:45;not null"`  // IPv6 length
	UserAgent     string      `gorm:"size:255"`          // Navegador/App info
	Device        string      `gorm:"size:100"`          // Dispositivo usado
	Location      string      `gorm:"size:100"`          // Ubicación basada en IP
	LoginMethod   LoginMethod `gorm:"size:20;not null"`  // Método de autenticación
	Status        string      `gorm:"size:20;not null"`  // Éxito/Fallo/Bloqueado
	FailureReason string      `gorm:"size:255"`          // Razón si falló
	SessionID     string      `gorm:"size:100"`          // ID de sesión
	Country       string      `gorm:"size:2"`            // Código ISO del país
	City          string      `gorm:"size:100"`          // Ciudad basada en IP
	Success       bool        `gorm:"not null"`          // Indicador de éxito
	User          User        `gorm:"foreignKey:UserID"` // Relación con Usuario
	gorm.Model
}

func NewLoginAudit() *LoginAudit {
	return &LoginAudit{
		ID:        uuid.New(),
		LoginTime: time.Now(),
	}
}

func (me *LoginAudit) BeforeCreate(tx *gorm.DB) error {
	me.ID = uuid.New()
	return nil
}

func (me *LoginAudit) Create(tx *gorm.DB) error {
	return tx.Create(me).Error
}

// RegisterSuccessfulLogin registra un login exitoso
func (me *LoginAudit) RegisterSuccessfulLogin(userID uuid.UUID, ipAddress, userAgent string, method LoginMethod, tx *gorm.DB) error {
	me.UserID = userID
	me.IPAddress = ipAddress
	me.UserAgent = userAgent
	me.LoginMethod = method
	me.Status = LoginStatusSuccess
	me.Success = true

	// Aquí podrías agregar lógica para obtener:
	// - Ubicación basada en IP
	// - Información del dispositivo desde User-Agent
	// - País y ciudad usando un servicio de geolocalización

	return me.Create(tx)
}

// RegisterFailedLogin registra un intento fallido
func (me *LoginAudit) RegisterFailedLogin(userID uuid.UUID, ipAddress, userAgent string, method LoginMethod, reason string, tx *gorm.DB) error {
	me.UserID = userID
	me.IPAddress = ipAddress
	me.UserAgent = userAgent
	me.LoginMethod = method
	me.Status = LoginStatusFailed
	me.FailureReason = reason
	me.Success = false

	return me.Create(tx)
}

// RegisterLogout registra el cierre de sesión
func (me *LoginAudit) RegisterLogout(sessionID string, tx *gorm.DB) error {
	now := time.Now()
	return tx.Model(me).
		Where("session_id = ? AND logout_time IS NULL", sessionID).
		Update("logout_time", now).Error
}

// GetUserLastLogins obtiene los últimos logins de un usuario
func (me *LoginAudit) GetUserLastLogins(userID uuid.UUID, limit int, tx *gorm.DB) ([]LoginAudit, error) {
	var logins []LoginAudit
	err := tx.Where("user_id = ?", userID).
		Order("login_time DESC").
		Limit(limit).
		Find(&logins).Error
	return logins, err
}

// GetFailedLoginAttempts obtiene los intentos fallidos recientes
func (me *LoginAudit) GetFailedLoginAttempts(userID uuid.UUID, duration time.Duration, tx *gorm.DB) (int64, error) {
	var count int64
	err := tx.Model(&LoginAudit{}).
		Where("user_id = ? AND success = ? AND login_time > ?",
			userID, false, time.Now().Add(-duration)).
		Count(&count).Error
	return count, err
}

// GetActiveSessionsCount obtiene el número de sesiones activas
func (me *LoginAudit) GetActiveSessionsCount(userID uuid.UUID, tx *gorm.DB) (int64, error) {
	var count int64
	err := tx.Model(&LoginAudit{}).
		Where("user_id = ? AND logout_time IS NULL", userID).
		Count(&count).Error
	return count, err
}
