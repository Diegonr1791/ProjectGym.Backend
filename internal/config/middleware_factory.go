package config

import (
	adapters "github.com/Diegonr1791/GymBro/internal/adapters"
	auth "github.com/Diegonr1791/GymBro/internal/auth"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	"github.com/gin-gonic/gin"
)

// MiddlewareFactory maneja la creación de middlewares
type MiddlewareFactory struct {
	container *Container
}

// NewMiddlewareFactory crea una nueva instancia del factory de middlewares
func NewMiddlewareFactory(container *Container) *MiddlewareFactory {
	return &MiddlewareFactory{
		container: container,
	}
}

// CreateUserDeletionAuthMiddleware crea el middleware de autorización para eliminación de usuarios
func (mf *MiddlewareFactory) CreateUserDeletionAuthMiddleware() gin.HandlerFunc {
	roleAdapter := adapters.NewRoleAdapter(mf.container.RoleRepo)
	allowedRoles := []string{models.RoleAdmin, models.RoleDev}
	return auth.RequireRoleMiddleware(roleAdapter, allowedRoles...)
}

// CreateJWTAuthMiddleware crea el middleware de autenticación JWT
func (mf *MiddlewareFactory) CreateJWTAuthMiddleware() gin.HandlerFunc {
	return auth.JWTAuthMiddleware(mf.container.JWTConfig)
}
