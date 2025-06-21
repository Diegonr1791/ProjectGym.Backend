package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware es un middleware que valida el token JWT
// Extrae el token del header Authorization: Bearer <token>
// Si es válido, agrega los claims al contexto
func JWTAuthMiddleware(cfg JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extraer el token del header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorización requerido"})
			c.Abort()
			return
		}

		// Verificar que el header tenga el formato correcto: "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de autorización inválido. Use: Bearer <token>"})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		// Validar el token
		claims, err := ValidateJWT(tokenString, cfg)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido: " + err.Error()})
			c.Abort()
			return
		}

		// Agregar los claims al contexto para que los handlers puedan acceder
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("claims", claims)

		c.Next()
	}
}

// GetUserIDFromContext extrae el userID del contexto de Gin
// Útil para los handlers que necesitan el ID del usuario autenticado
func GetUserIDFromContext(c *gin.Context) (int, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}
	if id, ok := userID.(int); ok {
		return id, true
	}
	return 0, false
}

// GetUserEmailFromContext extrae el email del contexto de Gin
func GetUserEmailFromContext(c *gin.Context) (string, bool) {
	email, exists := c.Get("user_email")
	if !exists {
		return "", false
	}
	if em, ok := email.(string); ok {
		return em, true
	}
	return "", false
}
