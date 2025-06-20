package repository

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
)

type ExerciseRepository interface {
	GetAll() ([]*model.Ejercicio, error)
	GetById(id uint) (*model.Ejercicio, error)
	Create(ejercicio *model.Ejercicio) error
	Update(ejercicio *model.Ejercicio) error
	Delete(id uint) error
	GetByMuscleGroup(muscleGroupID uint) ([]*model.Ejercicio, error)
}
