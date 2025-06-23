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

func (r *UsuarioGormRepository) GetAll() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	if err := r.DB.Find(&usuarios).Error; err != nil {
		return nil, errors.Wrap(err, "UsuarioGormRepository.GetAll")
	}
	return usuarios, nil
}

func (r *UsuarioGormRepository) GetByID(id uint) (*models.Usuario, error) {
	var usuario models.Usuario
	if err := r.DB.First(&usuario, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "UsuarioGormRepository.GetByID: id %d", id)
	}
	return &usuario, nil
}

func (r *UsuarioGormRepository) GetByEmail(email string) (*models.Usuario, error) {
	var usuario models.Usuario
	if err := r.DB.Where("email = ?", email).First(&usuario).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "UsuarioGormRepository.GetByEmail: email %s", email)
	}
	return &usuario, nil
}

func (r *UsuarioGormRepository) Create(usuario *models.Usuario) error {
	if err := r.DB.Create(usuario).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domainErrors.ErrConflict
		}
		return errors.Wrap(err, "UsuarioGormRepository.Create")
	}
	return nil
}

func (r *UsuarioGormRepository) Update(usuario *models.Usuario) error {
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
	if err := r.DB.Delete(&models.Usuario{}, id).Error; err != nil {
		return errors.Wrapf(err, "UsuarioGormRepository.Delete: id %d", id)
	}
	return nil
}
