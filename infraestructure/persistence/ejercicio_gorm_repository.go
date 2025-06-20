package persistence

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"

	"gorm.io/gorm"
)

type ExerciseGormRepository struct {
	db *gorm.DB
}

func NewExerciseGormRepository(db *gorm.DB) repository.ExerciseRepository {
	return &ExerciseGormRepository{db}
}

func (r *ExerciseGormRepository) GetAll() ([]*model.Ejercicio, error) {
	var ejercicios []*model.Ejercicio
	if err := r.db.Find(&ejercicios).Error; err != nil {
		return nil, err
	}
	return ejercicios, nil
}

func (r *ExerciseGormRepository) GetById(id uint) (*model.Ejercicio, error) {
	var ejercicio model.Ejercicio
	if err := r.db.First(&ejercicio, id).Error; err != nil {
		return nil, err
	}
	return &ejercicio, nil
}

func (r *ExerciseGormRepository) Create(ejercicio *model.Ejercicio) error {
	return r.db.Create(ejercicio).Error
}

func (r *ExerciseGormRepository) Update(ejercicio *model.Ejercicio) error {
	return r.db.Save(ejercicio).Error
}

func (r *ExerciseGormRepository) Delete(id uint) error {
	return r.db.Delete(&model.Ejercicio{}, id).Error
}

func (r *ExerciseGormRepository) GetByMuscleGroup(muscleGroupID uint) ([]*model.Ejercicio, error) {
	var ejercicios []*model.Ejercicio
	if err := r.db.Where("grupo_muscular_id = ?", muscleGroupID).Find(&ejercicios).Error; err != nil {
		return nil, err
	}
	return ejercicios, nil
}
