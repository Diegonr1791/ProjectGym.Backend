package http

import (
	"net/http"
	"strconv"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	"github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type RoutineMuscleGroupHandler struct {
	usecase *usecase.RoutineMuscleGroupUsecase
}

func NewRoutineMuscleGroupHandler(router gin.IRouter, uc *usecase.RoutineMuscleGroupUsecase) {
	handler := &RoutineMuscleGroupHandler{uc}

	routineMuscleGroupRoutes := router.Group("/routine-muscle-groups")
	{
		routineMuscleGroupRoutes.POST("", handler.Create)
		routineMuscleGroupRoutes.GET("", handler.GetAll)
		routineMuscleGroupRoutes.GET("/:id", handler.GetByID)
		routineMuscleGroupRoutes.PUT("/:id", handler.Update)
		routineMuscleGroupRoutes.DELETE("/:id", handler.Delete)
		routineMuscleGroupRoutes.GET("/routine/:id/muscle-groups", handler.GetMuscleGroupsByRoutine)
	}
}

// @Summary      Create a new routine muscle group
// @Description  Create a new routine muscle group in the system
// @Tags         routine-muscle-groups
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        routine_muscle_group body models.RutinaGrupoMuscular true "Routine muscle group data"
// @Success      201  {object}  models.RutinaGrupoMuscular
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /routine-muscle-groups [post]
func (h *RoutineMuscleGroupHandler) Create(c *gin.Context) {
	var rutinaGM models.RutinaGrupoMuscular
	if err := c.ShouldBindJSON(&rutinaGM); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	if err := h.usecase.Create(&rutinaGM); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, rutinaGM)
}

// @Summary      Get all routine muscle groups
// @Description  Get a complete list of routine muscle groups
// @Tags         routine-muscle-groups
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.RutinaGrupoMuscular
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /routine-muscle-groups [get]
func (h *RoutineMuscleGroupHandler) GetAll(c *gin.Context) {
	rutinasGM, err := h.usecase.GetAll()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, rutinasGM)
}

// @Summary      Get routine muscle group by ID
// @Description  Get a specific routine muscle group by its ID
// @Tags         routine-muscle-groups
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Routine muscle group ID"
// @Success      200  {object}  models.RutinaGrupoMuscular
// @Failure      400  {object}  errors.ErrorResponse "Invalid ID format"
// @Failure      404  {object}  errors.ErrorResponse "Routine muscle group not found"
// @Router       /routine-muscle-groups/{id} [get]
func (h *RoutineMuscleGroupHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Routine muscle group ID must be a valid number", err))
		return
	}

	rutinaGM, err := h.usecase.GetByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, rutinaGM)
}

// @Summary      Update routine muscle group
// @Description  Update an existing routine muscle group
// @Tags         routine-muscle-groups
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id                    path   int  true  "Routine muscle group ID"
// @Param        routine_muscle_group  body   models.RutinaGrupoMuscular true "Updated routine muscle group data"
// @Success      200  {object}  models.RutinaGrupoMuscular
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      404  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /routine-muscle-groups/{id} [put]
func (h *RoutineMuscleGroupHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Routine muscle group ID must be a valid number", err))
		return
	}

	var rutinaGM models.RutinaGrupoMuscular
	if err := c.ShouldBindJSON(&rutinaGM); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	rutinaGM.ID = uint(id)
	if err := h.usecase.Update(&rutinaGM); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, rutinaGM)
}

// @Summary      Delete routine muscle group
// @Description  Delete a routine muscle group from the system
// @Tags         routine-muscle-groups
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "Routine muscle group ID"
// @Success      204 "No Content"
// @Failure      400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure      500 {object} errors.ErrorResponse "Internal server error"
// @Router       /routine-muscle-groups/{id} [delete]
func (h *RoutineMuscleGroupHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Routine muscle group ID must be a valid number", err))
		return
	}

	if err := h.usecase.Delete(uint(id)); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary      Get muscle groups by routine
// @Description  Get all muscle groups associated with a specific routine
// @Tags         routine-muscle-groups
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "Routine ID"
// @Success      200 {object}  models.RutinaConGruposMusculares
// @Failure      400 {object}  errors.ErrorResponse "Invalid routine ID format"
// @Failure      500 {object}  errors.ErrorResponse "Internal server error"
// @Router       /routine-muscle-groups/routine/{id}/muscle-groups [get]
func (h *RoutineMuscleGroupHandler) GetMuscleGroupsByRoutine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Routine ID must be a valid number", err))
		return
	}

	grupos, err := h.usecase.GetMuscleGroupsByRoutine(uint(id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, grupos)
}
