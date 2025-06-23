package persistence

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type GrupoMuscularGormRepository struct {
	db *gorm.DB
}

func NewGrupoMuscularGormRepository(db *gorm.DB) repositories.GrupoMuscularRepository {
	return &GrupoMuscularGormRepository{db}
}

func (r *GrupoMuscularGormRepository) Create(g *models.GrupoMuscular) error {
	if err := r.db.Create(g).Error; err != nil {
		return errors.Wrap(err, "GrupoMuscularGormRepository.Create")
	}
	return nil
}

func (r *GrupoMuscularGormRepository) GetAll() ([]models.GrupoMuscular, error) {
	var grupos []models.GrupoMuscular
	if err := r.db.Order("id ASC").Find(&grupos).Error; err != nil {
		return nil, errors.Wrap(err, "GrupoMuscularGormRepository.GetAll")
	}
	return grupos, nil
}

func (r *GrupoMuscularGormRepository) GetByID(id uint) (*models.GrupoMuscular, error) {
	var g models.GrupoMuscular
	if err := r.db.First(&g, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "GrupoMuscularGormRepository.GetByID: id %d", id)
	}
	return &g, nil
}

func (r *GrupoMuscularGormRepository) Update(g *models.GrupoMuscular) error {
	if err := r.db.Save(g).Error; err != nil {
		return errors.Wrapf(err, "GrupoMuscularGormRepository.Update: id %d", g.ID)
	}
	return nil
}

func (r *GrupoMuscularGormRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.GrupoMuscular{}, id).Error; err != nil {
		return errors.Wrapf(err, "GrupoMuscularGormRepository.Delete: id %d", id)
	}
	return nil
}
