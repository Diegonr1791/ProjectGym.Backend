package persistence

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TypeExerciseGormRepository struct {
	db *gorm.DB
}

func NewTypeExerciseGormRepository(db *gorm.DB) repositories.TypeExerciseRepository {
	return &TypeExerciseGormRepository{db}
}

func (r *TypeExerciseGormRepository) GetAll() ([]*models.TipoEjercicio, error) {
	var tiposEjercicio []*models.TipoEjercicio
	if err := r.db.Find(&tiposEjercicio).Error; err != nil {
		return nil, errors.Wrap(err, "TypeExerciseGormRepository.GetAll")
	}
	return tiposEjercicio, nil
}

func (r *TypeExerciseGormRepository) GetById(id uint) (*models.TipoEjercicio, error) {
	var tipoEjercicio models.TipoEjercicio
	if err := r.db.First(&tipoEjercicio, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "TypeExerciseGormRepository.GetById: id %d", id)
	}
	return &tipoEjercicio, nil
}

func (r *TypeExerciseGormRepository) Create(tipoEjercicio *models.TipoEjercicio) error {
	if err := r.db.Create(tipoEjercicio).Error; err != nil {
		return errors.Wrap(err, "TypeExerciseGormRepository.Create")
	}
	return nil
}

func (r *TypeExerciseGormRepository) Update(tipoEjercicio *models.TipoEjercicio) error {
	if err := r.db.Save(tipoEjercicio).Error; err != nil {
		return errors.Wrapf(err, "TypeExerciseGormRepository.Update: id %d", tipoEjercicio.ID)
	}
	return nil
}

func (r *TypeExerciseGormRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.TipoEjercicio{}, id).Error; err != nil {
		return errors.Wrapf(err, "TypeExerciseGormRepository.Delete: id %d", id)
	}
	return nil
}
