package usecase

import (
	"time"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
)

type SessionUsecase struct {
	sesionRepo repositories.SessionRepository
}

func NewSessionUsecase(sesionRepo repositories.SessionRepository) *SessionUsecase {
	return &SessionUsecase{sesionRepo}
}

func (uc *SessionUsecase) CreateSession(sesion *models.Sesion) error {
	if err := uc.sesionRepo.Create(sesion); err != nil {
		return domainErrors.NewAppError(500, "DB_CREATE_SESSION_FAILED", "Failed to create session in database", err)
	}
	return nil
}

func (uc *SessionUsecase) GetAllSessions() ([]*models.Sesion, error) {
	sesiones, err := uc.sesionRepo.GetAll()
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_ALL_SESSIONS_FAILED", "Failed to get all sessions from database", err)
	}
	return sesiones, nil
}

func (uc *SessionUsecase) GetSessionByID(id uint) (*models.Sesion, error) {
	sesion, err := uc.sesionRepo.GetById(id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, domainErrors.NewAppError(500, "DB_GET_SESSION_FAILED", "Failed to get session from database", err)
	}
	return sesion, nil
}

func (uc *SessionUsecase) UpdateSession(sesion *models.Sesion) error {
	// Verify session exists before updating
	if _, err := uc.sesionRepo.GetById(sesion.ID); err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return domainErrors.ErrNotFound
		}
		return domainErrors.NewAppError(500, "DB_UPDATE_SESSION_FAILED", "Failed to verify session existence", err)
	}

	if err := uc.sesionRepo.Update(sesion); err != nil {
		return domainErrors.NewAppError(500, "DB_UPDATE_SESSION_FAILED", "Failed to update session in database", err)
	}
	return nil
}

func (uc *SessionUsecase) DeleteSession(id uint) error {
	if err := uc.sesionRepo.Delete(id); err != nil {
		return domainErrors.NewAppError(500, "DB_DELETE_SESSION_FAILED", "Failed to delete session from database", err)
	}
	return nil
}

func (uc *SessionUsecase) GetSessionsByUserID(userID uint) ([]*models.Sesion, error) {
	sesiones, err := uc.sesionRepo.GetByUserID(userID)
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_SESSIONS_BY_USER_FAILED", "Failed to get sessions by user from database", err)
	}
	return sesiones, nil
}

func (uc *SessionUsecase) GetSessionsByDateRange(startDate, endDate time.Time) ([]*models.Sesion, error) {
	sesiones, err := uc.sesionRepo.GetByDateRange(startDate, endDate)
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_SESSIONS_BY_DATE_RANGE_FAILED", "Failed to get sessions by date range from database", err)
	}
	return sesiones, nil
}
