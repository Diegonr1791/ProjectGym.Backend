package repository

import model "github.com/Diegonr1791/GymBro/internal/domain/models"

type FavoritaRepository interface {
	GetAll() ([]model.Favorita, error)
	GetByID(id uint) (*model.Favorita, error)
	Create(favorita *model.Favorita) error
	Update(favorita *model.Favorita) error
	Delete(id uint) error
	GetFavoritasByUsuarioID(usuarioID uint) ([]model.Favorita, error)
}
