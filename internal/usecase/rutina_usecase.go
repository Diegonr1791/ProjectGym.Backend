package usecase

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
)

type RutinaUsecase struct {
	repo repositories.RutinaRepository
}

func NewRutinaUsecase(repo repositories.RutinaRepository) *RutinaUsecase {
	return &RutinaUsecase{repo}
}

func (uc *RutinaUsecase) GetAllRoutines() ([]models.Rutina, error) {
	rutinas, err := uc.repo.GetAll()
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_ALL_ROUTINES_FAILED", "Failed to get all routines from database", err)
	}
	return rutinas, nil
}

func (uc *RutinaUsecase) GetRoutineByID(id uint) (*models.Rutina, error) {
	rutina, err := uc.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, domainErrors.NewAppError(500, "DB_GET_ROUTINE_FAILED", "Failed to get routine from database", err)
	}
	return rutina, nil
}

func (uc *RutinaUsecase) CreateRoutine(rutina *models.Rutina) error {
	if err := uc.repo.Create(rutina); err != nil {
		return domainErrors.NewAppError(500, "DB_CREATE_ROUTINE_FAILED", "Failed to create routine in database", err)
	}
	return nil
}

func (uc *RutinaUsecase) UpdateRoutine(rutina *models.Rutina) error {
	// Verify routine exists before updating
	if _, err := uc.repo.GetByID(rutina.ID); err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return domainErrors.ErrNotFound
		}
		return domainErrors.NewAppError(500, "DB_UPDATE_ROUTINE_FAILED", "Failed to verify routine existence", err)
	}

	if err := uc.repo.Update(rutina); err != nil {
		return domainErrors.NewAppError(500, "DB_UPDATE_ROUTINE_FAILED", "Failed to update routine in database", err)
	}
	return nil
}

func (uc *RutinaUsecase) DeleteRoutine(id uint) error {
	if err := uc.repo.Delete(id); err != nil {
		return domainErrors.NewAppError(500, "DB_DELETE_ROUTINE_FAILED", "Failed to delete routine from database", err)
	}
	return nil
}
