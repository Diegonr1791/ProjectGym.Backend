package usecase

import (
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
)

type RoutineMuscleGroupUsecase struct {
	repo repositories.RutinaGrupoMuscularRepository
}

func NewRoutineMuscleGroupUsecase(repo repositories.RutinaGrupoMuscularRepository) *RoutineMuscleGroupUsecase {
	return &RoutineMuscleGroupUsecase{repo}
}

func (uc *RoutineMuscleGroupUsecase) Create(rutinaGM *models.RutinaGrupoMuscular) error {
	if err := uc.repo.Create(rutinaGM); err != nil {
		return domainErrors.NewAppError(500, "DB_CREATE_ROUTINE_MUSCLE_GROUP_FAILED", "Failed to create routine muscle group in database", err)
	}
	return nil
}

func (uc *RoutineMuscleGroupUsecase) GetAll() ([]models.RutinaGrupoMuscular, error) {
	rutinasGM, err := uc.repo.GetAll()
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_ALL_ROUTINE_MUSCLE_GROUPS_FAILED", "Failed to get all routine muscle groups from database", err)
	}
	return rutinasGM, nil
}

func (uc *RoutineMuscleGroupUsecase) GetByID(id uint) (*models.RutinaGrupoMuscular, error) {
	rutinaGM, err := uc.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, domainErrors.NewAppError(500, "DB_GET_ROUTINE_MUSCLE_GROUP_FAILED", "Failed to get routine muscle group from database", err)
	}
	return rutinaGM, nil
}

func (uc *RoutineMuscleGroupUsecase) Update(rutinaGM *models.RutinaGrupoMuscular) error {
	if _, err := uc.repo.GetByID(rutinaGM.ID); err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return domainErrors.ErrNotFound
		}
		return domainErrors.NewAppError(500, "DB_UPDATE_ROUTINE_MUSCLE_GROUP_FAILED", "Failed to verify routine muscle group existence", err)
	}
	if err := uc.repo.Update(rutinaGM); err != nil {
		return domainErrors.NewAppError(500, "DB_UPDATE_ROUTINE_MUSCLE_GROUP_FAILED", "Failed to update routine muscle group in database", err)
	}
	return nil
}

func (uc *RoutineMuscleGroupUsecase) Delete(id uint) error {
	if err := uc.repo.Delete(id); err != nil {
		return domainErrors.NewAppError(500, "DB_DELETE_ROUTINE_MUSCLE_GROUP_FAILED", "Failed to delete routine muscle group from database", err)
	}
	return nil
}

func (uc *RoutineMuscleGroupUsecase) GetMuscleGroupsByRoutine(id uint) (*models.RutinaConGruposMusculares, error) {
	grupos, err := uc.repo.GetMusclesGroupByRutine(id)
	if err != nil {
		return nil, domainErrors.NewAppError(500, "DB_GET_MUSCLE_GROUPS_BY_ROUTINE_FAILED", "Failed to get muscle groups by routine from database", err)
	}
	return grupos, nil
}
