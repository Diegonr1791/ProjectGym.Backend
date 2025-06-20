package persistence

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	"gorm.io/gorm"
)

type FavoritaGormRepository struct {
	db *gorm.DB
}

func NewFavoritaGormRepository(db *gorm.DB) *FavoritaGormRepository {
	return &FavoritaGormRepository{db}
}

func (r *FavoritaGormRepository) GetAll() ([]model.Favorita, error) {
	var favoritas []model.Favorita
	if err := r.db.Find(&favoritas).Error; err != nil {
		return nil, err
	}
	return favoritas, nil
}

func (r *FavoritaGormRepository) GetByID(id uint) (*model.Favorita, error) {
	var favorita model.Favorita
	if err := r.db.First(&favorita, id).Error; err != nil {
		return nil, err
	}
	return &favorita, nil
}

func (r *FavoritaGormRepository) Create(favorita *model.Favorita) error {
	return r.db.Create(favorita).Error
}

func (r *FavoritaGormRepository) Update(favorita *model.Favorita) error {
	return r.db.Save(favorita).Error
}

func (r *FavoritaGormRepository) Delete(id uint) error {
	return r.db.Delete(&model.Favorita{}, id).Error
}

func (r *FavoritaGormRepository) GetFavoritasByUsuarioID(usuarioID uint) ([]model.Favorita, error) {
	var favoritas []model.Favorita
	if err := r.db.Where("usuario_id = ?", usuarioID).Find(&favoritas).Error; err != nil {
		return nil, err
	}
	return favoritas, nil
}
