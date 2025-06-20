package http

import (
	"net/http"
	"strconv"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	rutinaGM "github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type RutinaGMHandler struct {
	usecase *rutinaGM.RutinaGrupoMuscularUsecase
}

func NewRutinaGMHandler(router *gin.Engine, uc *rutinaGM.RutinaGrupoMuscularUsecase) {
	handler := &RutinaGMHandler{uc}

	router.POST("/rutina-gm", handler.Crear)
	router.GET("/rutina-gm", handler.ObtenerTodos)
	router.GET("/rutina-gm/:id", handler.ObtenerPorID)
	router.PUT("/rutina-gm/:id", handler.Actualizar)
	router.DELETE("/rutina-gm/:id", handler.Eliminar)
	router.GET("/rutina/grupo-muscular/:id", handler.ObtenerGruposMuscularesPorRutina)
}

func (h *RutinaGMHandler) Crear(c *gin.Context) {
	var rutinaGM model.RutinaGrupoMuscular
	if err := c.ShouldBindJSON(&rutinaGM); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.Crear(&rutinaGM); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear rutina grupo muscular"})
		return
	}

	c.JSON(http.StatusCreated, rutinaGM)
}

func (h *RutinaGMHandler) ObtenerTodos(c *gin.Context) {
	rutinasGM, err := h.usecase.ObtenerTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener rutinas grupo muscular"})
		return
	}

	c.JSON(http.StatusOK, rutinasGM)
}

func (h *RutinaGMHandler) ObtenerPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	rutinaGM, err := h.usecase.ObtenerPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rutina grupo muscular no encontrada"})
		return
	}

	c.JSON(http.StatusOK, rutinaGM)
}

func (h *RutinaGMHandler) Actualizar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.usecase.ObtenerPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rutina grupo muscular no encontrada"})
		return
	}

	var rutinaGM model.RutinaGrupoMuscular
	if err := c.ShouldBindJSON(&rutinaGM); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.Actualizar(&rutinaGM); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar rutina grupo muscular"})
		return
	}

	c.JSON(http.StatusOK, rutinaGM)
}

func (h *RutinaGMHandler) Eliminar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.Eliminar(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar rutina grupo muscular"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rutina grupo muscular eliminada correctamente"})
}

func (h *RutinaGMHandler) ObtenerGruposMuscularesPorRutina(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	grupos, err := h.usecase.ObtenerGruposMuscularesPorRutina(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grupos)
}
