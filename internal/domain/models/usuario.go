package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
// @Description User entity for authentication and profile management
type User struct {
	ID        uint   `gorm:"primaryKey" json:"id" example:"1"`
	Name      string `gorm:"not null" json:"name" example:"John Doe"`
	Email     string `gorm:"unique;not null" json:"email" example:"john@example.com"`
	Password  string `gorm:"not null" json:"password" example:"password123"`
	RoleID    uint   `gorm:"not null" json:"role_id" example:"1"`
	IsActive  bool   `gorm:"default:true" json:"is_active" example:"true"`
	IsDeleted bool   `gorm:"default:false" json:"is_deleted" example:"false"` // Soft delete

	// Relations
	Role         Role           `gorm:"foreignKey:RoleID" json:"role,omitempty" swaggerignore:"true"`
	RefreshToken []RefreshToken `gorm:"foreignKey:UserID" json:"-"`

	// Audit fields
	CreatedAt time.Time      `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" swaggerignore:"true"`
}

func (User) TableName() string {
	return "usuarios"
}

// BeforeCreate hook to set default values
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	if u.UpdatedAt.IsZero() {
		u.UpdatedAt = time.Now()
	}
	return nil
}

// BeforeUpdate hook to update UpdatedAt
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}

// SoftDelete marks the user as logically deleted
func (u *User) SoftDelete() error {
	u.IsDeleted = true
	u.IsActive = false
	u.UpdatedAt = time.Now()
	return nil
}

// Restore restores a logically deleted user
func (u *User) Restore() error {
	u.IsDeleted = false
	u.IsActive = true
	u.UpdatedAt = time.Now()
	return nil
}

// IsSoftDeleted checks if the user is logically deleted
func (u *User) IsSoftDeleted() bool {
	return u.IsDeleted
}
