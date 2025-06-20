package usecase

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UsuarioUsecase struct {
	repo repository.UsuarioRepository
}

func NewUsuarioUsecase(repo repository.UsuarioRepository) *UsuarioUsecase {
	return &UsuarioUsecase{repo}
}

func (s *UsuarioUsecase) RegistrarUsuario(u *model.Usuario) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return s.repo.Create(u)
}

func (s *UsuarioUsecase) ObtenerUsuario(id uint) (*model.Usuario, error) {
	return s.repo.GetByID(id)
}

func (s *UsuarioUsecase) ObtenerUsuarioPorEmail(email string) (*model.Usuario, error) {
	return s.repo.GetByEmail(email)
}

func (s *UsuarioUsecase) ObtenerTodosUsuarios() ([]model.Usuario, error) {
	return s.repo.GetAll()
}

func (s *UsuarioUsecase) ActualizarUsuario(u *model.Usuario) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return s.repo.Update(u)
}

func (s *UsuarioUsecase) EliminarUsuario(id uint) error {
	return s.repo.Delete(id)
}

func (s *UsuarioUsecase) Login(email, password string) (*model.Usuario, error) {
	usuario, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password)); err != nil {
		return nil, err
	}

	return usuario, nil
}
