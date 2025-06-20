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

func (h *SessionExerciseHandler) GetAllSessionExercises(c *gin.Context) {
	sesionesEjercicios, err := h.uc.GetAllSessionExercises()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sesionesEjercicios)
}

func (h *SessionExerciseHandler) GetSessionExerciseById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sesionEjercicio, err := h.uc.GetSessionExerciseById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sesionEjercicio)
}

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

func (h *SessionExerciseHandler) DeleteSessionExercise(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.uc.DeleteSessionExercise(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

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
