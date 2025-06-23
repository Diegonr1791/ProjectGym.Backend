package persistence

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ExerciseGormRepository struct {
	db *gorm.DB
}

func NewExerciseGormRepository(db *gorm.DB) repositories.ExerciseRepository {
	return &ExerciseGormRepository{db}
}

func (r *ExerciseGormRepository) GetAll() ([]*models.Ejercicio, error) {
	var ejercicios []*models.Ejercicio
	if err := r.db.Find(&ejercicios).Error; err != nil {
		return nil, errors.Wrap(err, "ExerciseGormRepository.GetAll")
	}
	return ejercicios, nil
}

func (r *ExerciseGormRepository) GetById(id uint) (*models.Ejercicio, error) {
	var ejercicio models.Ejercicio
	if err := r.db.First(&ejercicio, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "ExerciseGormRepository.GetById: id %d", id)
	}
	return &ejercicio, nil
}

func (r *ExerciseGormRepository) Create(ejercicio *models.Ejercicio) error {
	if err := r.db.Create(ejercicio).Error; err != nil {
		return errors.Wrap(err, "ExerciseGormRepository.Create")
	}
	return nil
}

func (r *ExerciseGormRepository) Update(ejercicio *models.Ejercicio) error {
	if err := r.db.Save(ejercicio).Error; err != nil {
		return errors.Wrapf(err, "ExerciseGormRepository.Update: id %d", ejercicio.ID)
	}
	return nil
}

func (r *ExerciseGormRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Ejercicio{}, id).Error; err != nil {
		return errors.Wrapf(err, "ExerciseGormRepository.Delete: id %d", id)
	}
	return nil
}

func (r *ExerciseGormRepository) GetByMuscleGroup(muscleGroupID uint) ([]*models.Ejercicio, error) {
	var ejercicios []*models.Ejercicio
	if err := r.db.Where("grupo_muscular_id = ?", muscleGroupID).Find(&ejercicios).Error; err != nil {
		return nil, errors.Wrapf(err, "ExerciseGormRepository.GetByMuscleGroup: muscleGroupID %d", muscleGroupID)
	}
	return ejercicios, nil
}
