package usecase

import (
	"time"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"
)

type SessionUsecase struct {
	sesionRepo repository.SessionRepository
}

func NewSessionUsecase(sesionRepo repository.SessionRepository) *SessionUsecase {
	return &SessionUsecase{sesionRepo}
}

func (u *SessionUsecase) CreateSession(sesion *model.Sesion) error {
	return u.sesionRepo.Create(sesion)
}

func (u *SessionUsecase) GetAllSessions() ([]*model.Sesion, error) {
	return u.sesionRepo.GetAll()
}

func (u *SessionUsecase) GetSessionById(id uint) (*model.Sesion, error) {
	return u.sesionRepo.GetById(id)
}

func (u *SessionUsecase) UpdateSession(sesion *model.Sesion) error {
	return u.sesionRepo.Update(sesion)
}

func (u *SessionUsecase) DeleteSession(id uint) error {
	return u.sesionRepo.Delete(id)
}

func (u *SessionUsecase) GetSessionsByUserID(userID uint) ([]*model.Sesion, error) {
	return u.sesionRepo.GetByUserID(userID)
}

func (u *SessionUsecase) GetSessionsByDateRange(startDate, endDate time.Time) ([]*model.Sesion, error) {
	return u.sesionRepo.GetByDateRange(startDate, endDate)
}
