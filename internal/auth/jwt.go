package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTConfig define la interfaz para la configuración JWT
// Esto evita el import cycle
type JWTConfig interface {
	GetJWTSecret() string
	GetJWTExpirationMinutes() int
	GetRefreshExpirationHours() int
	GetRefreshMaxAge() int // MaxAge en segundos para la cookie
}

// Claims personalizados para el token JWT de acceso
// Incluye ID, Email y RoleID
type CustomClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	RoleID uint   `json:"role_id"`
	jwt.RegisteredClaims
}

// Claims para el refresh token
// Solo incluye el userID y un tipo
type RefreshClaims struct {
	UserID int    `json:"user_id"`
	Type   string `json:"type"` // siempre "refresh"
	jwt.RegisteredClaims
}

// GenerateJWT genera un token JWT de acceso para un usuario dado
func GenerateJWT(userID uint, email string, roleID uint, cfg JWTConfig) (string, error) {
	expirationTime := time.Now().Add(time.Duration(cfg.GetJWTExpirationMinutes()) * time.Minute)
	claims := CustomClaims{
		UserID: int(userID),
		Email:  email,
		RoleID: roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.GetJWTSecret()))
}

// ValidateJWT valida un token JWT de acceso y retorna los claims si es válido
func ValidateJWT(tokenString string, cfg JWTConfig) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		// Validar el método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma inválido")
		}
		return []byte(cfg.GetJWTSecret()), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token inválido")
	}
	return claims, nil
}

// GenerateRefreshToken genera un refresh token para un usuario dado
// Expira en 7 días por defecto
func GenerateRefreshToken(userID uint, cfg JWTConfig) (string, error) {
	expirationTime := time.Now().Add(time.Duration(cfg.GetRefreshExpirationHours()) * time.Hour)
	claims := RefreshClaims{
		UserID: int(userID),
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.GetJWTSecret()))
}

// ValidateRefreshToken valida un refresh token y retorna los claims si es válido
func ValidateRefreshToken(tokenString string, cfg JWTConfig) (*RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma inválido")
		}
		return []byte(cfg.GetJWTSecret()), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*RefreshClaims)
	if !ok || !token.Valid {
		return nil, errors.New("refresh token inválido")
	}
	if claims.Type != "refresh" {
		return nil, errors.New("el token no es de tipo refresh")
	}
	return claims, nil
}
