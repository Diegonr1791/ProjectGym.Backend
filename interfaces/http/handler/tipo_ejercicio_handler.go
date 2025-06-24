package http

import (
	"net/http"
	"strconv"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	"github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TypeExerciseHandler struct {
	uc *usecase.TypeExerciseUsecase
}

func NewTypeExerciseHandler(router gin.IRouter, uc *usecase.TypeExerciseUsecase) {
	handler := &TypeExerciseHandler{uc}

	// Grouping exercise type routes under "exercise-types"
	exerciseTypeRoutes := router.Group("/exercise-types")
	{
		exerciseTypeRoutes.GET("", handler.GetAll)
		exerciseTypeRoutes.POST("", handler.Create)
		exerciseTypeRoutes.GET("/:id", handler.GetByID)
		exerciseTypeRoutes.PUT("/:id", handler.Update)
		exerciseTypeRoutes.DELETE("/:id", handler.Delete)
	}
}

// @Summary      Get all exercise types
// @Description  Get a complete list of exercise types
// @Tags         exercise-types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.TipoEjercicio
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /exercise-types [get]
func (h *TypeExerciseHandler) GetAll(c *gin.Context) {
	tiposEjercicio, err := h.uc.GetAllExerciseTypes()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, tiposEjercicio)
}

// @Summary      Get exercise type by ID
// @Description  Get a specific exercise type by its ID
// @Tags         exercise-types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Exercise Type ID"
// @Success      200  {object}  models.TipoEjercicio
// @Failure      400  {object}  errors.ErrorResponse "Invalid ID format"
// @Failure      404  {object}  errors.ErrorResponse "Exercise type not found"
// @Router       /exercise-types/{id} [get]
func (h *TypeExerciseHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Exercise type ID must be a valid number", err))
		return
	}

	tipoEjercicio, err := h.uc.GetExerciseTypeByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, tipoEjercicio)
}

// @Summary      Create a new exercise type
// @Description  Create a new exercise type in the system
// @Tags         exercise-types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        exercise-type body models.TipoEjercicio true "Exercise type data"
// @Success      201  {object}  models.TipoEjercicio
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /exercise-types [post]
func (h *TypeExerciseHandler) Create(c *gin.Context) {
	var tipoEjercicio models.TipoEjercicio
	if err := c.ShouldBindJSON(&tipoEjercicio); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	if err := h.uc.CreateExerciseType(&tipoEjercicio); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, tipoEjercicio)
}

// @Summary      Update exercise type
// @Description  Update an existing exercise type
// @Tags         exercise-types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id             path   int  true  "Exercise Type ID"
// @Param        exercise-type  body   models.TipoEjercicio true "Updated exercise type data"
// @Success      200  {object}  models.TipoEjercicio
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      404  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /exercise-types/{id} [put]
func (h *TypeExerciseHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Exercise type ID must be a valid number", err))
		return
	}

	var tipoEjercicio models.TipoEjercicio
	if err := c.ShouldBindJSON(&tipoEjercicio); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	tipoEjercicio.ID = uint(id)
	if err := h.uc.UpdateExerciseType(&tipoEjercicio); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, tipoEjercicio)
}

// @Summary      Delete exercise type
// @Description  Delete an exercise type from the system
// @Tags         exercise-types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "Exercise Type ID"
// @Success      204 "No Content"
// @Failure      400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure      500 {object} errors.ErrorResponse "Internal server error"
// @Router       /exercise-types/{id} [delete]
func (h *TypeExerciseHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Exercise type ID must be a valid number", err))
		return
	}

	if err := h.uc.DeleteExerciseType(uint(id)); err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
