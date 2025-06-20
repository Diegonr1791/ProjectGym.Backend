package http

import (
	"net/http"
	"strconv"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	usecase "github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ExerciseHandler struct {
	uc *usecase.ExerciseUsecase
}

func NewExerciseHandler(r *gin.Engine, uc *usecase.ExerciseUsecase) {
	h := &ExerciseHandler{uc}

	r.GET("/ejercicio", h.GetAll)
	r.GET("/ejercicio/:id", h.GetById)
	r.POST("/ejercicio", h.Create)
	r.PUT("/ejercicio/:id", h.Update)
	r.DELETE("/ejercicio/:id", h.Delete)
	r.GET("/ejercicio/grupo-muscular/:id", h.GetByMuscleGroup)
}

func (h *ExerciseHandler) GetAll(c *gin.Context) {
	ejercicios, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ejercicios)
}

func (h *ExerciseHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	ejercicio, err := h.uc.GetById(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ejercicio no encontrado"})
		return
	}
	c.JSON(http.StatusOK, ejercicio)
}

func (h *ExerciseHandler) Create(c *gin.Context) {
	var ejercicio model.Ejercicio
	if err := c.ShouldBindJSON(&ejercicio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.Create(&ejercicio); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ejercicio)
}

func (h *ExerciseHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var ejercicioUpdate model.Ejercicio

	if err := c.ShouldBindJSON(&ejercicioUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ejercicioUpdate.ID = uint(id)
	if err := h.uc.Update(&ejercicioUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ejercicioUpdate)
}

func (h *ExerciseHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := h.uc.Delete(uint(idUint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *ExerciseHandler) GetByMuscleGroup(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	ejercicios, err := h.uc.GetByMuscleGroup(uint(idUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ejercicios)
}
