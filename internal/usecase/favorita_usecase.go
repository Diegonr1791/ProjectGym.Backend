package usecase

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
)

type FavoriteUsecase struct {
	repo repositories.FavoritaRepository
}

func NewFavoriteUsecase(repo repositories.FavoritaRepository) *FavoriteUsecase {
	return &FavoriteUsecase{repo}
}

func (uc *FavoriteUsecase) GetAll() ([]models.Favorita, error) {
	favoritas, err := uc.repo.GetAll()
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_ALL_FAVORITES_FAILED", "Failed to get all favorites from database", err)
	}
	return favoritas, nil
}

func (uc *FavoriteUsecase) GetByID(id uint) (*models.Favorita, error) {
	favorita, err := uc.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, domainErrors.NewAppError(500, "DB_GET_FAVORITE_FAILED", "Failed to get favorite from database", err)
	}
	return favorita, nil
}

func (uc *FavoriteUsecase) Create(favorita *models.Favorita) error {
	if err := uc.repo.Create(favorita); err != nil {
		return domainErrors.NewAppError(500, "DB_CREATE_FAVORITE_FAILED", "Failed to create favorite in database", err)
	}
	return nil
}

func (uc *FavoriteUsecase) Update(favorita *models.Favorita) error {
	if _, err := uc.repo.GetByID(favorita.ID); err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return domainErrors.ErrNotFound
		}
		return domainErrors.NewAppError(500, "DB_UPDATE_FAVORITE_FAILED", "Failed to verify favorite existence", err)
	}
	if err := uc.repo.Update(favorita); err != nil {
		return domainErrors.NewAppError(500, "DB_UPDATE_FAVORITE_FAILED", "Failed to update favorite in database", err)
	}
	return nil
}

func (uc *FavoriteUsecase) Delete(id uint) error {
	if err := uc.repo.Delete(id); err != nil {
		return domainErrors.NewAppError(500, "DB_DELETE_FAVORITE_FAILED", "Failed to delete favorite from database", err)
	}
	return nil
}

func (uc *FavoriteUsecase) GetByUserID(userID uint) ([]models.Favorita, error) {
	favoritas, err := uc.repo.GetFavoritasByUsuarioID(userID)
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_FAVORITES_BY_USER_FAILED", "Failed to get favorites by user from database", err)
	}
	return favoritas, nil
}
