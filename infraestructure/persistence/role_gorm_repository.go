// infraestructure/persistence/role_gorm_repository.go
package persistence

import (
	"context"

	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// RoleGormRepository implements RoleRepository using GORM
type RoleGormRepository struct {
	db *gorm.DB
}

// NewRoleGormRepository creates a new role GORM repository
func NewRoleGormRepository(db *gorm.DB) repositories.RoleRepository {
	return &RoleGormRepository{
		db: db,
	}
}

// Create creates a new role
func (r *RoleGormRepository) Create(ctx context.Context, role *models.Role) error {
	if err := r.db.WithContext(ctx).Create(role).Error; err != nil {
		return errors.Wrap(err, "RoleGormRepository.Create")
	}
	return nil
}

// GetByID retrieves a role by its ID
func (r *RoleGormRepository) GetByID(ctx context.Context, id uint) (*models.Role, error) {
	var role models.Role
	if err := r.db.WithContext(ctx).First(&role, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(err, "RoleGormRepository.GetByID: id %d", id)
		}
		return nil, errors.Wrapf(err, "RoleGormRepository.GetByID: id %d", id)
	}
	return &role, nil
}

// GetByName retrieves a role by its name
func (r *RoleGormRepository) GetByName(ctx context.Context, name string) (*models.Role, error) {
	var role models.Role
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(err, "RoleGormRepository.GetByName: name %s", name)
		}
		return nil, errors.Wrapf(err, "RoleGormRepository.GetByName: name %s", name)
	}
	return &role, nil
}

// GetAll retrieves all active roles
func (r *RoleGormRepository) GetAll(ctx context.Context) ([]models.Role, error) {
	var roles []models.Role
	if err := r.db.WithContext(ctx).Where("is_deleted = ?", false).Find(&roles).Error; err != nil {
		return nil, errors.Wrap(err, "RoleGormRepository.GetAll")
	}
	return roles, nil
}

// GetAllWithDeleted retrieves all roles including deleted ones
func (r *RoleGormRepository) GetAllWithDeleted(ctx context.Context) ([]models.Role, error) {
	var roles []models.Role
	if err := r.db.WithContext(ctx).Find(&roles).Error; err != nil {
		return nil, errors.Wrap(err, "RoleGormRepository.GetAllWithDeleted")
	}
	return roles, nil
}

// Update updates an existing role
func (r *RoleGormRepository) Update(ctx context.Context, role *models.Role) error {
	if err := r.db.WithContext(ctx).Save(role).Error; err != nil {
		return errors.Wrapf(err, "RoleGormRepository.Update: id %d", role.ID)
	}
	return nil
}

// SoftDelete marks a role as deleted
func (r *RoleGormRepository) SoftDelete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Model(&models.Role{}).Where("id = ?", id).Updates(map[string]interface{}{
		"is_deleted": true,
		"is_active":  false,
	}).Error; err != nil {
		return errors.Wrapf(err, "RoleGormRepository.SoftDelete: id %d", id)
	}
	return nil
}

// HardDelete permanently deletes a role
func (r *RoleGormRepository) HardDelete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.Role{}, id).Error; err != nil {
		return errors.Wrapf(err, "RoleGormRepository.HardDelete: id %d", id)
	}
	return nil
}

// Restore restores a deleted role
func (r *RoleGormRepository) Restore(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Model(&models.Role{}).Where("id = ?", id).Updates(map[string]interface{}{
		"is_deleted": false,
		"is_active":  true,
	}).Error; err != nil {
		return errors.Wrapf(err, "RoleGormRepository.Restore: id %d", id)
	}
	return nil
}

// GetSystemRoles retrieves all system roles
func (r *RoleGormRepository) GetSystemRoles(ctx context.Context) ([]models.Role, error) {
	var roles []models.Role
	if err := r.db.WithContext(ctx).Where("is_system = ? AND is_deleted = ?", true, false).Find(&roles).Error; err != nil {
		return nil, errors.Wrap(err, "RoleGormRepository.GetSystemRoles")
	}
	return roles, nil
}

// GetActiveRoles retrieves all active roles
func (r *RoleGormRepository) GetActiveRoles(ctx context.Context) ([]models.Role, error) {
	var roles []models.Role
	if err := r.db.WithContext(ctx).Where("is_active = ? AND is_deleted = ?", true, false).Find(&roles).Error; err != nil {
		return nil, errors.Wrap(err, "RoleGormRepository.GetActiveRoles")
	}
	return roles, nil
}
