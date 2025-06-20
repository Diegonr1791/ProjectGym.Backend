package usecase

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"
)

type GrupoMuscularUseCase struct {
	repo repository.GrupoMuscularRepository
}

func NewGrupoMuscularUseCase(repo repository.GrupoMuscularRepository) *GrupoMuscularUseCase {
	return &GrupoMuscularUseCase{repo}
}

func (uc *GrupoMuscularUseCase) Crear(g *model.GrupoMuscular) error {
	return uc.repo.Create(g)
}

func (uc *GrupoMuscularUseCase) ObtenerTodos() ([]model.GrupoMuscular, error) {
	return uc.repo.GetAll()
}

func (uc *GrupoMuscularUseCase) ObtenerPorID(id uint) (*model.GrupoMuscular, error) {
	return uc.repo.GetByID(id)
}

func (uc *GrupoMuscularUseCase) Actualizar(g *model.GrupoMuscular) error {
	return uc.repo.Update(g)
}

func (uc *GrupoMuscularUseCase) Eliminar(id uint) error {
	return uc.repo.Delete(id)
}
