package repository

import model "github.com/Diegonr1791/GymBro/internal/domain/models"

type GrupoMuscularRepository interface {
	Create(g *model.GrupoMuscular) error
	GetAll() ([]model.GrupoMuscular, error)
	GetByID(id uint) (*model.GrupoMuscular, error)
	Update(g *model.GrupoMuscular) error
	Delete(id uint) error
}
