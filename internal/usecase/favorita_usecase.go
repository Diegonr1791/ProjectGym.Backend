package usecase

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"
)

type FavoritaUsecase struct {
	repo repository.FavoritaRepository
}

func NewFavoritaUsecase(repo repository.FavoritaRepository) *FavoritaUsecase {
	return &FavoritaUsecase{repo}
}

func (uc *FavoritaUsecase) ObtenerTodos() ([]model.Favorita, error) {
	return uc.repo.GetAll()
}

func (uc *FavoritaUsecase) ObtenerPorID(id uint) (*model.Favorita, error) {
	return uc.repo.GetByID(id)
}

func (uc *FavoritaUsecase) Crear(favorita *model.Favorita) error {
	return uc.repo.Create(favorita)
}

func (uc *FavoritaUsecase) Actualizar(favorita *model.Favorita) error {
	return uc.repo.Update(favorita)
}

func (uc *FavoritaUsecase) Eliminar(id uint) error {
	return uc.repo.Delete(id)
}

func (uc *FavoritaUsecase) ObtenerFavoritasPorUsuario(usuarioID uint) ([]model.Favorita, error) {
	return uc.repo.GetFavoritasByUsuarioID(usuarioID)
}
