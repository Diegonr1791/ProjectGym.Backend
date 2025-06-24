package http

import (
	"net/http"
	"strconv"
	"time"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	"github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type SessionExerciseHandler struct {
	uc *usecase.SessionExerciseUsecase
}

func NewSessionExerciseHandler(r gin.IRouter, uc *usecase.SessionExerciseUsecase) {
	h := &SessionExerciseHandler{uc}

	// Grouping session exercise routes under "session-exercises"
	sessionExerciseRoutes := r.Group("/session-exercises")
	{
		sessionExerciseRoutes.GET("", h.GetAll)
		sessionExerciseRoutes.POST("", h.Create)
		sessionExerciseRoutes.GET("/:id", h.GetByID)
		sessionExerciseRoutes.PUT("/:id", h.Update)
		sessionExerciseRoutes.DELETE("/:id", h.Delete)
		sessionExerciseRoutes.GET("/session/:id", h.GetBySessionID)
	}
}

// @Summary      Create a new session exercise
// @Description  Create a new session exercise in the system
// @Tags         session-exercises
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        session_exercise body models.SesionEjercicio true "Session exercise data"
// @Success      201  {object}  models.SesionEjercicio
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /session-exercises [post]
func (h *SessionExerciseHandler) Create(c *gin.Context) {
	var sesionEjercicio models.SesionEjercicio
	if err := c.ShouldBindJSON(&sesionEjercicio); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	if err := h.uc.CreateSessionExercise(&sesionEjercicio); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, sesionEjercicio)
}

// @Summary      Get all session exercises
// @Description  Get a complete list of session exercises
// @Tags         session-exercises
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.SesionEjercicio
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /session-exercises [get]
func (h *SessionExerciseHandler) GetAll(c *gin.Context) {
	sesionesEjercicios, err := h.uc.GetAllSessionExercises()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, sesionesEjercicios)
}

// @Summary      Get session exercise by ID
// @Description  Get a specific session exercise by its ID
// @Tags         session-exercises
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Session exercise ID"
// @Success      200  {object}  models.SesionEjercicio
// @Failure      400  {object}  errors.ErrorResponse "Invalid ID format"
// @Failure      404  {object}  errors.ErrorResponse "Session exercise not found"
// @Router       /session-exercises/{id} [get]
func (h *SessionExerciseHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Session exercise ID must be a valid number", err))
		return
	}

	sesionEjercicio, err := h.uc.GetSessionExerciseByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, sesionEjercicio)
}

// @Summary      Update session exercise
// @Description  Update an existing session exercise
// @Tags         session-exercises
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id                path   int  true  "Session exercise ID"
// @Param        session_exercise  body   models.SesionEjercicio true "Updated session exercise data"
// @Success      200  {object}  models.SesionEjercicio
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      404  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /session-exercises/{id} [put]
func (h *SessionExerciseHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Session exercise ID must be a valid number", err))
		return
	}

	var sesionEjercicio models.SesionEjercicio
	if err := c.ShouldBindJSON(&sesionEjercicio); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	sesionEjercicio.ID = uint(id)
	if err := h.uc.UpdateSessionExercise(&sesionEjercicio); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, sesionEjercicio)
}

// @Summary      Delete session exercise
// @Description  Delete a session exercise from the system
// @Tags         session-exercises
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "Session exercise ID"
// @Success      204 "No Content"
// @Failure      400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure      500 {object} errors.ErrorResponse "Internal server error"
// @Router       /session-exercises/{id} [delete]
func (h *SessionExerciseHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Session exercise ID must be a valid number", err))
		return
	}

	if err := h.uc.DeleteSessionExercise(uint(id)); err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary      Get session exercises by session ID
// @Description  Get all exercises of a specific session
// @Tags         session-exercises
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id           path   int     true  "Session ID"
// @Param        start_date   query  string  false "Start date (format: 2006-01-02)"
// @Param        end_date     query  string  false "End date (format: 2006-01-02)"
// @Success      200 {array}   models.SesionEjercicio
// @Failure      400 {object}  errors.ErrorResponse "Invalid ID format or date format"
// @Failure      500 {object}  errors.ErrorResponse "Internal server error"
// @Router       /session-exercises/session/{id} [get]
func (h *SessionExerciseHandler) GetBySessionID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Session ID must be a valid number", err))
		return
	}

	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate time.Time
	var err1, err2 error

	if startDateStr != "" {
		startDate, err1 = time.Parse("2006-01-02", startDateStr)
		if err1 != nil {
			c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_DATE_FORMAT", "Invalid start date format. Use format: 2006-01-02", err1))
			return
		}
	}

	if endDateStr != "" {
		endDate, err2 = time.Parse("2006-01-02", endDateStr)
		if err2 != nil {
			c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_DATE_FORMAT", "Invalid end date format. Use format: 2006-01-02", err2))
			return
		}
	}

	sesionesEjercicios, err := h.uc.GetSessionExercisesBySessionID(uint(id), startDate, endDate)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, sesionesEjercicios)
}
