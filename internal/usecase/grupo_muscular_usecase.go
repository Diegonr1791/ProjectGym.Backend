package usecase

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
)

type GrupoMuscularUseCase struct {
	repo repositories.GrupoMuscularRepository
}

func NewGrupoMuscularUseCase(repo repositories.GrupoMuscularRepository) *GrupoMuscularUseCase {
	return &GrupoMuscularUseCase{repo}
}

func (uc *GrupoMuscularUseCase) CreateMuscleGroup(g *models.GrupoMuscular) error {
	if err := uc.repo.Create(g); err != nil {
		return domainErrors.NewAppError(500, "DB_CREATE_MUSCLE_GROUP_FAILED", "Failed to create muscle group in database", err)
	}
	return nil
}

func (uc *GrupoMuscularUseCase) GetAllMuscleGroups() ([]models.GrupoMuscular, error) {
	grupos, err := uc.repo.GetAll()
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_ALL_MUSCLE_GROUPS_FAILED", "Failed to get all muscle groups from database", err)
	}
	return grupos, nil
}

func (uc *GrupoMuscularUseCase) GetMuscleGroupByID(id uint) (*models.GrupoMuscular, error) {
	grupo, err := uc.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, domainErrors.NewAppError(500, "DB_GET_MUSCLE_GROUP_FAILED", "Failed to get muscle group from database", err)
	}
	return grupo, nil
}

func (uc *GrupoMuscularUseCase) UpdateMuscleGroup(g *models.GrupoMuscular) error {
	// Verify muscle group exists before updating
	if _, err := uc.repo.GetByID(g.ID); err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return domainErrors.ErrNotFound
		}
		return domainErrors.NewAppError(500, "DB_UPDATE_MUSCLE_GROUP_FAILED", "Failed to verify muscle group existence", err)
	}

	if err := uc.repo.Update(g); err != nil {
		return domainErrors.NewAppError(500, "DB_UPDATE_MUSCLE_GROUP_FAILED", "Failed to update muscle group in database", err)
	}
	return nil
}

func (uc *GrupoMuscularUseCase) DeleteMuscleGroup(id uint) error {
	if err := uc.repo.Delete(id); err != nil {
		return domainErrors.NewAppError(500, "DB_DELETE_MUSCLE_GROUP_FAILED", "Failed to delete muscle group from database", err)
	}
	return nil
}
