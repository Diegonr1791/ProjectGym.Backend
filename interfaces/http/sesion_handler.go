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

type SessionHandler struct {
	uc *usecase.SessionUsecase
}

func NewSessionHandler(r gin.IRouter, uc *usecase.SessionUsecase) {
	h := &SessionHandler{uc}

	// Grouping session routes under "sessions"
	sessionRoutes := r.Group("/sessions")
	{
		sessionRoutes.GET("", h.GetAll)
		sessionRoutes.POST("", h.Create)
		sessionRoutes.GET("/:id", h.GetByID)
		sessionRoutes.PUT("/:id", h.Update)
		sessionRoutes.DELETE("/:id", h.Delete)
		sessionRoutes.GET("/user/:id", h.GetByUserID)
		sessionRoutes.GET("/date-range", h.GetByDateRange)
	}
}

// @Summary      Create a new session
// @Description  Create a new training session in the system
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        session body models.Sesion true "Session data"
// @Success      201  {object}  models.Sesion
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /sessions [post]
func (h *SessionHandler) Create(c *gin.Context) {
	var sesion models.Sesion
	if err := c.ShouldBindJSON(&sesion); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	if err := h.uc.CreateSession(&sesion); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, sesion)
}

// @Summary      Get all sessions
// @Description  Get a complete list of training sessions
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.Sesion
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /sessions [get]
func (h *SessionHandler) GetAll(c *gin.Context) {
	sesiones, err := h.uc.GetAllSessions()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, sesiones)
}

// @Summary      Get session by ID
// @Description  Get a specific session by its ID
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Session ID"
// @Success      200  {object}  models.Sesion
// @Failure      400  {object}  errors.ErrorResponse "Invalid ID format"
// @Failure      404  {object}  errors.ErrorResponse "Session not found"
// @Router       /sessions/{id} [get]
func (h *SessionHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Session ID must be a valid number", err))
		return
	}

	sesion, err := h.uc.GetSessionByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, sesion)
}

// @Summary      Update session
// @Description  Update an existing session
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path   int  true  "Session ID"
// @Param        session  body   models.Sesion true "Updated session data"
// @Success      200  {object}  models.Sesion
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      404  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /sessions/{id} [put]
func (h *SessionHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Session ID must be a valid number", err))
		return
	}

	var sesion models.Sesion
	if err := c.ShouldBindJSON(&sesion); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	sesion.ID = uint(id)
	if err := h.uc.UpdateSession(&sesion); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, sesion)
}

// @Summary      Delete session
// @Description  Delete a session from the system
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "Session ID"
// @Success      204 "No Content"
// @Failure      400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure      500 {object} errors.ErrorResponse "Internal server error"
// @Router       /sessions/{id} [delete]
func (h *SessionHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Session ID must be a valid number", err))
		return
	}

	if err := h.uc.DeleteSession(uint(id)); err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary      Get sessions by user
// @Description  Get all sessions of a specific user
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "User ID"
// @Success      200 {array}   models.Sesion
// @Failure      400 {object}  errors.ErrorResponse "Invalid ID format"
// @Failure      500 {object}  errors.ErrorResponse "Internal server error"
// @Router       /sessions/user/{id} [get]
func (h *SessionHandler) GetByUserID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "User ID must be a valid number", err))
		return
	}

	sesiones, err := h.uc.GetSessionsByUserID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, sesiones)
}

// @Summary      Get sessions by date range
// @Description  Get all sessions within a specific date range
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        start_date query string true "Start date (ISO format: 2024-01-02T15:04:05Z)"
// @Param        end_date   query string true "End date (ISO format: 2024-01-02T15:04:05Z)"
// @Success      200 {array}   models.Sesion
// @Failure      400 {object}  errors.ErrorResponse "Invalid date format"
// @Failure      500 {object}  errors.ErrorResponse "Internal server error"
// @Router       /sessions/date-range [get]
func (h *SessionHandler) GetByDateRange(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	// Parse date strings to time.Time (ISO format)
	startDate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_DATE_FORMAT", "Invalid start date format. Use ISO format (2024-01-02T15:04:05Z)", err))
		return
	}

	endDate, err := time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_DATE_FORMAT", "Invalid end date format. Use ISO format (2024-01-02T15:04:05Z)", err))
		return
	}

	sesiones, err := h.uc.GetSessionsByDateRange(startDate, endDate)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, sesiones)
}
