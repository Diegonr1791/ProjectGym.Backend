package http

import (
	"net/http"
	"strconv"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	usecase "github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type MedicionHandler struct {
	uc *usecase.MedicionUsecase
}

func NewMedicionHandler(r gin.IRouter, uc *usecase.MedicionUsecase) {
	h := &MedicionHandler{uc}

	r.POST("/medicion", h.Crear)
	r.GET("/medicion", h.ObtenerTodos)
	r.GET("/medicion/:id", h.ObtenerPorID)
	r.PUT("/medicion/:id", h.Actualizar)
	r.DELETE("/medicion/:id", h.Eliminar)
	r.GET("/medicion/usuario/:usuario_id", h.ObtenerPorUsuarioID)
}

// @Summary Crear nueva medición
// @Description Crea una nueva medición en el sistema
// @Tags mediciones
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param medicion body model.Medicion true "Datos de la medición"
// @Success 201 {object} model.Medicion
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /medicion [post]
func (h *MedicionHandler) Crear(c *gin.Context) {
	var medicion model.Medicion
	if err := c.ShouldBindJSON(&medicion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.Create(&medicion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, medicion)
}

// @Summary Obtener todas las mediciones
// @Description Obtiene la lista completa de mediciones
// @Tags mediciones
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} model.Medicion
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /medicion [get]
func (h *MedicionHandler) ObtenerTodos(c *gin.Context) {
	mediciones, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mediciones)
}

// @Summary Obtener medición por ID
// @Description Obtiene una medición específica por su ID
// @Tags mediciones
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID de la medición"
// @Success 200 {object} model.Medicion
// @Failure 404 {object} map[string]interface{} "Medición no encontrada"
// @Router /medicion/{id} [get]
func (h *MedicionHandler) ObtenerPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	medicion, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medicion)
}

// @Summary Actualizar medición
// @Description Actualiza una medición existente
// @Tags mediciones
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID de la medición"
// @Param medicion body model.Medicion true "Datos actualizados de la medición"
// @Success 200 {object} model.Medicion
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /medicion/{id} [put]
func (h *MedicionHandler) Actualizar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var medicion model.Medicion
	if err := c.ShouldBindJSON(&medicion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	medicion.ID = uint(id)
	if err := h.uc.Update(&medicion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medicion)
}

// @Summary Eliminar medición
// @Description Elimina una medición del sistema
// @Tags mediciones
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID de la medición"
// @Success 204 "Medición eliminada"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /medicion/{id} [delete]
func (h *MedicionHandler) Eliminar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.uc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary Obtener mediciones por usuario
// @Description Obtiene todas las mediciones de un usuario específico
// @Tags mediciones
// @Accept json
// @Produce json
// @Param usuario_id path int true "ID del usuario"
// @Success 200 {array} model.Medicion
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /medicion/usuario/{usuario_id} [get]
func (h *MedicionHandler) ObtenerPorUsuarioID(c *gin.Context) {
	usuarioID, _ := strconv.Atoi(c.Param("usuario_id"))
	mediciones, err := h.uc.GetMesurementsByUserID(uint(usuarioID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mediciones)
}
