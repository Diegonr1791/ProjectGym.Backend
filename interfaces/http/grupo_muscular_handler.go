package http

import (
	"net/http"
	"strconv"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	usecase "github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type GrupoMuscularHandler struct {
	uc *usecase.GrupoMuscularUseCase
}

func NewGrupoMuscularHandler(r *gin.Engine, uc *usecase.GrupoMuscularUseCase) {
	h := &GrupoMuscularHandler{uc}

	r.POST("/grupo-muscular", h.Crear)
	r.GET("/grupo-muscular", h.ObtenerTodos)
	r.GET("/grupo-muscular/:id", h.ObtenerPorID)
	r.PUT("/grupo-muscular/:id", h.Actualizar)
	r.DELETE("/grupo-muscular/:id", h.Eliminar)
}

// @Summary Crear nuevo grupo muscular
// @Description Crea un nuevo grupo muscular en el sistema
// @Tags grupos-musculares
// @Accept json
// @Produce json
// @Param grupo body model.GrupoMuscular true "Datos del grupo muscular"
// @Success 201 {object} model.GrupoMuscular
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /grupo-muscular [post]
func (h *GrupoMuscularHandler) Crear(c *gin.Context) {
	var g model.GrupoMuscular
	if err := c.ShouldBindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.Crear(&g); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, g)
}

// @Summary Obtener todos los grupos musculares
// @Description Obtiene la lista completa de grupos musculares
// @Tags grupos-musculares
// @Accept json
// @Produce json
// @Success 200 {array} model.GrupoMuscular
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /grupo-muscular [get]
func (h *GrupoMuscularHandler) ObtenerTodos(c *gin.Context) {
	grupos, err := h.uc.ObtenerTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grupos)
}

// @Summary Obtener grupo muscular por ID
// @Description Obtiene un grupo muscular específico por su ID
// @Tags grupos-musculares
// @Accept json
// @Produce json
// @Param id path int true "ID del grupo muscular"
// @Success 200 {object} model.GrupoMuscular
// @Failure 404 {object} map[string]interface{} "Grupo muscular no encontrado"
// @Router /grupo-muscular/{id} [get]
func (h *GrupoMuscularHandler) ObtenerPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	g, err := h.uc.ObtenerPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, g)
}

// @Summary Actualizar grupo muscular
// @Description Actualiza un grupo muscular existente
// @Tags grupos-musculares
// @Accept json
// @Produce json
// @Param id path int true "ID del grupo muscular"
// @Param grupo body model.GrupoMuscular true "Datos actualizados del grupo muscular"
// @Success 200 {object} model.GrupoMuscular
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /grupo-muscular/{id} [put]
func (h *GrupoMuscularHandler) Actualizar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var g model.GrupoMuscular
	if err := c.ShouldBindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.ID = uint(id)
	if err := h.uc.Actualizar(&g); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, g)
}

// @Summary Eliminar grupo muscular
// @Description Elimina un grupo muscular del sistema
// @Tags grupos-musculares
// @Accept json
// @Produce json
// @Param id path int true "ID del grupo muscular"
// @Success 204 "Grupo muscular eliminado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /grupo-muscular/{id} [delete]
func (h *GrupoMuscularHandler) Eliminar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.uc.Eliminar(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
