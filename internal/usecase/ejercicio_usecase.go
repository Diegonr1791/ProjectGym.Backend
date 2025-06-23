package usecase

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
)

type ExerciseUsecase struct {
	exerciseRepo repositories.ExerciseRepository
}

func NewExerciseUsecase(exerciseRepo repositories.ExerciseRepository) *ExerciseUsecase {
	return &ExerciseUsecase{exerciseRepo}
}

func (uc *ExerciseUsecase) GetAllExercises() ([]*models.Ejercicio, error) {
	ejercicios, err := uc.exerciseRepo.GetAll()
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_ALL_EXERCISES_FAILED", "Failed to get all exercises from database", err)
	}
	return ejercicios, nil
}

func (uc *ExerciseUsecase) GetExerciseByID(id uint) (*models.Ejercicio, error) {
	ejercicio, err := uc.exerciseRepo.GetById(id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, domainErrors.NewAppError(500, "DB_GET_EXERCISE_FAILED", "Failed to get exercise from database", err)
	}
	return ejercicio, nil
}

func (uc *ExerciseUsecase) CreateExercise(exercise *models.Ejercicio) error {
	if err := uc.exerciseRepo.Create(exercise); err != nil {
		return domainErrors.NewAppError(500, "DB_CREATE_EXERCISE_FAILED", "Failed to create exercise in database", err)
	}
	return nil
}

func (uc *ExerciseUsecase) UpdateExercise(exercise *models.Ejercicio) error {
	// Verify exercise exists before updating
	if _, err := uc.exerciseRepo.GetById(exercise.ID); err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return domainErrors.ErrNotFound
		}
		return domainErrors.NewAppError(500, "DB_UPDATE_EXERCISE_FAILED", "Failed to verify exercise existence", err)
	}

	if err := uc.exerciseRepo.Update(exercise); err != nil {
		return domainErrors.NewAppError(500, "DB_UPDATE_EXERCISE_FAILED", "Failed to update exercise in database", err)
	}
	return nil
}

func (uc *ExerciseUsecase) DeleteExercise(id uint) error {
	if err := uc.exerciseRepo.Delete(id); err != nil {
		return domainErrors.NewAppError(500, "DB_DELETE_EXERCISE_FAILED", "Failed to delete exercise from database", err)
	}
	return nil
}

func (uc *ExerciseUsecase) GetExercisesByMuscleGroup(muscleGroupID uint) ([]*models.Ejercicio, error) {
	ejercicios, err := uc.exerciseRepo.GetByMuscleGroup(muscleGroupID)
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_EXERCISES_BY_MUSCLE_GROUP_FAILED", "Failed to get exercises by muscle group from database", err)
	}
	return ejercicios, nil
}
