package usecase

import (
	"time"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"
)

type SessionExerciseUsecase struct {
	sessionExerciseRepo repository.SessionExerciseRepository
}

func NewSessionExerciseUsecase(sessionExerciseRepo repository.SessionExerciseRepository) *SessionExerciseUsecase {
	return &SessionExerciseUsecase{sessionExerciseRepo}
}

func (u *SessionExerciseUsecase) CreateSessionExercise(sessionExercise *model.SesionEjercicio) error {
	return u.sessionExerciseRepo.Create(sessionExercise)
}

func (u *SessionExerciseUsecase) GetAllSessionExercises() ([]*model.SesionEjercicio, error) {
	return u.sessionExerciseRepo.GetAll()
}

func (u *SessionExerciseUsecase) GetSessionExerciseById(id uint) (*model.SesionEjercicio, error) {
	return u.sessionExerciseRepo.GetById(id)
}

func (u *SessionExerciseUsecase) UpdateSessionExercise(sessionExercise *model.SesionEjercicio) error {
	return u.sessionExerciseRepo.Update(sessionExercise)
}

func (u *SessionExerciseUsecase) DeleteSessionExercise(id uint) error {
	return u.sessionExerciseRepo.Delete(id)
}

func (u *SessionExerciseUsecase) GetSessionExercisesBySessionID(sessionID uint, fechaDesde, fechaHasta time.Time) ([]*model.SesionEjercicio, error) {
	return u.sessionExerciseRepo.GetBySessionID(sessionID, fechaDesde, fechaHasta)
}
