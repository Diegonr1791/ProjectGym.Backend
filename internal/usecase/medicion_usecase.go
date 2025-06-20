package usecase

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"
)

type MedicionUsecase struct {
	medicionRepo repository.MedicionRepository
}

func NewMedicionUsecase(medicionRepo repository.MedicionRepository) *MedicionUsecase {
	return &MedicionUsecase{medicionRepo}
}

func (u *MedicionUsecase) GetAll() ([]model.Medicion, error) {
	return u.medicionRepo.GetAll()
}

func (u *MedicionUsecase) GetByID(id uint) (*model.Medicion, error) {
	return u.medicionRepo.GetByID(id)
}

func (u *MedicionUsecase) Create(medicion *model.Medicion) error {
	return u.medicionRepo.Create(medicion)
}

func (u *MedicionUsecase) Update(medicion *model.Medicion) error {
	return u.medicionRepo.Update(medicion)
}

func (u *MedicionUsecase) Delete(id uint) error {
	return u.medicionRepo.Delete(id)
}

func (u *MedicionUsecase) GetMesurementsByUserID(usuarioID uint) ([]model.Medicion, error) {
	return u.medicionRepo.GetMesurementsByUserID(usuarioID)
}
