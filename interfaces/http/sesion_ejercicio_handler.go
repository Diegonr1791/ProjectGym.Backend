package http

import (
	"net/http"
	"strconv"
	"time"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	usecase "github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type SessionExerciseHandler struct {
	uc *usecase.SessionExerciseUsecase
}

func NewSessionExerciseHandler(r *gin.Engine, uc *usecase.SessionExerciseUsecase) {
	h := &SessionExerciseHandler{uc}

	r.POST("/sesion-ejercicio", h.CreateSessionExercise)
	r.GET("/sesion-ejercicio", h.GetAllSessionExercises)
	r.GET("/sesion-ejercicio/:id", h.GetSessionExerciseById)
	r.PUT("/sesion-ejercicio/:id", h.UpdateSessionExercise)
	r.DELETE("/sesion-ejercicio/:id", h.DeleteSessionExercise)
	r.GET("/sesion-ejercicio/sesion/:id", h.GetSessionExercisesBySessionID)
}

// @Summary Crear nuevo ejercicio de sesión
// @Description Crea un nuevo ejercicio de sesión en el sistema
// @Tags sesiones-ejercicios
// @Accept json
// @Produce json
// @Param sesionEjercicio body model.SesionEjercicio true "Datos del ejercicio de sesión"
// @Success 201 {object} model.SesionEjercicio
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /sesion-ejercicio [post]
func (h *SessionExerciseHandler) CreateSessionExercise(c *gin.Context) {
	var sesionEjercicio model.SesionEjercicio
	if err := c.ShouldBindJSON(&sesionEjercicio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.CreateSessionExercise(&sesionEjercicio); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sesionEjercicio)
}

// @Summary Obtener todos los ejercicios de sesión
// @Description Obtiene la lista completa de ejercicios de sesión
// @Tags sesiones-ejercicios
// @Accept json
// @Produce json
// @Success 200 {array} model.SesionEjercicio
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /sesion-ejercicio [get]
func (h *SessionExerciseHandler) GetAllSessionExercises(c *gin.Context) {
	sesionesEjercicios, err := h.uc.GetAllSessionExercises()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sesionesEjercicios)
}

// @Summary Obtener ejercicio de sesión por ID
// @Description Obtiene un ejercicio de sesión específico por su ID
// @Tags sesiones-ejercicios
// @Accept json
// @Produce json
// @Param id path int true "ID del ejercicio de sesión"
// @Success 200 {object} model.SesionEjercicio
// @Failure 404 {object} map[string]interface{} "Ejercicio de sesión no encontrado"
// @Router /sesion-ejercicio/{id} [get]
func (h *SessionExerciseHandler) GetSessionExerciseById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sesionEjercicio, err := h.uc.GetSessionExerciseById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sesionEjercicio)
}

// @Summary Actualizar ejercicio de sesión
// @Description Actualiza un ejercicio de sesión existente
// @Tags sesiones-ejercicios
// @Accept json
// @Produce json
// @Param id path int true "ID del ejercicio de sesión"
// @Param sesionEjercicio body model.SesionEjercicio true "Datos actualizados del ejercicio de sesión"
// @Success 200 {object} model.SesionEjercicio
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /sesion-ejercicio/{id} [put]
func (h *SessionExerciseHandler) UpdateSessionExercise(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var sesionEjercicio model.SesionEjercicio
	if err := c.ShouldBindJSON(&sesionEjercicio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sesionEjercicio.ID = uint(id)
	if err := h.uc.UpdateSessionExercise(&sesionEjercicio); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sesionEjercicio)
}

// @Summary Eliminar ejercicio de sesión
// @Description Elimina un ejercicio de sesión del sistema
// @Tags sesiones-ejercicios
// @Accept json
// @Produce json
// @Param id path int true "ID del ejercicio de sesión"
// @Success 204 "Ejercicio de sesión eliminado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /sesion-ejercicio/{id} [delete]
func (h *SessionExerciseHandler) DeleteSessionExercise(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.uc.DeleteSessionExercise(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary Obtener ejercicios de sesión por ID de sesión
// @Description Obtiene todos los ejercicios de una sesión específica
// @Tags sesiones-ejercicios
// @Accept json
// @Produce json
// @Param id path int true "ID de la sesión"
// @Param fechaDesde query string false "Fecha desde (formato: 2006-01-02)"
// @Param fechaHasta query string false "Fecha hasta (formato: 2006-01-02)"
// @Success 200 {array} model.SesionEjercicio
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /sesion-ejercicio/sesion/{id} [get]
func (h *SessionExerciseHandler) GetSessionExercisesBySessionID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fechaDesde, _ := time.Parse("2006-01-02", c.Query("fechaDesde"))
	fechaHasta, _ := time.Parse("2006-01-02", c.Query("fechaHasta"))
	sesionesEjercicios, err := h.uc.GetSessionExercisesBySessionID(uint(id), fechaDesde, fechaHasta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sesionesEjercicios)
}
