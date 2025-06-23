package usecase

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
)

type MeasurementUsecase struct {
	repo repositories.MedicionRepository
}

func NewMeasurementUsecase(repo repositories.MedicionRepository) *MeasurementUsecase {
	return &MeasurementUsecase{repo}
}

func (uc *MeasurementUsecase) GetAll() ([]models.Medicion, error) {
	mediciones, err := uc.repo.GetAll()
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_ALL_MEASUREMENTS_FAILED", "Failed to get all measurements from database", err)
	}
	return mediciones, nil
}

func (uc *MeasurementUsecase) GetByID(id uint) (*models.Medicion, error) {
	medicion, err := uc.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, domainErrors.NewAppError(500, "DB_GET_MEASUREMENT_FAILED", "Failed to get measurement from database", err)
	}
	return medicion, nil
}

func (uc *MeasurementUsecase) Create(medicion *models.Medicion) error {
	if err := uc.repo.Create(medicion); err != nil {
		return domainErrors.NewAppError(500, "DB_CREATE_MEASUREMENT_FAILED", "Failed to create measurement in database", err)
	}
	return nil
}

func (uc *MeasurementUsecase) Update(medicion *models.Medicion) error {
	if _, err := uc.repo.GetByID(medicion.ID); err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return domainErrors.ErrNotFound
		}
		return domainErrors.NewAppError(500, "DB_UPDATE_MEASUREMENT_FAILED", "Failed to verify measurement existence", err)
	}
	if err := uc.repo.Update(medicion); err != nil {
		return domainErrors.NewAppError(500, "DB_UPDATE_MEASUREMENT_FAILED", "Failed to update measurement in database", err)
	}
	return nil
}

func (uc *MeasurementUsecase) Delete(id uint) error {
	if err := uc.repo.Delete(id); err != nil {
		return domainErrors.NewAppError(500, "DB_DELETE_MEASUREMENT_FAILED", "Failed to delete measurement from database", err)
	}
	return nil
}

func (uc *MeasurementUsecase) GetByUserID(userID uint) ([]models.Medicion, error) {
	mediciones, err := uc.repo.GetMesurementsByUserID(userID)
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_MEASUREMENTS_BY_USER_FAILED", "Failed to get measurements by user from database", err)
	}
	return mediciones, nil
}
