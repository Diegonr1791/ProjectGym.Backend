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

func NewFavoritaHandler(r gin.IRouter, uc *usecase.FavoritaUsecase) {
	h := &FavoritaHandler{uc}

	r.POST("/favorita", h.Crear)
	r.GET("/favorita", h.ObtenerTodos)
	r.GET("/favorita/:id", h.ObtenerPorID)
	r.GET("/favorita/usuario/:usuario_id", h.ObtenerFavoritasPorUsuario)
	r.PUT("/favorita/:id", h.Actualizar)
	r.DELETE("/favorita/:id", h.Eliminar)
}

// @Summary Crear nueva favorita
// @Description Crea una nueva favorita en el sistema
// @Tags favoritas
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param favorita body model.Favorita true "Datos de la favorita"
// @Success 201 {object} model.Favorita
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /favorita [post]
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

// @Summary Obtener todas las favoritas
// @Description Obtiene la lista completa de favoritas
// @Tags favoritas
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} model.Favorita
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /favorita [get]
func (h *FavoritaHandler) ObtenerTodos(c *gin.Context) {
	favoritas, err := h.uc.ObtenerTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favoritas)
}

// @Summary Obtener favorita por ID
// @Description Obtiene una favorita específica por su ID
// @Tags favoritas
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID de la favorita"
// @Success 200 {object} model.Favorita
// @Failure 404 {object} map[string]interface{} "Favorita no encontrada"
// @Router /favorita/{id} [get]
func (h *FavoritaHandler) ObtenerPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	favorita, err := h.uc.ObtenerPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favorita)
}

// @Summary Actualizar favorita
// @Description Actualiza una favorita existente
// @Tags favoritas
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID de la favorita"
// @Param favorita body model.Favorita true "Datos actualizados de la favorita"
// @Success 200 {object} model.Favorita
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /favorita/{id} [put]
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

// @Summary Eliminar favorita
// @Description Elimina una favorita del sistema
// @Tags favoritas
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID de la favorita"
// @Success 204 "Favorita eliminada"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /favorita/{id} [delete]
func (h *FavoritaHandler) Eliminar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.uc.Eliminar(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary Obtener favoritas por usuario
// @Description Obtiene todas las favoritas de un usuario específico
// @Tags favoritas
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param usuario_id path int true "ID del usuario"
// @Success 200 {array} model.Favorita
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /favorita/usuario/{usuario_id} [get]
func (h *FavoritaHandler) ObtenerFavoritasPorUsuario(c *gin.Context) {
	usuarioID, _ := strconv.Atoi(c.Param("usuario_id"))
	favoritas, err := h.uc.ObtenerFavoritasPorUsuario(uint(usuarioID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favoritas)
}
