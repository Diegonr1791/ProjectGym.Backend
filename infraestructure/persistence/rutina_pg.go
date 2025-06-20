package persistence

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	"gorm.io/gorm"
)

type RutinaPgRepo struct {
	db *gorm.DB
}

func NewRutinaPgRepo(db *gorm.DB) *RutinaPgRepo {
	return &RutinaPgRepo{db}
}

func (r *RutinaPgRepo) GetAll() ([]model.Rutina, error) {
	var rutinas []model.Rutina
	if err := r.db.Find(&rutinas).Error; err != nil {
		return nil, err
	}

	return rutinas, nil
}

func (r *RutinaPgRepo) GetByID(id uint) (*model.Rutina, error) {
	var rutina model.Rutina
	if err := r.db.First(&rutina, id).Error; err != nil {
		return nil, err
	}

	return &rutina, nil
}

func (r *RutinaPgRepo) Create(rutina *model.Rutina) error {
	return r.db.Create(rutina).Error
}

func (r *RutinaPgRepo) Update(rutina *model.Rutina) error {
	return r.db.Save(rutina).Error
}

func (r *RutinaPgRepo) Delete(id uint) error {
	return r.db.Delete(&model.Rutina{}, id).Error
}
