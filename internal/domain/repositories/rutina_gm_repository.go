package repository

import model "github.com/Diegonr1791/GymBro/internal/domain/models"

type RutinaGrupoMuscularRepository interface {
	GetAll() ([]model.RutinaGrupoMuscular, error)
	GetByID(id uint) (*model.RutinaGrupoMuscular, error)
	Create(rutina *model.RutinaGrupoMuscular) error
	Update(rutina *model.RutinaGrupoMuscular) error
	Delete(id uint) error
	GetMusclesGroupByRutine(id uint) (*model.RutinaConGruposMusculares, error)
}
