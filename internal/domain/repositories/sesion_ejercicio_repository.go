package repository

import (
	"time"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
)

type SessionExerciseRepository interface {
	Create(sesionEjercicio *model.SesionEjercicio) error
	GetAll() ([]*model.SesionEjercicio, error)
	GetById(id uint) (*model.SesionEjercicio, error)
	Update(sesionEjercicio *model.SesionEjercicio) error
	Delete(id uint) error
	GetBySessionID(sessionID uint, fechaDesde, fechaHasta time.Time) ([]*model.SesionEjercicio, error)
}
