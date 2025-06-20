package repository

import (
	"time"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
)

type SessionRepository interface {
	Create(sesion *model.Sesion) error
	GetAll() ([]*model.Sesion, error)
	GetById(id uint) (*model.Sesion, error)
	Update(sesion *model.Sesion) error
	Delete(id uint) error
	GetByUserID(userID uint) ([]*model.Sesion, error)
	GetByDateRange(startDate, endDate time.Time) ([]*model.Sesion, error)
}
