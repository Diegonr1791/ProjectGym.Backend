package persistence

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	"gorm.io/gorm"
)

type MedicionGormRepository struct {
	db *gorm.DB
}

func NewMedicionGormRepository(db *gorm.DB) *MedicionGormRepository {
	return &MedicionGormRepository{db}
}

func (r *MedicionGormRepository) GetAll() ([]model.Medicion, error) {
	var mediciones []model.Medicion
	if err := r.db.Find(&mediciones).Error; err != nil {
		return nil, err
	}
	return mediciones, nil
}

func (r *MedicionGormRepository) GetByID(id uint) (*model.Medicion, error) {
	var medicion model.Medicion
	if err := r.db.First(&medicion, id).Error; err != nil {
		return nil, err
	}
	return &medicion, nil
}

func (r *MedicionGormRepository) Create(medicion *model.Medicion) error {
	return r.db.Create(medicion).Error
}

func (r *MedicionGormRepository) Update(medicion *model.Medicion) error {
	return r.db.Save(medicion).Error
}

func (r *MedicionGormRepository) Delete(id uint) error {
	return r.db.Delete(&model.Medicion{}, id).Error
}

func (r *MedicionGormRepository) GetMesurementsByUserID(usuarioID uint) ([]model.Medicion, error) {
	var mediciones []model.Medicion
	if err := r.db.Where("usuario_id = ?", usuarioID).Find(&mediciones).Error; err != nil {
		return nil, err
	}
	return mediciones, nil
}
