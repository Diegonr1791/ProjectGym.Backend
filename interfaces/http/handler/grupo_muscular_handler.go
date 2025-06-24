package http

import (
	"net/http"
	"strconv"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	"github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type GrupoMuscularHandler struct {
	uc *usecase.GrupoMuscularUseCase
}

func NewGrupoMuscularHandler(r gin.IRouter, uc *usecase.GrupoMuscularUseCase) {
	h := &GrupoMuscularHandler{uc}

	// Grouping muscle group routes under "muscle-groups"
	muscleGroupRoutes := r.Group("/muscle-groups")
	{
		muscleGroupRoutes.GET("", h.GetAll)
		muscleGroupRoutes.POST("", h.Create)
		muscleGroupRoutes.GET("/:id", h.GetByID)
		muscleGroupRoutes.PUT("/:id", h.Update)
		muscleGroupRoutes.DELETE("/:id", h.Delete)
	}
}

// @Summary      Create a new muscle group
// @Description  Create a new muscle group in the system
// @Tags         muscle-groups
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        muscle-group body models.GrupoMuscular true "Muscle group data"
// @Success      201  {object}  models.GrupoMuscular
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /muscle-groups [post]
func (h *GrupoMuscularHandler) Create(c *gin.Context) {
	var g models.GrupoMuscular
	if err := c.ShouldBindJSON(&g); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	if err := h.uc.CreateMuscleGroup(&g); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, g)
}

// @Summary      Get all muscle groups
// @Description  Get a complete list of muscle groups
// @Tags         muscle-groups
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.GrupoMuscular
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /muscle-groups [get]
func (h *GrupoMuscularHandler) GetAll(c *gin.Context) {
	grupos, err := h.uc.GetAllMuscleGroups()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, grupos)
}

// @Summary      Get muscle group by ID
// @Description  Get a specific muscle group by its ID
// @Tags         muscle-groups
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Muscle Group ID"
// @Success      200  {object}  models.GrupoMuscular
// @Failure      400  {object}  errors.ErrorResponse "Invalid ID format"
// @Failure      404  {object}  errors.ErrorResponse "Muscle group not found"
// @Router       /muscle-groups/{id} [get]
func (h *GrupoMuscularHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Muscle group ID must be a valid number", err))
		return
	}

	grupo, err := h.uc.GetMuscleGroupByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, grupo)
}

// @Summary      Update muscle group
// @Description  Update an existing muscle group
// @Tags         muscle-groups
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id           path   int  true  "Muscle Group ID"
// @Param        muscle-group body   models.GrupoMuscular true "Updated muscle group data"
// @Success      200  {object}  models.GrupoMuscular
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      404  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /muscle-groups/{id} [put]
func (h *GrupoMuscularHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Muscle group ID must be a valid number", err))
		return
	}

	var g models.GrupoMuscular
	if err := c.ShouldBindJSON(&g); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	g.ID = uint(id)
	if err := h.uc.UpdateMuscleGroup(&g); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, g)
}

// @Summary      Delete muscle group
// @Description  Delete a muscle group from the system
// @Tags         muscle-groups
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "Muscle Group ID"
// @Success      204 "No Content"
// @Failure      400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure      500 {object} errors.ErrorResponse "Internal server error"
// @Router       /muscle-groups/{id} [delete]
func (h *GrupoMuscularHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Muscle group ID must be a valid number", err))
		return
	}

	if err := h.uc.DeleteMuscleGroup(uint(id)); err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
