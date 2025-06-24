package http

import (
	"net/http"
	"strconv"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	"github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type RutinaHandler struct {
	usecase *usecase.RutinaUsecase
}

func NewRoutineHandler(r gin.IRouter, usecase *usecase.RutinaUsecase) {
	handler := &RutinaHandler{usecase}

	// Grouping routine routes under "routines"
	routineRoutes := r.Group("/routines")
	{
		routineRoutes.GET("", handler.GetAll)
		routineRoutes.POST("", handler.Create)
		routineRoutes.GET("/:id", handler.GetByID)
		routineRoutes.PUT("/:id", handler.Update)
		routineRoutes.DELETE("/:id", handler.Delete)
	}
}

// @Summary      Get all routines
// @Description  Get a complete list of routines
// @Tags         routines
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.Rutina
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /routines [get]
func (h *RutinaHandler) GetAll(c *gin.Context) {
	rutinas, err := h.usecase.GetAllRoutines()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, rutinas)
}

// @Summary      Get routine by ID
// @Description  Get a specific routine by its ID
// @Tags         routines
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Routine ID"
// @Success      200  {object}  models.Rutina
// @Failure      400  {object}  errors.ErrorResponse "Invalid ID format"
// @Failure      404  {object}  errors.ErrorResponse "Routine not found"
// @Router       /routines/{id} [get]
func (h *RutinaHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Routine ID must be a valid number", err))
		return
	}

	rutina, err := h.usecase.GetRoutineByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, rutina)
}

// @Summary      Create a new routine
// @Description  Create a new routine in the system
// @Tags         routines
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        routine body models.Rutina true "Routine data"
// @Success      201  {object}  models.Rutina
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /routines [post]
func (h *RutinaHandler) Create(c *gin.Context) {
	var rutina models.Rutina
	if err := c.ShouldBindJSON(&rutina); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	if err := h.usecase.CreateRoutine(&rutina); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, rutina)
}

// @Summary      Update routine
// @Description  Update an existing routine
// @Tags         routines
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path   int  true  "Routine ID"
// @Param        routine  body   models.Rutina true "Updated routine data"
// @Success      200  {object}  models.Rutina
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      404  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /routines/{id} [put]
func (h *RutinaHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Routine ID must be a valid number", err))
		return
	}

	var r models.Rutina
	if err := c.ShouldBindJSON(&r); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	r.ID = uint(id)
	if err := h.usecase.UpdateRoutine(&r); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary      Delete routine
// @Description  Delete a routine from the system
// @Tags         routines
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "Routine ID"
// @Success      204 "No Content"
// @Failure      400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure      500 {object} errors.ErrorResponse "Internal server error"
// @Router       /routines/{id} [delete]
func (h *RutinaHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Routine ID must be a valid number", err))
		return
	}

	if err := h.usecase.DeleteRoutine(uint(id)); err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
