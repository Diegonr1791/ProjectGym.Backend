package usecase

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"
)

type ExerciseUsecase struct {
	exerciseRepo repository.ExerciseRepository
}

func NewExerciseUsecase(exerciseRepo repository.ExerciseRepository) *ExerciseUsecase {
	return &ExerciseUsecase{exerciseRepo}
}

func (u *ExerciseUsecase) GetAll() ([]*model.Ejercicio, error) {
	return u.exerciseRepo.GetAll()
}

func (u *ExerciseUsecase) GetById(id uint) (*model.Ejercicio, error) {
	return u.exerciseRepo.GetById(id)
}

func (u *ExerciseUsecase) Create(exercise *model.Ejercicio) error {
	return u.exerciseRepo.Create(exercise)
}

func (u *ExerciseUsecase) Update(exercise *model.Ejercicio) error {
	return u.exerciseRepo.Update(exercise)
}

func (u *ExerciseUsecase) Delete(id uint) error {
	return u.exerciseRepo.Delete(id)
}

func (u *ExerciseUsecase) GetByMuscleGroup(muscleGroupID uint) ([]*model.Ejercicio, error) {
	return u.exerciseRepo.GetByMuscleGroup(muscleGroupID)
}
