package persistence

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"gorm.io/gorm"
)

type RutinaGrupoMuscularGormRepository struct {
	db *gorm.DB
}

func NewRutinaGrupoMuscularGormRepository(db *gorm.DB) repository.RutinaGrupoMuscularRepository {
	return &RutinaGrupoMuscularGormRepository{db}
}

func (r *RutinaGrupoMuscularGormRepository) Create(rutinaGM *model.RutinaGrupoMuscular) error {
	return r.db.Create(rutinaGM).Error
}

func (r *RutinaGrupoMuscularGormRepository) GetAll() ([]model.RutinaGrupoMuscular, error) {
	var rutinasGM []model.RutinaGrupoMuscular
	if err := r.db.Find(&rutinasGM).Error; err != nil {
		return nil, err
	}
	return rutinasGM, nil
}

func (r *RutinaGrupoMuscularGormRepository) GetByID(id uint) (*model.RutinaGrupoMuscular, error) {
	var rutinaGM model.RutinaGrupoMuscular
	if err := r.db.First(&rutinaGM, id).Error; err != nil {
		return nil, err
	}
	return &rutinaGM, nil
}

func (r *RutinaGrupoMuscularGormRepository) Update(rutinaGM *model.RutinaGrupoMuscular) error {
	return r.db.Save(rutinaGM).Error
}

func (r *RutinaGrupoMuscularGormRepository) Delete(id uint) error {
	return r.db.Delete(&model.RutinaGrupoMuscular{}, id).Error
}

func (r *RutinaGrupoMuscularGormRepository) ObtenerRutinasPorGrupoMuscular(grupoMuscularID uint) ([]model.RutinaGrupoMuscular, error) {
	var rutinasGM []model.RutinaGrupoMuscular
	if err := r.db.Where("grupo_muscular_id = ?", grupoMuscularID).Find(&rutinasGM).Error; err != nil {
		return nil, err
	}
	return rutinasGM, nil
}

// Obtener grupos musculares por rutina
func (r *RutinaGrupoMuscularGormRepository) GetMusclesGroupByRutine(id uint) (*model.RutinaConGruposMusculares, error) {
	var gruposMusculares []model.GrupoMuscular
	if err := r.db.
		Joins("JOIN rutina_grupo_muscular ON grupos_musculares.id = rutina_grupo_muscular.grupo_muscular_id").
		Where("rutina_grupo_muscular.rutina_id = ?", id).
		Find(&gruposMusculares).Error; err != nil {
		return nil, err
	}

	response := &model.RutinaConGruposMusculares{
		ID:               id,
		RutinaID:         id,
		GruposMusculares: gruposMusculares,
	}

	return response, nil
}
