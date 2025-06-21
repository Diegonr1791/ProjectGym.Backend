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

func NewExerciseHandler(r gin.IRouter, uc *usecase.ExerciseUsecase) {
	h := &ExerciseHandler{uc}

	r.GET("/ejercicio", h.GetAll)
	r.GET("/ejercicio/:id", h.GetById)
	r.POST("/ejercicio", h.Create)
	r.PUT("/ejercicio/:id", h.Update)
	r.DELETE("/ejercicio/:id", h.Delete)
	r.GET("/ejercicio/grupo-muscular/:id", h.GetByMuscleGroup)
}

// @Summary Obtener todos los ejercicios
// @Description Obtiene la lista completa de ejercicios
// @Tags ejercicios
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} model.Ejercicio
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /ejercicio [get]
func (h *ExerciseHandler) GetAll(c *gin.Context) {
	ejercicios, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ejercicios)
}

// @Summary Obtener ejercicio por ID
// @Description Obtiene un ejercicio específico por su ID
// @Tags ejercicios
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID del ejercicio"
// @Success 200 {object} model.Ejercicio
// @Failure 400 {object} map[string]interface{} "ID inválido"
// @Failure 404 {object} map[string]interface{} "Ejercicio no encontrado"
// @Router /ejercicio/{id} [get]
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

// @Summary Crear nuevo ejercicio
// @Description Crea un nuevo ejercicio en el sistema
// @Tags ejercicios
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param ejercicio body model.Ejercicio true "Datos del ejercicio"
// @Success 201 {object} model.Ejercicio
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /ejercicio [post]
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

// @Summary Actualizar ejercicio
// @Description Actualiza un ejercicio existente
// @Tags ejercicios
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID del ejercicio"
// @Param ejercicio body model.Ejercicio true "Datos actualizados del ejercicio"
// @Success 200 {object} model.Ejercicio
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /ejercicio/{id} [put]
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

// @Summary Eliminar ejercicio
// @Description Elimina un ejercicio del sistema
// @Tags ejercicios
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID del ejercicio"
// @Success 204 "Ejercicio eliminado"
// @Failure 400 {object} map[string]interface{} "ID inválido"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /ejercicio/{id} [delete]
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

// @Summary Obtener ejercicios por grupo muscular
// @Description Obtiene todos los ejercicios de un grupo muscular específico
// @Tags ejercicios
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID del grupo muscular"
// @Success 200 {array} model.Ejercicio
// @Failure 400 {object} map[string]interface{} "ID inválido"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /ejercicio/grupo-muscular/{id} [get]
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
