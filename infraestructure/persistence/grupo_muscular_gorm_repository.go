package persistence

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"gorm.io/gorm"
)

type grupoMuscularGormRepository struct {
	db *gorm.DB
}

func NewGrupoMuscularGormRepository(db *gorm.DB) repository.GrupoMuscularRepository {
	return &grupoMuscularGormRepository{db}
}

func (r *grupoMuscularGormRepository) Create(g *model.GrupoMuscular) error {
	return r.db.Create(g).Error
}

func (r *grupoMuscularGormRepository) GetAll() ([]model.GrupoMuscular, error) {
	var grupos []model.GrupoMuscular
	err := r.db.Order("id ASC").Find(&grupos).Error
	return grupos, err
}

func (r *grupoMuscularGormRepository) GetByID(id uint) (*model.GrupoMuscular, error) {
	var g model.GrupoMuscular
	err := r.db.First(&g, id).Error
	return &g, err
}

func (r *grupoMuscularGormRepository) Update(g *model.GrupoMuscular) error {
	return r.db.Save(g).Error
}

func (r *grupoMuscularGormRepository) Delete(id uint) error {
	return r.db.Delete(&model.GrupoMuscular{}, id).Error
}
