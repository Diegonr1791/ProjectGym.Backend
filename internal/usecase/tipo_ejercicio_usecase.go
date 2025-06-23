package usecase

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
)

type TypeExerciseUsecase struct {
	typeExerciseRepo repositories.TypeExerciseRepository
}

func NewTypeExerciseUsecase(typeExerciseRepo repositories.TypeExerciseRepository) *TypeExerciseUsecase {
	return &TypeExerciseUsecase{typeExerciseRepo}
}

func (uc *TypeExerciseUsecase) GetAllExerciseTypes() ([]*models.TipoEjercicio, error) {
	tiposEjercicio, err := uc.typeExerciseRepo.GetAll()
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_ALL_EXERCISE_TYPES_FAILED", "Failed to get all exercise types from database", err)
	}
	return tiposEjercicio, nil
}

func (uc *TypeExerciseUsecase) GetExerciseTypeByID(id uint) (*models.TipoEjercicio, error) {
	tipoEjercicio, err := uc.typeExerciseRepo.GetById(id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, domainErrors.NewAppError(500, "DB_GET_EXERCISE_TYPE_FAILED", "Failed to get exercise type from database", err)
	}
	return tipoEjercicio, nil
}

func (uc *TypeExerciseUsecase) CreateExerciseType(tipoEjercicio *models.TipoEjercicio) error {
	if err := uc.typeExerciseRepo.Create(tipoEjercicio); err != nil {
		return domainErrors.NewAppError(500, "DB_CREATE_EXERCISE_TYPE_FAILED", "Failed to create exercise type in database", err)
	}
	return nil
}

func (uc *TypeExerciseUsecase) UpdateExerciseType(tipoEjercicio *models.TipoEjercicio) error {
	// Verify exercise type exists before updating
	if _, err := uc.typeExerciseRepo.GetById(tipoEjercicio.ID); err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return domainErrors.ErrNotFound
		}
		return domainErrors.NewAppError(500, "DB_UPDATE_EXERCISE_TYPE_FAILED", "Failed to verify exercise type existence", err)
	}

	if err := uc.typeExerciseRepo.Update(tipoEjercicio); err != nil {
		return domainErrors.NewAppError(500, "DB_UPDATE_EXERCISE_TYPE_FAILED", "Failed to update exercise type in database", err)
	}
	return nil
}

func (uc *TypeExerciseUsecase) DeleteExerciseType(id uint) error {
	if err := uc.typeExerciseRepo.Delete(id); err != nil {
		return domainErrors.NewAppError(500, "DB_DELETE_EXERCISE_TYPE_FAILED", "Failed to delete exercise type from database", err)
	}
	return nil
}
