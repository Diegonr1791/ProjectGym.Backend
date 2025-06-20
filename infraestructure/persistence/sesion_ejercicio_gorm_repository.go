package persistence

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"

	"time"

	"gorm.io/gorm"
)

type SessionExerciseGormRepository struct {
	db *gorm.DB
}

func NewSessionExerciseGormRepository(db *gorm.DB) repository.SessionExerciseRepository {
	return &SessionExerciseGormRepository{db}
}

func (r *SessionExerciseGormRepository) Create(sesionEjercicio *model.SesionEjercicio) error {
	return r.db.Create(sesionEjercicio).Error
}

func (r *SessionExerciseGormRepository) GetAll() ([]*model.SesionEjercicio, error) {
	var sesionesEjercicios []*model.SesionEjercicio
	if err := r.db.Find(&sesionesEjercicios).Error; err != nil {
		return nil, err
	}
	return sesionesEjercicios, nil
}

func (r *SessionExerciseGormRepository) GetById(id uint) (*model.SesionEjercicio, error) {
	var sesionEjercicio model.SesionEjercicio
	if err := r.db.First(&sesionEjercicio, id).Error; err != nil {
		return nil, err
	}
	return &sesionEjercicio, nil
}

func (r *SessionExerciseGormRepository) Update(sesionEjercicio *model.SesionEjercicio) error {
	return r.db.Save(sesionEjercicio).Error
}

func (r *SessionExerciseGormRepository) Delete(id uint) error {
	return r.db.Delete(&model.SesionEjercicio{}, id).Error
}

func (r *SessionExerciseGormRepository) GetBySessionID(sessionID uint, fechaDesde, fechaHasta time.Time) ([]*model.SesionEjercicio, error) {
	var sesionesEjercicios []*model.SesionEjercicio
	if err := r.db.
		Joins("JOIN sesiones ON sesion_ejercicio.sesion_id = sesiones.id").
		Where("sesion_ejercicio.sesion_id = ? AND sesiones.fecha BETWEEN ? AND ?", sessionID, fechaDesde, fechaHasta).
		Find(&sesionesEjercicios).Error; err != nil {
		return nil, err
	}
	return sesionesEjercicios, nil
}
