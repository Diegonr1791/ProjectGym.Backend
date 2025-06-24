package adapters

import (
	"context"

	auth "github.com/Diegonr1791/GymBro/internal/auth"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
)

// RoleAdapter adapta el repositorio de roles del dominio para el middleware de autorización
type RoleAdapter struct {
	roleRepo repositories.RoleRepository
}

// NewRoleAdapter crea un nuevo adaptador de roles
func NewRoleAdapter(roleRepo repositories.RoleRepository) auth.RoleRepository {
	return &RoleAdapter{
		roleRepo: roleRepo,
	}
}

// GetByID obtiene un rol por ID y lo convierte al formato del middleware
func (ra *RoleAdapter) GetByID(id uint) (*auth.Role, error) {
	role, err := ra.roleRepo.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &auth.Role{
		ID:   role.ID,
		Name: role.Name,
	}, nil
}
