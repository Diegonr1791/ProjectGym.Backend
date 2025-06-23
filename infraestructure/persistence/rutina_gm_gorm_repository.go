package persistence

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RutinaGrupoMuscularGormRepository struct {
	db *gorm.DB
}

func NewRutinaGrupoMuscularGormRepository(db *gorm.DB) repositories.RutinaGrupoMuscularRepository {
	return &RutinaGrupoMuscularGormRepository{db}
}

func (r *RutinaGrupoMuscularGormRepository) Create(rutinaGM *models.RutinaGrupoMuscular) error {
	if err := r.db.Create(rutinaGM).Error; err != nil {
		return errors.Wrap(err, "RutinaGrupoMuscularGormRepository.Create")
	}
	return nil
}

func (r *RutinaGrupoMuscularGormRepository) GetAll() ([]models.RutinaGrupoMuscular, error) {
	var rutinasGM []models.RutinaGrupoMuscular
	if err := r.db.Find(&rutinasGM).Error; err != nil {
		return nil, errors.Wrap(err, "RutinaGrupoMuscularGormRepository.GetAll")
	}
	return rutinasGM, nil
}

func (r *RutinaGrupoMuscularGormRepository) GetByID(id uint) (*models.RutinaGrupoMuscular, error) {
	var rutinaGM models.RutinaGrupoMuscular
	if err := r.db.First(&rutinaGM, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, errors.Wrapf(err, "RutinaGrupoMuscularGormRepository.GetByID: id %d", id)
	}
	return &rutinaGM, nil
}

func (r *RutinaGrupoMuscularGormRepository) Update(rutinaGM *models.RutinaGrupoMuscular) error {
	if err := r.db.Save(rutinaGM).Error; err != nil {
		return errors.Wrapf(err, "RutinaGrupoMuscularGormRepository.Update: id %d", rutinaGM.ID)
	}
	return nil
}

func (r *RutinaGrupoMuscularGormRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.RutinaGrupoMuscular{}, id).Error; err != nil {
		return errors.Wrapf(err, "RutinaGrupoMuscularGormRepository.Delete: id %d", id)
	}
	return nil
}

func (r *RutinaGrupoMuscularGormRepository) ObtenerRutinasPorGrupoMuscular(grupoMuscularID uint) ([]models.RutinaGrupoMuscular, error) {
	var rutinasGM []models.RutinaGrupoMuscular
	if err := r.db.Where("grupo_muscular_id = ?", grupoMuscularID).Find(&rutinasGM).Error; err != nil {
		return nil, errors.Wrapf(err, "RutinaGrupoMuscularGormRepository.ObtenerRutinasPorGrupoMuscular: grupoMuscularID %d", grupoMuscularID)
	}
	return rutinasGM, nil
}

// Obtener grupos musculares por rutina
func (r *RutinaGrupoMuscularGormRepository) GetMusclesGroupByRutine(id uint) (*models.RutinaConGruposMusculares, error) {
	var gruposMusculares []models.GrupoMuscular
	if err := r.db.
		Joins("JOIN rutina_grupo_muscular ON grupos_musculares.id = rutina_grupo_muscular.grupo_muscular_id").
		Where("rutina_grupo_muscular.rutina_id = ?", id).
		Find(&gruposMusculares).Error; err != nil {
		return nil, errors.Wrapf(err, "RutinaGrupoMuscularGormRepository.GetMusclesGroupByRutine: id %d", id)
	}

	response := &models.RutinaConGruposMusculares{
		ID:               id,
		RutinaID:         id,
		GruposMusculares: gruposMusculares,
	}

	return response, nil
}
