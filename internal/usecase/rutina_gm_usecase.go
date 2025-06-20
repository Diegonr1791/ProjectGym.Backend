package usecase

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"
)

type RutinaGrupoMuscularUsecase struct {
	repo repository.RutinaGrupoMuscularRepository
}

func NewRutinaGrupoMuscularUsecase(repo repository.RutinaGrupoMuscularRepository) *RutinaGrupoMuscularUsecase {
	return &RutinaGrupoMuscularUsecase{repo}
}

func (uc *RutinaGrupoMuscularUsecase) Crear(rutinaGM *model.RutinaGrupoMuscular) error {
	return uc.repo.Create(rutinaGM)
}

func (uc *RutinaGrupoMuscularUsecase) ObtenerTodos() ([]model.RutinaGrupoMuscular, error) {
	return uc.repo.GetAll()
}

func (uc *RutinaGrupoMuscularUsecase) ObtenerPorID(id uint) (*model.RutinaGrupoMuscular, error) {
	return uc.repo.GetByID(id)
}

func (uc *RutinaGrupoMuscularUsecase) Actualizar(rutinaGM *model.RutinaGrupoMuscular) error {
	return uc.repo.Update(rutinaGM)
}

func (uc *RutinaGrupoMuscularUsecase) Eliminar(id uint) error {
	return uc.repo.Delete(id)
}

func (uc *RutinaGrupoMuscularUsecase) ObtenerGruposMuscularesPorRutina(id uint) (*model.RutinaConGruposMusculares, error) {
	return uc.repo.GetMusclesGroupByRutine(id)
}
