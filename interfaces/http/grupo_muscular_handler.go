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

func (h *GrupoMuscularHandler) ObtenerTodos(c *gin.Context) {
	grupos, err := h.uc.ObtenerTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grupos)
}

func (h *GrupoMuscularHandler) ObtenerPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	g, err := h.uc.ObtenerPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, g)
}

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

func (h *GrupoMuscularHandler) Eliminar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.uc.Eliminar(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
