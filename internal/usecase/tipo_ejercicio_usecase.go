package usecase

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"
)

type TypeExerciseUsecase struct {
	typeExerciseRepo repository.TypeExerciseRepository
}

func NewTypeExerciseUsecase(typeExerciseRepo repository.TypeExerciseRepository) *TypeExerciseUsecase {
	return &TypeExerciseUsecase{typeExerciseRepo}
}

func (u *TypeExerciseUsecase) GetAll() ([]*model.TipoEjercicio, error) {
	return u.typeExerciseRepo.GetAll()
}

func (u *TypeExerciseUsecase) GetById(id uint) (*model.TipoEjercicio, error) {
	return u.typeExerciseRepo.GetById(id)
}

func (u *TypeExerciseUsecase) Create(tipoEjercicio *model.TipoEjercicio) error {
	return u.typeExerciseRepo.Create(tipoEjercicio)
}

func (u *TypeExerciseUsecase) Update(tipoEjercicio *model.TipoEjercicio) error {
	return u.typeExerciseRepo.Update(tipoEjercicio)
}

func (u *TypeExerciseUsecase) Delete(id uint) error {
	return u.typeExerciseRepo.Delete(id)
}
