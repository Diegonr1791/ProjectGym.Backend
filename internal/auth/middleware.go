package auth

import (
	"net/http"
	"slices"
	"strings"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	"github.com/gin-gonic/gin"
)

// RoleRepository define la interfaz para el repositorio de roles
// Esto evita el import cycle
type RoleRepository interface {
	GetByID(id uint) (*Role, error)
}

// Role representa un rol simplificado para el middleware
type Role struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// JWTAuthMiddleware es un middleware que valida el token JWT
// Extrae el token del header Authorization: Bearer <token>
// Si es válido, agrega los claims al contexto
func JWTAuthMiddleware(cfg JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extraer el token del header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(domainErrors.ErrUnauthorized)
			c.Abort()
			return
		}

		// Verificar que el header tenga el formato correcto: "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.Error(domainErrors.NewAppError(http.StatusUnauthorized, "INVALID_AUTH_FORMAT", "Formato de autorización inválido. Use: Bearer <token>", nil))
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		// Validar el token
		claims, err := ValidateJWT(tokenString, cfg)
		if err != nil {
			c.Error(domainErrors.NewAppError(http.StatusUnauthorized, "INVALID_TOKEN", "Token inválido: "+err.Error(), err))
			c.Abort()
			return
		}

		// Agregar los claims al contexto para que los handlers puedan acceder
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("role_id", claims.RoleID)
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

// GetRoleIDFromContext extrae el roleID del contexto de Gin
func GetRoleIDFromContext(c *gin.Context) (uint, bool) {
	roleID, exists := c.Get("role_id")
	if !exists {
		return 0, false
	}
	if rid, ok := roleID.(uint); ok {
		return rid, true
	}
	return 0, false
}

// RequireRoleMiddleware es un middleware que verifica que el usuario tenga uno de los roles especificados
func RequireRoleMiddleware(roleRepo RoleRepository, allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el roleID del contexto
		roleID, exists := GetRoleIDFromContext(c)
		if !exists {
			c.Error(domainErrors.NewAppError(http.StatusForbidden, "ROLE_INFO_UNAVAILABLE", "Role information not available", nil))
			c.Abort()
			return
		}

		// Obtener el rol desde la base de datos
		role, err := roleRepo.GetByID(roleID)
		if err != nil {
			c.Error(domainErrors.NewAppError(http.StatusForbidden, "INVALID_ROLE_INFO", "Invalid role information", err))
			c.Abort()
			return
		}

		// Verificar si el usuario tiene uno de los roles permitidos
		if !slices.Contains(allowedRoles, role.Name) {
			c.Error(domainErrors.NewAppError(http.StatusForbidden, "INSUFFICIENT_PERMISSIONS", "Insufficient permissions. Only admin and dev roles can perform this action.", nil))
			c.Abort()
			return
		}

		c.Next()
	}
}
