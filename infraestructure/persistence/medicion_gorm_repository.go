package persistence

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MedicionGormRepository struct {
	db *gorm.DB
}

func NewMedicionGormRepository(db *gorm.DB) repositories.MedicionRepository {
	return &MedicionGormRepository{db}
}

func (r *MedicionGormRepository) GetAll() ([]models.Medicion, error) {
	var mediciones []models.Medicion
	if err := r.db.Find(&mediciones).Error; err != nil {
		return nil, errors.Wrap(err, "MedicionGormRepository.GetAll")
	}
	return mediciones, nil
}

func (r *MedicionGormRepository) GetByID(id uint) (*models.Medicion, error) {
	var medicion models.Medicion
	if err := r.db.First(&medicion, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "MedicionGormRepository.GetByID: id %d", id)
	}
	return &medicion, nil
}

func (r *MedicionGormRepository) Create(medicion *models.Medicion) error {
	if err := r.db.Create(medicion).Error; err != nil {
		return errors.Wrap(err, "MedicionGormRepository.Create")
	}
	return nil
}

func (r *MedicionGormRepository) Update(medicion *models.Medicion) error {
	if err := r.db.Save(medicion).Error; err != nil {
		return errors.Wrapf(err, "MedicionGormRepository.Update: id %d", medicion.ID)
	}
	return nil
}

func (r *MedicionGormRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Medicion{}, id).Error; err != nil {
		return errors.Wrapf(err, "MedicionGormRepository.Delete: id %d", id)
	}
	return nil
}

func (r *MedicionGormRepository) GetMesurementsByUserID(usuarioID uint) ([]models.Medicion, error) {
	var mediciones []models.Medicion
	if err := r.db.Where("usuario_id = ?", usuarioID).Find(&mediciones).Error; err != nil {
		return nil, errors.Wrapf(err, "MedicionGormRepository.GetMesurementsByUserID: usuarioID %d", usuarioID)
	}
	return mediciones, nil
}
