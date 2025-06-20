package http

import (
	"net/http"
	"strconv"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	usecase "github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type FavoritaHandler struct {
	uc *usecase.FavoritaUsecase
}

func NewFavoritaHandler(r *gin.Engine, uc *usecase.FavoritaUsecase) {
	h := &FavoritaHandler{uc}

	r.POST("/favorita", h.Crear)
	r.GET("/favorita", h.ObtenerTodos)
	r.GET("/favorita/:id", h.ObtenerPorID)
	r.GET("/favorita/usuario/:usuario_id", h.ObtenerFavoritasPorUsuario)
	r.PUT("/favorita/:id", h.Actualizar)
	r.DELETE("/favorita/:id", h.Eliminar)
}

func (h *FavoritaHandler) Crear(c *gin.Context) {
	var favorita model.Favorita
	if err := c.ShouldBindJSON(&favorita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.Crear(&favorita); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, favorita)
}

func (h *FavoritaHandler) ObtenerTodos(c *gin.Context) {
	favoritas, err := h.uc.ObtenerTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favoritas)
}

func (h *FavoritaHandler) ObtenerPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	favorita, err := h.uc.ObtenerPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favorita)
}

func (h *FavoritaHandler) Actualizar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var favorita model.Favorita
	if err := c.ShouldBindJSON(&favorita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	favorita.ID = uint(id)
	if err := h.uc.Actualizar(&favorita); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favorita)
}

func (h *FavoritaHandler) Eliminar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.uc.Eliminar(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *FavoritaHandler) ObtenerFavoritasPorUsuario(c *gin.Context) {
	usuarioID, _ := strconv.Atoi(c.Param("usuario_id"))
	favoritas, err := h.uc.ObtenerFavoritasPorUsuario(uint(usuarioID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favoritas)
}
