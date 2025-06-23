package usecase

import (
	"regexp"
	"strings"
	"time"

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

// validateEmail validates email format
func (uc *UsuarioUsecase) validateEmail(email string) error {
	if email == "" {
		return domainErrors.NewAppError(400, "EMAIL_REQUIRED", "Email is required", nil)
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return domainErrors.NewAppError(400, "INVALID_EMAIL_FORMAT", "Invalid email format", nil)
	}

	return nil
}

// validatePassword validates password strength
func (uc *UsuarioUsecase) validatePassword(password string) error {
	if password == "" {
		return domainErrors.NewAppError(400, "PASSWORD_REQUIRED", "Password is required", nil)
	}

	if len(password) < 8 {
		return domainErrors.NewAppError(400, "PASSWORD_TOO_SHORT", "Password must be at least 8 characters long", nil)
	}

	if len(password) > 128 {
		return domainErrors.NewAppError(400, "PASSWORD_TOO_LONG", "Password must not exceed 128 characters", nil)
	}

	// Check for at least one uppercase letter, one lowercase letter, and one number
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

	if !hasUpper || !hasLower || !hasNumber {
		return domainErrors.NewAppError(400, "PASSWORD_WEAK", "Password must contain at least one uppercase letter, one lowercase letter, and one number", nil)
	}

	return nil
}

// validateName validates user name
func (uc *UsuarioUsecase) validateName(name string) error {
	if name == "" {
		return domainErrors.NewAppError(400, "NAME_REQUIRED", "Name is required", nil)
	}

	name = strings.TrimSpace(name)
	if len(name) < 2 {
		return domainErrors.NewAppError(400, "NAME_TOO_SHORT", "Name must be at least 2 characters long", nil)
	}

	if len(name) > 100 {
		return domainErrors.NewAppError(400, "NAME_TOO_LONG", "Name must not exceed 100 characters", nil)
	}

	// Check for valid characters (letters, spaces, hyphens, apostrophes)
	nameRegex := regexp.MustCompile(`^[a-zA-ZÀ-ÿ\s\-']+$`)
	if !nameRegex.MatchString(name) {
		return domainErrors.NewAppError(400, "INVALID_NAME_CHARACTERS", "Name contains invalid characters", nil)
	}

	return nil
}

// validateUser validates all user fields
func (uc *UsuarioUsecase) validateUser(u *models.User, isUpdate bool) error {
	if !isUpdate {
		// Validations for new user creation
		if err := uc.validateName(u.Name); err != nil {
			return err
		}

		if err := uc.validateEmail(u.Email); err != nil {
			return err
		}

		if err := uc.validatePassword(u.Password); err != nil {
			return err
		}

		if u.RoleID == 0 {
			return domainErrors.NewAppError(400, "ROLE_REQUIRED", "Role is required", nil)
		}
	} else {
		// Validations for user update
		if u.Name != "" {
			if err := uc.validateName(u.Name); err != nil {
				return err
			}
		}

		if u.Email != "" {
			if err := uc.validateEmail(u.Email); err != nil {
				return err
			}
		}

		if u.Password != "" {
			if err := uc.validatePassword(u.Password); err != nil {
				return err
			}
		}
	}

	return nil
}

func (uc *UsuarioUsecase) CreateUsuario(u *models.User) error {
	// Validate user data
	if err := uc.validateUser(u, false); err != nil {
		return err
	}

	// Check if email already exists
	existingUser, err := uc.repo.GetByEmailIncludingDeleted(u.Email)
	if err != nil && !errors.Is(err, domainErrors.ErrNotFound) {
		return domainErrors.NewAppError(500, "DB_CHECK_EMAIL_FAILED", "Failed to check email existence", err)
	}

	if existingUser != nil {
		if existingUser.IsSoftDeleted() {
			return domainErrors.NewAppError(409, "EMAIL_SOFT_DELETED", "Email belongs to a deleted user. Please contact support to restore the account.", nil)
		}
		return domainErrors.ErrEmailAlreadyExists
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return domainErrors.NewAppError(500, "HASH_PASSWORD_FAILED", "Failed to hash password", err)
	}
	u.Password = string(hash)

	// Set default values
	u.IsActive = true
	u.IsDeleted = false
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	if err := uc.repo.Create(u); err != nil {
		if errors.Is(err, domainErrors.ErrConflict) {
			return domainErrors.ErrEmailAlreadyExists
		}
		return domainErrors.NewAppError(500, "DB_CREATE_USER_FAILED", "Failed to create user in database", err)
	}

	return nil
}

func (uc *UsuarioUsecase) GetUsuarioByID(id uint) (*models.User, error) {
	user, err := uc.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, domainErrors.NewAppError(500, "DB_GET_USER_FAILED", "Failed to get user from database", err)
	}
	return user, nil
}

func (uc *UsuarioUsecase) GetUsuarioByEmail(email string) (*models.User, error) {
	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, domainErrors.NewAppError(500, "DB_GET_USER_FAILED", "Failed to get user by email from database", err)
	}
	return user, nil
}

func (uc *UsuarioUsecase) GetAllUsuarios() ([]models.User, error) {
	users, err := uc.repo.GetAll()
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_ALL_USERS_FAILED", "Failed to get all users from database", err)
	}
	return users, nil
}

func (uc *UsuarioUsecase) GetAllUsuariosIncludingDeleted() ([]models.User, error) {
	users, err := uc.repo.GetAllIncludingDeleted()
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_ALL_USERS_FAILED", "Failed to get all users from database", err)
	}
	return users, nil
}

func (uc *UsuarioUsecase) GetDeletedUsuarios() ([]models.User, error) {
	users, err := uc.repo.GetDeletedUsers()
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_DELETED_USERS_FAILED", "Failed to get deleted users from database", err)
	}
	return users, nil
}

func (uc *UsuarioUsecase) UpdateUsuario(u *models.User) error {
	// Check if user exists
	existingUser, err := uc.repo.GetByID(u.ID)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return domainErrors.ErrNotFound
		}
		return domainErrors.NewAppError(500, "DB_GET_USER_FAILED", "Failed to get user for update", err)
	}

	// Validate user data
	if err := uc.validateUser(u, true); err != nil {
		return err
	}

	// Check email uniqueness if email is being updated
	if u.Email != "" && u.Email != existingUser.Email {
		emailUser, err := uc.repo.GetByEmailIncludingDeleted(u.Email)
		if err != nil && !errors.Is(err, domainErrors.ErrNotFound) {
			return domainErrors.NewAppError(500, "DB_CHECK_EMAIL_FAILED", "Failed to check email existence", err)
		}

		if emailUser != nil {
			if emailUser.IsSoftDeleted() {
				return domainErrors.NewAppError(409, "EMAIL_SOFT_DELETED", "Email belongs to a deleted user", nil)
			}
			return domainErrors.ErrEmailAlreadyExists
		}
	}

	// Hash password if provided
	if u.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return domainErrors.NewAppError(500, "HASH_PASSWORD_FAILED", "Failed to hash password during update", err)
		}
		u.Password = string(hash)
	} else {
		// Keep existing password if not provided
		u.Password = existingUser.Password
	}

	// Preserve existing values if not provided
	if u.Name == "" {
		u.Name = existingUser.Name
	}
	if u.Email == "" {
		u.Email = existingUser.Email
	}
	if u.RoleID == 0 {
		u.RoleID = existingUser.RoleID
	}

	if err := uc.repo.Update(u); err != nil {
		if errors.Is(err, domainErrors.ErrConflict) {
			return domainErrors.ErrEmailAlreadyExists
		}
		return domainErrors.NewAppError(500, "DB_UPDATE_USER_FAILED", "Failed to apply user updates", err)
	}

	//Default values
	u.UpdatedAt = time.Now()

	return nil
}

func (uc *UsuarioUsecase) DeleteUsuario(id uint) error {
	// Check if user exists and is not already deleted
	user, err := uc.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return domainErrors.ErrNotFound
		}
		return domainErrors.NewAppError(500, "DB_GET_USER_FAILED", "Failed to get user for deletion", err)
	}

	if user.IsSoftDeleted() {
		return domainErrors.NewAppError(400, "USER_ALREADY_DELETED", "User is already deleted", nil)
	}

	if err := uc.repo.Delete(id); err != nil {
		return domainErrors.NewAppError(500, "DB_DELETE_USER_FAILED", "Failed to delete user from database", err)
	}
	return nil
}

func (uc *UsuarioUsecase) RestoreUsuario(id uint) error {
	// Check if user exists (including deleted ones)
	user, err := uc.repo.GetByIDIncludingDeleted(id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return domainErrors.ErrNotFound
		}
		return domainErrors.NewAppError(500, "DB_GET_USER_FAILED", "Failed to get user for restoration", err)
	}

	if !user.IsSoftDeleted() {
		return domainErrors.NewAppError(400, "USER_NOT_DELETED", "User is not deleted", nil)
	}

	if err := uc.repo.Restore(id); err != nil {
		return domainErrors.NewAppError(500, "DB_RESTORE_USER_FAILED", "Failed to restore user", err)
	}
	return nil
}

func (uc *UsuarioUsecase) HardDeleteUsuario(id uint) error {
	// Check if user exists (including deleted ones)
	_, err := uc.repo.GetByIDIncludingDeleted(id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return domainErrors.ErrNotFound
		}
		return domainErrors.NewAppError(500, "DB_GET_USER_FAILED", "Failed to get user for hard deletion", err)
	}

	if err := uc.repo.HardDelete(id); err != nil {
		return domainErrors.NewAppError(500, "DB_HARD_DELETE_USER_FAILED", "Failed to permanently delete user", err)
	}
	return nil
}

func (uc *UsuarioUsecase) Login(email, password string) (*models.User, error) {
	// Validate email format
	if err := uc.validateEmail(email); err != nil {
		return nil, err
	}

	// Validate password is not empty
	if password == "" {
		return nil, domainErrors.NewAppError(400, "PASSWORD_REQUIRED", "Password is required", nil)
	}

	usuario, err := uc.repo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return nil, domainErrors.ErrInvalidCredentials
		}
		return nil, domainErrors.NewAppError(500, "DB_LOGIN_FAILED", "Database error during login", err)
	}

	// Check if user is active
	if !usuario.IsActive {
		return nil, domainErrors.NewAppError(401, "USER_INACTIVE", "User account is inactive", nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password)); err != nil {
		return nil, domainErrors.ErrInvalidCredentials
	}

	return usuario, nil
}
