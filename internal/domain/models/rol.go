// internal/domain/models/role.go
package models

import (
	"time"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	"gorm.io/gorm"
)

// Role represents a role in the system
// @Description Role entity for user permissions and access control
type Role struct {
	ID          uint   `gorm:"primaryKey" json:"id" example:"1"`
	Name        string `gorm:"unique;not null" json:"name" example:"admin"`
	Description string `json:"description" example:"Administrator role with full access"`
	IsActive    bool   `gorm:"default:true" json:"is_active" example:"true"`
	IsSystem    bool   `gorm:"default:false" json:"is_system" example:"true"`   // System roles cannot be deleted
	Priority    int    `gorm:"default:0" json:"priority" example:"1"`           // For role hierarchy
	IsDeleted   bool   `gorm:"default:false" json:"is_deleted" example:"false"` // Soft delete

	// Audit fields
	CreatedAt time.Time      `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" swaggerignore:"true"`
}

func (Role) TableName() string {
	return "roles"
}

// BeforeCreate hook to set default values
func (r *Role) BeforeCreate(tx *gorm.DB) error {
	if r.CreatedAt.IsZero() {
		r.CreatedAt = time.Now()
	}
	if r.UpdatedAt.IsZero() {
		r.UpdatedAt = time.Now()
	}
	return nil
}

// BeforeUpdate hook to update UpdatedAt
func (r *Role) BeforeUpdate(tx *gorm.DB) error {
	r.UpdatedAt = time.Now()
	return nil
}

// SoftDelete marks the role as logically deleted
func (r *Role) SoftDelete() error {
	if r.IsSystem {
		return domainErrors.ErrSystemRoleNotDeletable
	}
	r.IsDeleted = true
	r.IsActive = false
	r.UpdatedAt = time.Now()
	return nil
}

// Restore restores a logically deleted role
func (r *Role) Restore() error {
	r.IsDeleted = false
	r.IsActive = true
	r.UpdatedAt = time.Now()
	return nil
}

// IsSoftDeleted checks if the role is logically deleted
func (r *Role) IsSoftDeleted() bool {
	return r.IsDeleted
}

// System role constants
const (
	RoleAdmin = "admin"
	RoleUser  = "user"
	RoleDev   = "dev"
)
