package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTConfig define la interfaz para la configuración JWT
// Esto evita el import cycle
type JWTConfig interface {
	GetJWTSecret() string
	GetJWTExpirationMinutes() string
}

// Claims personalizados para el token JWT de acceso
// Incluye ID y Email
type CustomClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
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
func GenerateJWT(userID int, email string, cfg JWTConfig) (string, error) {
	expMin, err := strconv.Atoi(cfg.GetJWTExpirationMinutes())
	if err != nil {
		expMin = 60 // fallback
	}
	claims := CustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expMin) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
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
func GenerateRefreshToken(userID int, cfg JWTConfig) (string, error) {
	days := 7
	claims := RefreshClaims{
		UserID: userID,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * time.Duration(days))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
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
