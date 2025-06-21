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

// @Summary Crear nueva sesión
// @Description Crea una nueva sesión de entrenamiento en el sistema
// @Tags sesiones
// @Accept json
// @Produce json
// @Param sesion body model.Sesion true "Datos de la sesión"
// @Success 201 {object} model.Sesion
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /sesion [post]
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

// @Summary Obtener todas las sesiones
// @Description Obtiene la lista completa de sesiones de entrenamiento
// @Tags sesiones
// @Accept json
// @Produce json
// @Success 200 {array} model.Sesion
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /sesion [get]
func (h *SessionHandler) GetAllSessions(c *gin.Context) {
	sesiones, err := h.uc.GetAllSessions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sesiones)
}

// @Summary Obtener sesión por ID
// @Description Obtiene una sesión específica por su ID
// @Tags sesiones
// @Accept json
// @Produce json
// @Param id path int true "ID de la sesión"
// @Success 200 {object} model.Sesion
// @Failure 404 {object} map[string]interface{} "Sesión no encontrada"
// @Router /sesion/{id} [get]
func (h *SessionHandler) GetSessionById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sesion, err := h.uc.GetSessionById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sesion)
}

// @Summary Actualizar sesión
// @Description Actualiza una sesión existente
// @Tags sesiones
// @Accept json
// @Produce json
// @Param id path int true "ID de la sesión"
// @Param sesion body model.Sesion true "Datos actualizados de la sesión"
// @Success 200 {object} model.Sesion
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /sesion/{id} [put]
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

// @Summary Eliminar sesión
// @Description Elimina una sesión del sistema
// @Tags sesiones
// @Accept json
// @Produce json
// @Param id path int true "ID de la sesión"
// @Success 204 "Sesión eliminada"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /sesion/{id} [delete]
func (h *SessionHandler) DeleteSession(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.uc.DeleteSession(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary Obtener sesiones por usuario
// @Description Obtiene todas las sesiones de un usuario específico
// @Tags sesiones
// @Accept json
// @Produce json
// @Param id path int true "ID del usuario"
// @Success 200 {array} model.Sesion
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /sesion/usuario/{id} [get]
func (h *SessionHandler) GetSessionsByUserID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sesiones, err := h.uc.GetSessionsByUserID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sesiones)
}

// @Summary Obtener sesiones por rango de fechas
// @Description Obtiene todas las sesiones dentro de un rango de fechas específico
// @Tags sesiones
// @Accept json
// @Produce json
// @Param fechaDesde query string true "Fecha desde (formato ISO: 2024-01-02T15:04:05Z)"
// @Param fechaHasta query string true "Fecha hasta (formato ISO: 2024-01-02T15:04:05Z)"
// @Success 200 {array} model.Sesion
// @Failure 400 {object} map[string]interface{} "Formato de fecha inválido"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /sesion/fecha [get]
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
