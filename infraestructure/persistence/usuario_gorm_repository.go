package persistence

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	"github.com/jackc/pgconn"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UsuarioGormRepository struct {
	DB *gorm.DB
}

func NewUsuarioGormRepository(db *gorm.DB) *UsuarioGormRepository {
	return &UsuarioGormRepository{DB: db}
}

func (r *UsuarioGormRepository) GetAll() ([]models.User, error) {
	var usuarios []models.User
	if err := r.DB.Where("is_deleted = ? AND is_active = ?", false, true).Find(&usuarios).Error; err != nil {
		return nil, errors.Wrap(err, "UsuarioGormRepository.GetAll")
	}
	return usuarios, nil
}

func (r *UsuarioGormRepository) GetAllIncludingDeleted() ([]models.User, error) {
	var usuarios []models.User
	if err := r.DB.Find(&usuarios).Error; err != nil {
		return nil, errors.Wrap(err, "UsuarioGormRepository.GetAllIncludingDeleted")
	}
	return usuarios, nil
}

func (r *UsuarioGormRepository) GetByID(id uint) (*models.User, error) {
	var usuario models.User
	if err := r.DB.Where("id = ? AND is_deleted = ? AND is_active = ?", id, false, true).First(&usuario).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "UsuarioGormRepository.GetByID: id %d", id)
	}
	return &usuario, nil
}

func (r *UsuarioGormRepository) GetByIDIncludingDeleted(id uint) (*models.User, error) {
	var usuario models.User
	if err := r.DB.First(&usuario, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "UsuarioGormRepository.GetByIDIncludingDeleted: id %d", id)
	}
	return &usuario, nil
}

func (r *UsuarioGormRepository) GetByEmail(email string) (*models.User, error) {
	var usuario models.User
	if err := r.DB.Where("email = ? AND is_deleted = ? AND is_active = ?", email, false, true).First(&usuario).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "UsuarioGormRepository.GetByEmail: email %s", email)
	}
	return &usuario, nil
}

func (r *UsuarioGormRepository) GetByEmailIncludingDeleted(email string) (*models.User, error) {
	var usuario models.User
	if err := r.DB.Where("email = ?", email).First(&usuario).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "UsuarioGormRepository.GetByEmailIncludingDeleted: email %s", email)
	}
	return &usuario, nil
}

func (r *UsuarioGormRepository) GetByEmailWithRole(email string) (*models.User, error) {
	var usuario models.User
	if err := r.DB.Preload("Role").Where("email = ? AND is_deleted = ? AND is_active = ?", email, false, true).First(&usuario).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "UsuarioGormRepository.GetByEmailWithRole: email %s", email)
	}
	return &usuario, nil
}

func (r *UsuarioGormRepository) Create(usuario *models.User) error {
	if err := r.DB.Create(usuario).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domainErrors.ErrConflict
		}
		return errors.Wrap(err, "UsuarioGormRepository.Create")
	}
	return nil
}

func (r *UsuarioGormRepository) Update(usuario *models.User) error {
	if err := r.DB.Save(usuario).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domainErrors.ErrConflict
		}
		return errors.Wrapf(err, "UsuarioGormRepository.Update: id %d", usuario.ID)
	}
	return nil
}

func (r *UsuarioGormRepository) Delete(id uint) error {
	usuario, err := r.GetByIDIncludingDeleted(id)
	if err != nil {
		return errors.Wrapf(err, "UsuarioGormRepository.Delete: failed to get user with id %d", id)
	}

	if err := usuario.SoftDelete(); err != nil {
		return errors.Wrapf(err, "UsuarioGormRepository.Delete: failed to soft delete user with id %d", id)
	}

	if err := r.DB.Save(usuario).Error; err != nil {
		return errors.Wrapf(err, "UsuarioGormRepository.Delete: failed to save soft deleted user with id %d", id)
	}

	return nil
}

func (r *UsuarioGormRepository) Restore(id uint) error {
	usuario, err := r.GetByIDIncludingDeleted(id)
	if err != nil {
		return errors.Wrapf(err, "UsuarioGormRepository.Restore: failed to get user with id %d", id)
	}

	if !usuario.IsSoftDeleted() {
		return domainErrors.NewAppError(400, "USER_NOT_DELETED", "User is not deleted", nil)
	}

	if err := usuario.Restore(); err != nil {
		return errors.Wrapf(err, "UsuarioGormRepository.Restore: failed to restore user with id %d", id)
	}

	if err := r.DB.Save(usuario).Error; err != nil {
		return errors.Wrapf(err, "UsuarioGormRepository.Restore: failed to save restored user with id %d", id)
	}

	return nil
}

func (r *UsuarioGormRepository) HardDelete(id uint) error {
	if err := r.DB.Unscoped().Delete(&models.User{}, id).Error; err != nil {
		return errors.Wrapf(err, "UsuarioGormRepository.HardDelete: id %d", id)
	}
	return nil
}

func (r *UsuarioGormRepository) GetDeletedUsers() ([]models.User, error) {
	var usuarios []models.User
	if err := r.DB.Where("is_deleted = ?", true).Find(&usuarios).Error; err != nil {
		return nil, errors.Wrap(err, "UsuarioGormRepository.GetDeletedUsers")
	}
	return usuarios, nil
}
