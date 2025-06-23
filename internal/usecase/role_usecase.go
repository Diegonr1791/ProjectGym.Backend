// internal/usecase/role_usecase.go
package usecase

import (
	"context"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
)

// RoleUseCase defines the business logic for role operations
type RoleUseCase struct {
	roleRepo repositories.RoleRepository
}

// NewRoleUseCase creates a new role use case
func NewRoleUseCase(roleRepo repositories.RoleRepository) *RoleUseCase {
	return &RoleUseCase{
		roleRepo: roleRepo,
	}
}

// CreateRole creates a new role
func (uc *RoleUseCase) CreateRole(ctx context.Context, role *models.Role) error {
	// Validate role name
	if role.Name == "" {
		return domainErrors.ErrBadRequest
	}

	// Check if role name already exists
	existingRole, err := uc.roleRepo.GetByName(ctx, role.Name)
	if err == nil && existingRole != nil && !existingRole.IsSoftDeleted() {
		return domainErrors.ErrConflict
	}

	return uc.roleRepo.Create(ctx, role)
}

// GetRoleByID retrieves a role by its ID
func (uc *RoleUseCase) GetRoleByID(ctx context.Context, id uint) (*models.Role, error) {
	role, err := uc.roleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if role.IsSoftDeleted() {
		return nil, domainErrors.ErrNotFound
	}

	return role, nil
}

// GetRoleByName retrieves a role by its name
func (uc *RoleUseCase) GetRoleByName(ctx context.Context, name string) (*models.Role, error) {
	role, err := uc.roleRepo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	if role.IsSoftDeleted() {
		return nil, domainErrors.ErrNotFound
	}

	return role, nil
}

// GetAllRoles retrieves all active roles
func (uc *RoleUseCase) GetAllRoles(ctx context.Context) ([]models.Role, error) {
	return uc.roleRepo.GetAll(ctx)
}

// GetAllRolesWithDeleted retrieves all roles including deleted ones
func (uc *RoleUseCase) GetAllRolesWithDeleted(ctx context.Context) ([]models.Role, error) {
	return uc.roleRepo.GetAllWithDeleted(ctx)
}

// UpdateRole updates an existing role
func (uc *RoleUseCase) UpdateRole(ctx context.Context, role *models.Role) error {
	// Validate role exists and is not deleted
	existingRole, err := uc.roleRepo.GetByID(ctx, role.ID)
	if err != nil {
		return err
	}

	if existingRole.IsSoftDeleted() {
		return domainErrors.ErrNotFound
	}

	// System roles cannot be modified
	if existingRole.IsSystem {
		return domainErrors.ErrSystemRoleNotDeletable
	}

	// Check if new name conflicts with existing role
	if role.Name != existingRole.Name {
		conflictingRole, err := uc.roleRepo.GetByName(ctx, role.Name)
		if err == nil && conflictingRole != nil && !conflictingRole.IsSoftDeleted() {
			return domainErrors.ErrConflict
		}
	}

	return uc.roleRepo.Update(ctx, role)
}

// SoftDeleteRole marks a role as deleted
func (uc *RoleUseCase) SoftDeleteRole(ctx context.Context, id uint) error {
	role, err := uc.roleRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if role.IsSoftDeleted() {
		return domainErrors.ErrNotFound
	}

	// System roles cannot be deleted
	if role.IsSystem {
		return domainErrors.ErrSystemRoleNotDeletable
	}

	return uc.roleRepo.SoftDelete(ctx, id)
}

// HardDeleteRole permanently deletes a role
func (uc *RoleUseCase) HardDeleteRole(ctx context.Context, id uint) error {
	role, err := uc.roleRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// System roles cannot be deleted
	if role.IsSystem {
		return domainErrors.ErrSystemRoleNotDeletable
	}

	return uc.roleRepo.HardDelete(ctx, id)
}

// RestoreRole restores a deleted role
func (uc *RoleUseCase) RestoreRole(ctx context.Context, id uint) error {
	role, err := uc.roleRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if !role.IsSoftDeleted() {
		return domainErrors.ErrBadRequest
	}

	return uc.roleRepo.Restore(ctx, id)
}

// GetSystemRoles retrieves all system roles
func (uc *RoleUseCase) GetSystemRoles(ctx context.Context) ([]models.Role, error) {
	return uc.roleRepo.GetSystemRoles(ctx)
}

// GetActiveRoles retrieves all active roles
func (uc *RoleUseCase) GetActiveRoles(ctx context.Context) ([]models.Role, error) {
	return uc.roleRepo.GetActiveRoles(ctx)
}
