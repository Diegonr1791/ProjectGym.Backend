package usecase

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"
)

type RutinaService struct {
	repo repository.RutinaRepository
}

func NewRutinaUseCase(repo repository.RutinaRepository) *RutinaService {
	return &RutinaService{repo}
}

func (s *RutinaService) Obtener() ([]model.Rutina, error) {
	return s.repo.GetAll()
}

func (s *RutinaService) Crear(rutina *model.Rutina) error {
	return s.repo.Create(rutina)
}

func (s *RutinaService) ObtenerPorID(id uint) (*model.Rutina, error) {
	return s.repo.GetByID(id)
}

func (s *RutinaService) Actualizar(rutina *model.Rutina) error {
	return s.repo.Update(rutina)
}

func (s *RutinaService) Eliminar(id uint) error {
	return s.repo.Delete(id)
}
