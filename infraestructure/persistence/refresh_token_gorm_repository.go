package persistence

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RefreshTokenGormRepository struct {
	db *gorm.DB
}

func NewRefreshTokenGormRepository(db *gorm.DB) repositories.RefreshTokenRepository {
	return &RefreshTokenGormRepository{db}
}

func (r *RefreshTokenGormRepository) Save(token *models.RefreshToken) error {
	if err := r.db.Create(token).Error; err != nil {
		return errors.Wrap(err, "RefreshTokenGormRepository.Save")
	}
	return nil
}

func (r *RefreshTokenGormRepository) FindByTokenHash(tokenHash string) (*models.RefreshToken, error) {
	var token models.RefreshToken
	err := r.db.Where("token_hash = ?", tokenHash).First(&token).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "RefreshTokenGormRepository.FindByTokenHash: tokenHash %s", tokenHash)
	}
	return &token, nil
}

func (r *RefreshTokenGormRepository) Revoke(tokenHash string) error {
	if err := r.db.Where("token_hash = ?", tokenHash).Delete(&models.RefreshToken{}).Error; err != nil {
		return errors.Wrapf(err, "RefreshTokenGormRepository.Revoke: tokenHash %s", tokenHash)
	}
	return nil
}

func (r *RefreshTokenGormRepository) RevokeByUserID(userID uint) error {
	if err := r.db.Where("user_id = ?", userID).Delete(&models.RefreshToken{}).Error; err != nil {
		return errors.Wrapf(err, "RefreshTokenGormRepository.RevokeByUserID: userID %d", userID)
	}
	return nil
}
