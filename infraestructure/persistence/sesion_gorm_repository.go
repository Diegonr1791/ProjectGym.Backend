package persistence

import (
	"time"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"

	"gorm.io/gorm"
)

type SessionGormRepository struct {
	db *gorm.DB
}

func NewSessionGormRepository(db *gorm.DB) repository.SessionRepository {
	return &SessionGormRepository{db}
}

func (r *SessionGormRepository) Create(sesion *model.Sesion) error {
	return r.db.Create(sesion).Error
}

func (r *SessionGormRepository) GetAll() ([]*model.Sesion, error) {
	var sesiones []*model.Sesion
	if err := r.db.Find(&sesiones).Error; err != nil {
		return nil, err
	}
	return sesiones, nil
}

func (r *SessionGormRepository) GetById(id uint) (*model.Sesion, error) {
	var sesion model.Sesion
	if err := r.db.First(&sesion, id).Error; err != nil {
		return nil, err
	}
	return &sesion, nil
}

func (r *SessionGormRepository) Update(sesion *model.Sesion) error {
	return r.db.Save(sesion).Error
}

func (r *SessionGormRepository) Delete(id uint) error {
	return r.db.Delete(&model.Sesion{}, id).Error
}

func (r *SessionGormRepository) GetByUserID(userID uint) ([]*model.Sesion, error) {
	var sesiones []*model.Sesion
	if err := r.db.Where("usuario_id = ?", userID).Find(&sesiones).Error; err != nil {
		return nil, err
	}
	return sesiones, nil
}

func (r *SessionGormRepository) GetByDateRange(startDate, endDate time.Time) ([]*model.Sesion, error) {
	var sesiones []*model.Sesion
	if err := r.db.Where("fecha BETWEEN ? AND ?", startDate, endDate).Find(&sesiones).Error; err != nil {
		return nil, err
	}
	return sesiones, nil
}
