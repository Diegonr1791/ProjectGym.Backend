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

// @Summary Crear nueva rutina de grupo muscular
// @Description Crea una nueva rutina de grupo muscular en el sistema
// @Tags rutinas-grupo-muscular
// @Accept json
// @Produce json
// @Param rutina body model.RutinaGrupoMuscular true "Datos de la rutina de grupo muscular"
// @Success 201 {object} model.RutinaGrupoMuscular
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /rutina-gm [post]
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

// @Summary Obtener todas las rutinas de grupo muscular
// @Description Obtiene la lista completa de rutinas de grupo muscular
// @Tags rutinas-grupo-muscular
// @Accept json
// @Produce json
// @Success 200 {array} model.RutinaGrupoMuscular
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /rutina-gm [get]
func (h *RutinaGMHandler) ObtenerTodos(c *gin.Context) {
	rutinasGM, err := h.usecase.ObtenerTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener rutinas grupo muscular"})
		return
	}

	c.JSON(http.StatusOK, rutinasGM)
}

// @Summary Obtener rutina de grupo muscular por ID
// @Description Obtiene una rutina de grupo muscular específica por su ID
// @Tags rutinas-grupo-muscular
// @Accept json
// @Produce json
// @Param id path int true "ID de la rutina de grupo muscular"
// @Success 200 {object} model.RutinaGrupoMuscular
// @Failure 404 {object} map[string]interface{} "Rutina grupo muscular no encontrada"
// @Router /rutina-gm/{id} [get]
func (h *RutinaGMHandler) ObtenerPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	rutinaGM, err := h.usecase.ObtenerPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rutina grupo muscular no encontrada"})
		return
	}

	c.JSON(http.StatusOK, rutinaGM)
}

// @Summary Actualizar rutina de grupo muscular
// @Description Actualiza una rutina de grupo muscular existente
// @Tags rutinas-grupo-muscular
// @Accept json
// @Produce json
// @Param id path int true "ID de la rutina de grupo muscular"
// @Param rutina body model.RutinaGrupoMuscular true "Datos actualizados de la rutina de grupo muscular"
// @Success 200 {object} model.RutinaGrupoMuscular
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 404 {object} map[string]interface{} "Rutina grupo muscular no encontrada"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /rutina-gm/{id} [put]
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

// @Summary Eliminar rutina de grupo muscular
// @Description Elimina una rutina de grupo muscular del sistema
// @Tags rutinas-grupo-muscular
// @Accept json
// @Produce json
// @Param id path int true "ID de la rutina de grupo muscular"
// @Success 200 {object} map[string]interface{} "Rutina grupo muscular eliminada correctamente"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /rutina-gm/{id} [delete]
func (h *RutinaGMHandler) Eliminar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.Eliminar(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar rutina grupo muscular"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rutina grupo muscular eliminada correctamente"})
}

// @Summary Obtener grupos musculares por rutina
// @Description Obtiene todos los grupos musculares asociados a una rutina específica
// @Tags rutinas-grupo-muscular
// @Accept json
// @Produce json
// @Param id path int true "ID de la rutina"
// @Success 200 {array} model.GrupoMuscular
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /rutina/grupo-muscular/{id} [get]
func (h *RutinaGMHandler) ObtenerGruposMuscularesPorRutina(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	grupos, err := h.usecase.ObtenerGruposMuscularesPorRutina(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grupos)
}
