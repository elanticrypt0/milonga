# Password Token System Documentation

## Overview
The Password Token system provides a secure, temporary access mechanism for users. It generates single-use tokens that expire after a configurable time period. These tokens are stored encrypted in the database and can be used for one-time authentication.

## Security Features
- AES-256-GCM encryption for token storage
- Single-use tokens (invalidated after use)
- Configurable expiration time
- User-friendly token format
- Cryptographically secure random generation
- Protection against token reuse

## Technical Specifications

### Token Format
- Length: 6 characters by default (configurable via `DefaultTokenLength`)
- Character Set: Alphanumeric (excluding ambiguous characters I, O, 0, 1)
- Format Example: `ABC-123`

### Encryption
- Algorithm: AES-256-GCM
- Key Size: 32 bytes
- Nonce: 12 bytes (randomly generated per token)
- Storage: Base64 encoded

### Default Settings
- Token Validity: 48 hours
- Token Length: 6 characters
- Auto-expiration: Yes

## API Reference

### Create New Token
```go
func (me *PasswordToken) Create(userID uuid.UUID, tx *gorm.DB) error
```
Creates a new token with default validity period (48 hours).

### Create Token with Custom Validity
```go
func (me *PasswordToken) CreateWithValidity(userID uuid.UUID, validity time.Duration, tx *gorm.DB) error
```
Creates a new token with custom validity period.

### Refresh Token
```go
func (me *PasswordToken) RefreshToken(token string, tx *gorm.DB) (string, error)
```
Invalidates existing token and generates a new one with default validity.

### Check Token
```go
func (me *PasswordToken) CheckToken(userID uuid.UUID, token string, tx *gorm.DB) error
```
Validates and consumes a token.

## Database Schema
```sql
CREATE TABLE password_tokens (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    token VARCHAR NOT NULL UNIQUE,
    is_used BOOLEAN DEFAULT FALSE,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

## Usage Examples

### Creating a New Token
```go
token := NewPasswordToken()
err := token.Create(userID, db)
```

### Creating a Token with Custom Validity
```go
token := NewPasswordToken()
err := token.CreateWithValidity(userID, 24*time.Hour, db)
```

### Validating a Token
```go
token := NewPasswordToken()
err := token.CheckToken(userID, receivedToken, db)
if err != nil {
    // Token is invalid, expired, or already used
}
```

## Security Considerations

### Production Deployment
1. Store encryption key in environment variables
2. Consider using a Key Management Service (KMS)
3. Implement rate limiting for token validation attempts
4. Monitor and log failed validation attempts
5. Regular token cleanup for expired entries

### Best Practices
1. Never store unencrypted tokens in logs
2. Implement token cleanup routine
3. Monitor token usage patterns for abuse
4. Set appropriate token validity periods based on use case
5. Implement proper error handling and logging

## Error Handling

The system provides specific errors for different scenarios:
- Token not found
- Token already used
- Token expired
- Encryption/decryption errors

## Performance Considerations

1. Database Indexes
```sql
CREATE INDEX idx_password_tokens_token ON password_tokens(token);
CREATE INDEX idx_password_tokens_user_id ON password_tokens(user_id);
```

2. Token Cleanup
Implement regular cleanup of expired tokens:
```sql
DELETE FROM password_tokens WHERE expires_at < NOW() OR is_used = true;
```

## Integration with Authentication System

The password token system integrates with the main authentication system through:
1. User relationship (foreign key)
2. Token validation in login process
3. Automatic token invalidation after use

For implementation details, see `LoginByPasswordToken` in the authentication handler.