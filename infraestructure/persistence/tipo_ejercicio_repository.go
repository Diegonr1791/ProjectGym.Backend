package persistence

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"

	"gorm.io/gorm"
)

type TypeExerciseGormRepository struct {
	db *gorm.DB
}

func NewTypeExerciseGormRepository(db *gorm.DB) repository.TypeExerciseRepository {
	return &TypeExerciseGormRepository{db}
}

func (r *TypeExerciseGormRepository) GetAll() ([]*model.TipoEjercicio, error) {
	var tiposEjercicio []*model.TipoEjercicio
	if err := r.db.Find(&tiposEjercicio).Error; err != nil {
		return nil, err
	}
	return tiposEjercicio, nil
}

func (r *TypeExerciseGormRepository) GetById(id uint) (*model.TipoEjercicio, error) {
	var tipoEjercicio model.TipoEjercicio
	if err := r.db.First(&tipoEjercicio, id).Error; err != nil {
		return nil, err
	}
	return &tipoEjercicio, nil
}

func (r *TypeExerciseGormRepository) Create(tipoEjercicio *model.TipoEjercicio) error {
	return r.db.Create(tipoEjercicio).Error
}

func (r *TypeExerciseGormRepository) Update(tipoEjercicio *model.TipoEjercicio) error {
	return r.db.Save(tipoEjercicio).Error
}

func (r *TypeExerciseGormRepository) Delete(id uint) error {
	return r.db.Delete(&model.TipoEjercicio{}, id).Error
}
