package persistence

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RutinaGormRepository struct {
	db *gorm.DB
}

func NewRutinaGormRepository(db *gorm.DB) repositories.RutinaRepository {
	return &RutinaGormRepository{db}
}

func (r *RutinaGormRepository) GetAll() ([]models.Rutina, error) {
	var rutinas []models.Rutina
	if err := r.db.Find(&rutinas).Error; err != nil {
		return nil, errors.Wrap(err, "RutinaGormRepository.GetAll")
	}
	return rutinas, nil
}

func (r *RutinaGormRepository) GetByID(id uint) (*models.Rutina, error) {
	var rutina models.Rutina
	if err := r.db.First(&rutina, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "RutinaGormRepository.GetByID: id %d", id)
	}
	return &rutina, nil
}

func (r *RutinaGormRepository) Create(rutina *models.Rutina) error {
	if err := r.db.Create(rutina).Error; err != nil {
		return errors.Wrap(err, "RutinaGormRepository.Create")
	}
	return nil
}

func (r *RutinaGormRepository) Update(rutina *models.Rutina) error {
	if err := r.db.Save(rutina).Error; err != nil {
		return errors.Wrapf(err, "RutinaGormRepository.Update: id %d", rutina.ID)
	}
	return nil
}

func (r *RutinaGormRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Rutina{}, id).Error; err != nil {
		return errors.Wrapf(err, "RutinaGormRepository.Delete: id %d", id)
	}
	return nil
}
