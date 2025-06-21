package http

import (
	"net/http"
	"strconv"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	rutina "github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type RutinaHandler struct {
	usecase *rutina.RutinaService
}

func NewRutinaHandler(r *gin.Engine, usecase *rutina.RutinaService) {
	handler := &RutinaHandler{usecase}

	r.GET("/rutinas", handler.ObtenerTodas)
	r.GET("/rutinas/:id", handler.ObtenerporId)
	r.POST("/rutinas", handler.Crear)
	r.PUT("/rutinas/:id", handler.Actualizar)
	r.DELETE("/rutinas/:id", handler.Eliminar)
}

// @Summary Obtener todas las rutinas
// @Description Obtiene la lista completa de rutinas
// @Tags rutinas
// @Accept json
// @Produce json
// @Success 200 {array} model.Rutina
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /rutinas [get]
func (h *RutinaHandler) ObtenerTodas(c *gin.Context) {
	rutinas, err := h.usecase.Obtener()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener rutinas"})
		return
	}

	c.JSON(http.StatusOK, rutinas)
}

// @Summary Obtener rutina por ID
// @Description Obtiene una rutina específica por su ID
// @Tags rutinas
// @Accept json
// @Produce json
// @Param id path int true "ID de la rutina"
// @Success 200 {object} model.Rutina
// @Failure 404 {object} map[string]interface{} "Rutina no encontrada"
// @Router /rutinas/{id} [get]
func (h *RutinaHandler) ObtenerporId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	rutina, err := h.usecase.ObtenerPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rutina no encontrada"})
		return
	}

	c.JSON(http.StatusOK, rutina)
}

// @Summary Crear nueva rutina
// @Description Crea una nueva rutina en el sistema
// @Tags rutinas
// @Accept json
// @Produce json
// @Param rutina body model.Rutina true "Datos de la rutina"
// @Success 200 {object} model.Rutina
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /rutinas [post]
func (h *RutinaHandler) Crear(c *gin.Context) {
	var rutina model.Rutina
	if err := c.ShouldBindJSON(&rutina); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.Crear(&rutina); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear rutina"})
		return
	}

	c.JSON(http.StatusOK, rutina)
}

// @Summary Actualizar rutina
// @Description Actualiza una rutina existente
// @Tags rutinas
// @Accept json
// @Produce json
// @Param id path int true "ID de la rutina"
// @Param rutina body model.Rutina true "Datos actualizados de la rutina"
// @Success 200 {object} model.Rutina
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 404 {object} map[string]interface{} "Rutina no encontrada"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /rutinas/{id} [put]
func (h *RutinaHandler) Actualizar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.usecase.ObtenerPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rutina no encontrada"})
		return
	}

	var r model.Rutina
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	if err := h.usecase.Actualizar(&r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar rutina"})
		return
	}

	c.JSON(http.StatusOK, r)
}

// @Summary Eliminar rutina
// @Description Elimina una rutina del sistema
// @Tags rutinas
// @Accept json
// @Produce json
// @Param id path int true "ID de la rutina"
// @Success 200 {object} map[string]interface{} "Rutina eliminada correctamente"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /rutinas/{id} [delete]
func (h *RutinaHandler) Eliminar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.usecase.Eliminar(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar rutina"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rutina eliminada correctamente"})
}
