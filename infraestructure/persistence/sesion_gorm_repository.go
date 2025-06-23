package persistence

import (
	"time"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SessionGormRepository struct {
	db *gorm.DB
}

func NewSessionGormRepository(db *gorm.DB) repositories.SessionRepository {
	return &SessionGormRepository{db}
}

func (r *SessionGormRepository) Create(sesion *models.Sesion) error {
	if err := r.db.Create(sesion).Error; err != nil {
		return errors.Wrap(err, "SessionGormRepository.Create")
	}
	return nil
}

func (r *SessionGormRepository) GetAll() ([]*models.Sesion, error) {
	var sesiones []*models.Sesion
	if err := r.db.Find(&sesiones).Error; err != nil {
		return nil, errors.Wrap(err, "SessionGormRepository.GetAll")
	}
	return sesiones, nil
}

func (r *SessionGormRepository) GetById(id uint) (*models.Sesion, error) {
	var sesion models.Sesion
	if err := r.db.First(&sesion, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "SessionGormRepository.GetById: id %d", id)
	}
	return &sesion, nil
}

func (r *SessionGormRepository) Update(sesion *models.Sesion) error {
	if err := r.db.Save(sesion).Error; err != nil {
		return errors.Wrapf(err, "SessionGormRepository.Update: id %d", sesion.ID)
	}
	return nil
}

func (r *SessionGormRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Sesion{}, id).Error; err != nil {
		return errors.Wrapf(err, "SessionGormRepository.Delete: id %d", id)
	}
	return nil
}

func (r *SessionGormRepository) GetByUserID(userID uint) ([]*models.Sesion, error) {
	var sesiones []*models.Sesion
	if err := r.db.Where("usuario_id = ?", userID).Find(&sesiones).Error; err != nil {
		return nil, errors.Wrapf(err, "SessionGormRepository.GetByUserID: userID %d", userID)
	}
	return sesiones, nil
}

func (r *SessionGormRepository) GetByDateRange(startDate, endDate time.Time) ([]*models.Sesion, error) {
	var sesiones []*models.Sesion
	if err := r.db.Where("fecha BETWEEN ? AND ?", startDate, endDate).Find(&sesiones).Error; err != nil {
		return nil, errors.Wrapf(err, "SessionGormRepository.GetByDateRange: startDate %v, endDate %v", startDate, endDate)
	}
	return sesiones, nil
}
