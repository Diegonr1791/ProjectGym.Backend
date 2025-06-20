package repository

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
)

type MedicionRepository interface {
	GetAll() ([]model.Medicion, error)
	GetByID(id uint) (*model.Medicion, error)
	Create(medicion *model.Medicion) error
	Update(medicion *model.Medicion) error
	Delete(id uint) error
	GetMesurementsByUserID(usuarioID uint) ([]model.Medicion, error)
}
