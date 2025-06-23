package persistence

import (
	"time"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SessionExerciseGormRepository struct {
	db *gorm.DB
}

func NewSessionExerciseGormRepository(db *gorm.DB) repositories.SessionExerciseRepository {
	return &SessionExerciseGormRepository{db}
}

func (r *SessionExerciseGormRepository) Create(sesionEjercicio *models.SesionEjercicio) error {
	if err := r.db.Create(sesionEjercicio).Error; err != nil {
		return errors.Wrap(err, "SessionExerciseGormRepository.Create")
	}
	return nil
}

func (r *SessionExerciseGormRepository) GetAll() ([]*models.SesionEjercicio, error) {
	var sesionesEjercicios []*models.SesionEjercicio
	if err := r.db.Find(&sesionesEjercicios).Error; err != nil {
		return nil, errors.Wrap(err, "SessionExerciseGormRepository.GetAll")
	}
	return sesionesEjercicios, nil
}

func (r *SessionExerciseGormRepository) GetById(id uint) (*models.SesionEjercicio, error) {
	var sesionEjercicio models.SesionEjercicio
	if err := r.db.First(&sesionEjercicio, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "SessionExerciseGormRepository.GetById: id %d", id)
	}
	return &sesionEjercicio, nil
}

func (r *SessionExerciseGormRepository) Update(sesionEjercicio *models.SesionEjercicio) error {
	if err := r.db.Save(sesionEjercicio).Error; err != nil {
		return errors.Wrapf(err, "SessionExerciseGormRepository.Update: id %d", sesionEjercicio.ID)
	}
	return nil
}

func (r *SessionExerciseGormRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.SesionEjercicio{}, id).Error; err != nil {
		return errors.Wrapf(err, "SessionExerciseGormRepository.Delete: id %d", id)
	}
	return nil
}

func (r *SessionExerciseGormRepository) GetBySessionID(sessionID uint, fechaDesde, fechaHasta time.Time) ([]*models.SesionEjercicio, error) {
	var sesionesEjercicios []*models.SesionEjercicio
	if err := r.db.
		Joins("JOIN sesiones ON sesion_ejercicio.sesion_id = sesiones.id").
		Where("sesion_ejercicio.sesion_id = ? AND sesiones.fecha BETWEEN ? AND ?", sessionID, fechaDesde, fechaHasta).
		Find(&sesionesEjercicios).Error; err != nil {
		return nil, errors.Wrapf(err, "SessionExerciseGormRepository.GetBySessionID: sessionID %d, fechaDesde %v, fechaHasta %v", sessionID, fechaDesde, fechaHasta)
	}
	return sesionesEjercicios, nil
}
