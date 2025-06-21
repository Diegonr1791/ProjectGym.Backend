package persistence

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	"gorm.io/gorm"
)

type refreshTokenGormRepository struct {
	db *gorm.DB
}

func NewRefreshTokenGormRepository(db *gorm.DB) *refreshTokenGormRepository {
	return &refreshTokenGormRepository{db}
}

func (r *refreshTokenGormRepository) Save(token *model.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *refreshTokenGormRepository) FindByTokenHash(tokenHash string) (*model.RefreshToken, error) {
	var token model.RefreshToken
	err := r.db.Where("token_hash = ?", tokenHash).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *refreshTokenGormRepository) Revoke(tokenHash string) error {
	return r.db.Where("token_hash = ?", tokenHash).Delete(&model.RefreshToken{}).Error
}

func (r *refreshTokenGormRepository) RevokeByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&model.RefreshToken{}).Error
}
