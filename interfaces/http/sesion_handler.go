package http

import (
	"net/http"
	"strconv"
	"time"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	usecase "github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	uc *usecase.SessionUsecase
}

func NewSessionHandler(r *gin.Engine, uc *usecase.SessionUsecase) {
	h := &SessionHandler{uc}

	r.POST("/sesion", h.CreateSession)
	r.GET("/sesion", h.GetAllSessions)
	r.GET("/sesion/:id", h.GetSessionById)
	r.PUT("/sesion/:id", h.UpdateSession)
	r.DELETE("/sesion/:id", h.DeleteSession)
	r.GET("/sesion/usuario/:id", h.GetSessionsByUserID)
	r.GET("/sesion/fecha", h.GetSessionsByDateRange)
}

func (h *SessionHandler) CreateSession(c *gin.Context) {
	var sesion model.Sesion
	if err := c.ShouldBindJSON(&sesion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.CreateSession(&sesion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sesion)
}

func (h *SessionHandler) GetAllSessions(c *gin.Context) {
	sesiones, err := h.uc.GetAllSessions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sesiones)
}

func (h *SessionHandler) GetSessionById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sesion, err := h.uc.GetSessionById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sesion)
}

func (h *SessionHandler) UpdateSession(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var sesion model.Sesion
	if err := c.ShouldBindJSON(&sesion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sesion.ID = uint(id)
	if err := h.uc.UpdateSession(&sesion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sesion)
}

func (h *SessionHandler) DeleteSession(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.uc.DeleteSession(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *SessionHandler) GetSessionsByUserID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sesiones, err := h.uc.GetSessionsByUserID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sesiones)
}

func (h *SessionHandler) GetSessionsByDateRange(c *gin.Context) {
	fechaDesdeStr := c.Query("fechaDesde")
	fechaHastaStr := c.Query("fechaHasta")

	// Parse date strings to time.Time (ISO format)
	fechaDesde, err := time.Parse(time.RFC3339, fechaDesdeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fecha inválido. Use formato ISO (2024-01-02T15:04:05Z)"})
		return
	}

	fechaHasta, err := time.Parse(time.RFC3339, fechaHastaStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fecha inválido. Use formato ISO (2024-01-02T15:04:05Z)"})
		return
	}

	sesiones, err := h.uc.GetSessionsByDateRange(fechaDesde, fechaHasta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sesiones)
}
