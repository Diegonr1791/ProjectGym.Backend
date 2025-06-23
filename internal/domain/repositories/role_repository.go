// internal/domain/repositories/role_repository.go
package repository

import (
	"context"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
)

// RoleRepository defines the interface for role data access
type RoleRepository interface {
	// Create creates a new role
	Create(ctx context.Context, role *model.Role) error

	// GetByID retrieves a role by its ID
	GetByID(ctx context.Context, id uint) (*model.Role, error)

	// GetByName retrieves a role by its name
	GetByName(ctx context.Context, name string) (*model.Role, error)

	// GetAll retrieves all active roles
	GetAll(ctx context.Context) ([]model.Role, error)

	// GetAllWithDeleted retrieves all roles including deleted ones
	GetAllWithDeleted(ctx context.Context) ([]model.Role, error)

	// Update updates an existing role
	Update(ctx context.Context, role *model.Role) error

	// SoftDelete marks a role as deleted
	SoftDelete(ctx context.Context, id uint) error

	// HardDelete permanently deletes a role
	HardDelete(ctx context.Context, id uint) error

	// Restore restores a deleted role
	Restore(ctx context.Context, id uint) error

	// GetSystemRoles retrieves all system roles
	GetSystemRoles(ctx context.Context) ([]model.Role, error)

	// GetActiveRoles retrieves all active roles
	GetActiveRoles(ctx context.Context) ([]model.Role, error)
}
