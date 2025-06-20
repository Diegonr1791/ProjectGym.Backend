package repository

import model "github.com/Diegonr1791/GymBro/internal/domain/models"

type UsuarioRepository interface {
	GetAll() ([]model.Usuario, error)
	GetByID(id uint) (*model.Usuario, error)
	GetByEmail(email string) (*model.Usuario, error)
	Create(usuario *model.Usuario) error
	Update(usuario *model.Usuario) error
	Delete(id uint) error
}
