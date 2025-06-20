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

func (h *RutinaHandler) ObtenerTodas(c *gin.Context) {
	rutinas, err := h.usecase.Obtener()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener rutinas"})
		return
	}

	c.JSON(http.StatusOK, rutinas)
}

func (h *RutinaHandler) ObtenerporId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	rutina, err := h.usecase.ObtenerPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rutina no encontrada"})
		return
	}

	c.JSON(http.StatusOK, rutina)
}

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

func (h *RutinaHandler) Eliminar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.usecase.Eliminar(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar rutina"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rutina eliminada correctamente"})
}
