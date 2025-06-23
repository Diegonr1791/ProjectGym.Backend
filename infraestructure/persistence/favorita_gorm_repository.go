package persistence

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type FavoritaGormRepository struct {
	db *gorm.DB
}

func NewFavoritaGormRepository(db *gorm.DB) repositories.FavoritaRepository {
	return &FavoritaGormRepository{db}
}

func (r *FavoritaGormRepository) GetAll() ([]models.Favorita, error) {
	var favoritas []models.Favorita
	if err := r.db.Find(&favoritas).Error; err != nil {
		return nil, errors.Wrap(err, "FavoritaGormRepository.GetAll")
	}
	return favoritas, nil
}

func (r *FavoritaGormRepository) GetByID(id uint) (*models.Favorita, error) {
	var favorita models.Favorita
	if err := r.db.First(&favorita, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "FavoritaGormRepository.GetByID: id %d", id)
	}
	return &favorita, nil
}

func (r *FavoritaGormRepository) Create(favorita *models.Favorita) error {
	if err := r.db.Create(favorita).Error; err != nil {
		return errors.Wrap(err, "FavoritaGormRepository.Create")
	}
	return nil
}

func (r *FavoritaGormRepository) Update(favorita *models.Favorita) error {
	if err := r.db.Save(favorita).Error; err != nil {
		return errors.Wrapf(err, "FavoritaGormRepository.Update: id %d", favorita.ID)
	}
	return nil
}

func (r *FavoritaGormRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Favorita{}, id).Error; err != nil {
		return errors.Wrapf(err, "FavoritaGormRepository.Delete: id %d", id)
	}
	return nil
}

func (r *FavoritaGormRepository) GetFavoritasByUsuarioID(usuarioID uint) ([]models.Favorita, error) {
	var favoritas []models.Favorita
	if err := r.db.Where("usuario_id = ?", usuarioID).Find(&favoritas).Error; err != nil {
		return nil, errors.Wrapf(err, "FavoritaGormRepository.GetFavoritasByUsuarioID: usuarioID %d", usuarioID)
	}
	return favoritas, nil
}
