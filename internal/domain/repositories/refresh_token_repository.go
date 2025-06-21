package repository

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
)

type RefreshTokenRepository interface {
	Save(token *model.RefreshToken) error
	FindByTokenHash(tokenHash string) (*model.RefreshToken, error)
	Revoke(tokenHash string) error
	RevokeByUserID(userID uint) error
}
