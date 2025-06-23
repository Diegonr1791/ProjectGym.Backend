package repository

import model "github.com/Diegonr1791/GymBro/internal/domain/models"

type UsuarioRepository interface {
	GetAll() ([]model.User, error)
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(usuario *model.User) error
	Update(usuario *model.User) error
	Delete(id uint) error
}
