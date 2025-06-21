package model

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `gorm:"not null"`
	User      Usuario        `gorm:"foreignKey:UserID"`
	TokenHash string         `gorm:"not null;uniqueIndex"`
	ExpiresAt time.Time      `gorm:"not null"`
	RevokedAt gorm.DeletedAt `gorm:"index"`
}
