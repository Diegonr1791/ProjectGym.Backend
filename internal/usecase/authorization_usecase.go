package usecase

import (
	"context"

	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
)

// AuthorizationUsecase maneja la lógica de autorización
type AuthorizationUsecase struct {
	roleRepo repositories.RoleRepository
}

// NewAuthorizationUsecase crea una nueva instancia del usecase de autorización
func NewAuthorizationUsecase(roleRepo repositories.RoleRepository) *AuthorizationUsecase {
	return &AuthorizationUsecase{
		roleRepo: roleRepo,
	}
}

// CanDeleteUsers verifica si un usuario puede eliminar otros usuarios
func (uc *AuthorizationUsecase) CanDeleteUsers(ctx context.Context, userRoleID uint) (bool, error) {
	role, err := uc.roleRepo.GetByID(ctx, userRoleID)
	if err != nil {
		return false, err
	}

	// Solo admin y dev pueden eliminar usuarios
	allowedRoles := []string{models.RoleAdmin, models.RoleDev}

	for _, allowedRole := range allowedRoles {
		if role.Name == allowedRole {
			return true, nil
		}
	}

	return false, nil
}

// GetAllowedRolesForUserDeletion retorna los roles que pueden eliminar usuarios
func (uc *AuthorizationUsecase) GetAllowedRolesForUserDeletion() []string {
	return []string{models.RoleAdmin, models.RoleDev}
}
