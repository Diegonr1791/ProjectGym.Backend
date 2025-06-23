package http

import (
	"net/http"
	"strconv"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	"github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ExerciseHandler struct {
	uc *usecase.ExerciseUsecase
}

func NewExerciseHandler(r gin.IRouter, uc *usecase.ExerciseUsecase) {
	h := &ExerciseHandler{uc}

	// Grouping exercise routes under "exercises"
	exerciseRoutes := r.Group("/exercises")
	{
		exerciseRoutes.GET("", h.GetAll)
		exerciseRoutes.POST("", h.Create)
		exerciseRoutes.GET("/:id", h.GetByID)
		exerciseRoutes.PUT("/:id", h.Update)
		exerciseRoutes.DELETE("/:id", h.Delete)
		exerciseRoutes.GET("/muscle-group/:id", h.GetByMuscleGroup)
	}
}

// @Summary      Get all exercises
// @Description  Get a complete list of exercises
// @Tags         exercises
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.Ejercicio
// @Failure      500  {object}  domainErrors.ErrorResponse
// @Router       /exercises [get]
func (h *ExerciseHandler) GetAll(c *gin.Context) {
	ejercicios, err := h.uc.GetAllExercises()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, ejercicios)
}

// @Summary      Get exercise by ID
// @Description  Get a specific exercise by its ID
// @Tags         exercises
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Exercise ID"
// @Success      200  {object}  models.Ejercicio
// @Failure      400  {object}  domainErrors.ErrorResponse "Invalid ID format"
// @Failure      404  {object}  domainErrors.ErrorResponse "Exercise not found"
// @Router       /exercises/{id} [get]
func (h *ExerciseHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Exercise ID must be a valid number", err))
		return
	}

	ejercicio, err := h.uc.GetExerciseByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, ejercicio)
}

// @Summary      Create a new exercise
// @Description  Create a new exercise in the system
// @Tags         exercises
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        exercise body models.Ejercicio true "Exercise data"
// @Success      201  {object}  models.Ejercicio
// @Failure      400  {object}  domainErrors.ErrorResponse
// @Failure      500  {object}  domainErrors.ErrorResponse
// @Router       /exercises [post]
func (h *ExerciseHandler) Create(c *gin.Context) {
	var ejercicio models.Ejercicio
	if err := c.ShouldBindJSON(&ejercicio); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	if err := h.uc.CreateExercise(&ejercicio); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, ejercicio)
}

// @Summary      Update exercise
// @Description  Update an existing exercise
// @Tags         exercises
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id        path   int  true  "Exercise ID"
// @Param        exercise  body   models.Ejercicio true "Updated exercise data"
// @Success      200  {object}  models.Ejercicio
// @Failure      400  {object}  domainErrors.ErrorResponse
// @Failure      404  {object}  domainErrors.ErrorResponse
// @Failure      500  {object}  domainErrors.ErrorResponse
// @Router       /exercises/{id} [put]
func (h *ExerciseHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Exercise ID must be a valid number", err))
		return
	}

	var ejercicioUpdate models.Ejercicio
	if err := c.ShouldBindJSON(&ejercicioUpdate); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	ejercicioUpdate.ID = uint(id)
	if err := h.uc.UpdateExercise(&ejercicioUpdate); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, ejercicioUpdate)
}

// @Summary      Delete exercise
// @Description  Delete an exercise from the system
// @Tags         exercises
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "Exercise ID"
// @Success      204 "No Content"
// @Failure      400 {object} domainErrors.ErrorResponse "Invalid ID format"
// @Failure      500 {object} domainErrors.ErrorResponse "Internal server error"
// @Router       /exercises/{id} [delete]
func (h *ExerciseHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Exercise ID must be a valid number", err))
		return
	}

	if err := h.uc.DeleteExercise(uint(id)); err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary      Get exercises by muscle group
// @Description  Get all exercises of a specific muscle group
// @Tags         exercises
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "Muscle Group ID"
// @Success      200 {array}   models.Ejercicio
// @Failure      400 {object}  domainErrors.ErrorResponse "Invalid ID format"
// @Failure      500 {object}  domainErrors.ErrorResponse "Internal server error"
// @Router       /exercises/muscle-group/{id} [get]
func (h *ExerciseHandler) GetByMuscleGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Muscle group ID must be a valid number", err))
		return
	}

	ejercicios, err := h.uc.GetExercisesByMuscleGroup(uint(id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, ejercicios)
}
