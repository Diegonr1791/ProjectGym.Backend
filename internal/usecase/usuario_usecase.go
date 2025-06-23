package usecase

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UsuarioUsecase struct {
	repo repositories.UsuarioRepository
}

func NewUsuarioUsecase(repo repositories.UsuarioRepository) *UsuarioUsecase {
	return &UsuarioUsecase{repo}
}

func (uc *UsuarioUsecase) CreateUsuario(u *models.Usuario) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return domainErrors.NewAppError(500, "HASH_PASSWORD_FAILED", "Failed to hash password", err)
	}
	u.Password = string(hash)

	if err := uc.repo.Create(u); err != nil {
		if errors.Is(err, domainErrors.ErrConflict) {
			return domainErrors.ErrEmailAlreadyExists
		}
		return domainErrors.NewAppError(500, "DB_CREATE_USER_FAILED", "Failed to create user in database", err)
	}

	return nil
}

func (uc *UsuarioUsecase) GetUsuarioByID(id uint) (*models.Usuario, error) {
	user, err := uc.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, domainErrors.NewAppError(500, "DB_GET_USER_FAILED", "Failed to get user from database", err)
	}
	return user, nil
}

func (uc *UsuarioUsecase) GetUsuarioByEmail(email string) (*models.Usuario, error) {
	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, domainErrors.NewAppError(500, "DB_GET_USER_FAILED", "Failed to get user by email from database", err)
	}
	return user, nil
}

func (uc *UsuarioUsecase) GetAllUsuarios() ([]models.Usuario, error) {
	users, err := uc.repo.GetAll()
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_ALL_USERS_FAILED", "Failed to get all users from database", err)
	}
	return users, nil
}

func (uc *UsuarioUsecase) UpdateUsuario(u *models.Usuario) error {
	if _, err := uc.repo.GetByID(u.ID); err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return domainErrors.ErrNotFound
		}
		return domainErrors.NewAppError(500, "DB_UPDATE_USER_FAILED", "Failed to update user in database", err)
	}

	if u.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return domainErrors.NewAppError(500, "HASH_PASSWORD_FAILED", "Failed to hash password during update", err)
		}
		u.Password = string(hash)
	}

	if err := uc.repo.Update(u); err != nil {
		if errors.Is(err, domainErrors.ErrConflict) {
			return domainErrors.ErrEmailAlreadyExists
		}
		return domainErrors.NewAppError(500, "DB_UPDATE_USER_FAILED", "Failed to apply user updates", err)
	}

	return nil
}

func (uc *UsuarioUsecase) DeleteUsuario(id uint) error {
	if err := uc.repo.Delete(id); err != nil {
		return domainErrors.NewAppError(500, "DB_DELETE_USER_FAILED", "Failed to delete user from database", err)
	}
	return nil
}

func (uc *UsuarioUsecase) Login(email, password string) (*models.Usuario, error) {
	usuario, err := uc.repo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return nil, domainErrors.ErrInvalidCredentials
		}
		return nil, domainErrors.NewAppError(500, "DB_LOGIN_FAILED", "Database error during login", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password)); err != nil {
		return nil, domainErrors.ErrInvalidCredentials
	}

	return usuario, nil
}
