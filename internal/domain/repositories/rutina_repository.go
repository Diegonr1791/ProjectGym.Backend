package repository

import model "github.com/Diegonr1791/GymBro/internal/domain/models"

type RutinaRepository interface {
	GetAll() ([]model.Rutina, error)
	GetByID(id uint) (*model.Rutina, error)
	Create(rutina *model.Rutina) error
	Update(rutina *model.Rutina) error
	Delete(id uint) error
}
