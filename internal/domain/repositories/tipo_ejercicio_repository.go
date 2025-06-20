package repository

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
)

type TypeExerciseRepository interface {
	GetAll() ([]*model.TipoEjercicio, error)
	GetById(id uint) (*model.TipoEjercicio, error)
	Create(tipoEjercicio *model.TipoEjercicio) error
	Update(tipoEjercicio *model.TipoEjercicio) error
	Delete(id uint) error
}
