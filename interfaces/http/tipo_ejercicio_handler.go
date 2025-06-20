package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	usecase "github.com/Diegonr1791/GymBro/internal/usecase"
)

type TypeExerciseHandler struct {
	uc *usecase.TypeExerciseUsecase
}

func NewTypeExerciseHandler(router *gin.Engine, uc *usecase.TypeExerciseUsecase) {
	handler := &TypeExerciseHandler{uc}

	router.GET("/tipo-ejercicio", handler.GetAll)
	router.GET("/tipo-ejercicio/:id", handler.GetById)
	router.POST("/tipo-ejercicio", handler.Create)
	router.PUT("/tipo-ejercicio/:id", handler.Update)
	router.DELETE("/tipo-ejercicio/:id", handler.Delete)
}

func (h *TypeExerciseHandler) GetAll(c *gin.Context) {
	tiposEjercicio, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tiposEjercicio)
}

func (h *TypeExerciseHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tipoEjercicio, err := h.uc.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tipoEjercicio)
}

func (h *TypeExerciseHandler) Create(c *gin.Context) {
	var tipoEjercicio model.TipoEjercicio
	if err := c.ShouldBindJSON(&tipoEjercicio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.Create(&tipoEjercicio); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tipoEjercicio)
}

func (h *TypeExerciseHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := h.uc.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tipo de ejercicio no encontrado"})
		return
	}

	var tipoEjercicio model.TipoEjercicio

	if err := h.uc.Update(&tipoEjercicio); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tipoEjercicio)
}

func (h *TypeExerciseHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.uc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
