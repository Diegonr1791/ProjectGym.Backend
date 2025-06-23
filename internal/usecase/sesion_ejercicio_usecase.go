package usecase

import (
	"time"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
)

type SessionExerciseUsecase struct {
	sessionExerciseRepo repositories.SessionExerciseRepository
}

func NewSessionExerciseUsecase(sessionExerciseRepo repositories.SessionExerciseRepository) *SessionExerciseUsecase {
	return &SessionExerciseUsecase{sessionExerciseRepo}
}

func (uc *SessionExerciseUsecase) CreateSessionExercise(sessionExercise *models.SesionEjercicio) error {
	if err := uc.sessionExerciseRepo.Create(sessionExercise); err != nil {
		return domainErrors.NewAppError(500, "DB_CREATE_SESSION_EXERCISE_FAILED", "Failed to create session exercise in database", err)
	}
	return nil
}

func (uc *SessionExerciseUsecase) GetAllSessionExercises() ([]*models.SesionEjercicio, error) {
	sesionesEjercicios, err := uc.sessionExerciseRepo.GetAll()
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_ALL_SESSION_EXERCISES_FAILED", "Failed to get all session exercises from database", err)
	}
	return sesionesEjercicios, nil
}

func (uc *SessionExerciseUsecase) GetSessionExerciseByID(id uint) (*models.SesionEjercicio, error) {
	sesionEjercicio, err := uc.sessionExerciseRepo.GetById(id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, domainErrors.NewAppError(500, "DB_GET_SESSION_EXERCISE_FAILED", "Failed to get session exercise from database", err)
	}
	return sesionEjercicio, nil
}

func (uc *SessionExerciseUsecase) UpdateSessionExercise(sessionExercise *models.SesionEjercicio) error {
	// Verify session exercise exists before updating
	if _, err := uc.sessionExerciseRepo.GetById(sessionExercise.ID); err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return domainErrors.ErrNotFound
		}
		return domainErrors.NewAppError(500, "DB_UPDATE_SESSION_EXERCISE_FAILED", "Failed to verify session exercise existence", err)
	}

	if err := uc.sessionExerciseRepo.Update(sessionExercise); err != nil {
		return domainErrors.NewAppError(500, "DB_UPDATE_SESSION_EXERCISE_FAILED", "Failed to update session exercise in database", err)
	}
	return nil
}

func (uc *SessionExerciseUsecase) DeleteSessionExercise(id uint) error {
	if err := uc.sessionExerciseRepo.Delete(id); err != nil {
		return domainErrors.NewAppError(500, "DB_DELETE_SESSION_EXERCISE_FAILED", "Failed to delete session exercise from database", err)
	}
	return nil
}

func (uc *SessionExerciseUsecase) GetSessionExercisesBySessionID(sessionID uint, fechaDesde, fechaHasta time.Time) ([]*models.SesionEjercicio, error) {
	sesionesEjercicios, err := uc.sessionExerciseRepo.GetBySessionID(sessionID, fechaDesde, fechaHasta)
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_SESSION_EXERCISES_BY_SESSION_FAILED", "Failed to get session exercises by session from database", err)
	}
	return sesionesEjercicios, nil
}
