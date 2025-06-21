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

// @Summary Obtener todos los tipos de ejercicio
// @Description Obtiene la lista completa de tipos de ejercicio
// @Tags tipos-ejercicio
// @Accept json
// @Produce json
// @Success 200 {array} model.TipoEjercicio
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /tipo-ejercicio [get]
func (h *TypeExerciseHandler) GetAll(c *gin.Context) {
	tiposEjercicio, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tiposEjercicio)
}

// @Summary Obtener tipo de ejercicio por ID
// @Description Obtiene un tipo de ejercicio específico por su ID
// @Tags tipos-ejercicio
// @Accept json
// @Produce json
// @Param id path int true "ID del tipo de ejercicio"
// @Success 200 {object} model.TipoEjercicio
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /tipo-ejercicio/{id} [get]
func (h *TypeExerciseHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tipoEjercicio, err := h.uc.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tipoEjercicio)
}

// @Summary Crear nuevo tipo de ejercicio
// @Description Crea un nuevo tipo de ejercicio en el sistema
// @Tags tipos-ejercicio
// @Accept json
// @Produce json
// @Param tipo body model.TipoEjercicio true "Datos del tipo de ejercicio"
// @Success 201 {object} model.TipoEjercicio
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /tipo-ejercicio [post]
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

// @Summary Actualizar tipo de ejercicio
// @Description Actualiza un tipo de ejercicio existente
// @Tags tipos-ejercicio
// @Accept json
// @Produce json
// @Param id path int true "ID del tipo de ejercicio"
// @Param tipo body model.TipoEjercicio true "Datos actualizados del tipo de ejercicio"
// @Success 200 {object} model.TipoEjercicio
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 404 {object} map[string]interface{} "Tipo de ejercicio no encontrado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /tipo-ejercicio/{id} [put]
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

// @Summary Eliminar tipo de ejercicio
// @Description Elimina un tipo de ejercicio del sistema
// @Tags tipos-ejercicio
// @Accept json
// @Produce json
// @Param id path int true "ID del tipo de ejercicio"
// @Success 204 "Tipo de ejercicio eliminado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /tipo-ejercicio/{id} [delete]
func (h *TypeExerciseHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.uc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
