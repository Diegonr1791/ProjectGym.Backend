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

func NewMedicionHandler(r *gin.Engine, uc *usecase.MedicionUsecase) {
	h := &MedicionHandler{uc}

	r.POST("/medicion", h.Crear)
	r.GET("/medicion", h.ObtenerTodos)
	r.GET("/medicion/:id", h.ObtenerPorID)
	r.PUT("/medicion/:id", h.Actualizar)
	r.DELETE("/medicion/:id", h.Eliminar)
	r.GET("/medicion/usuario/:usuario_id", h.ObtenerPorUsuarioID)
}

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

func (h *MedicionHandler) ObtenerTodos(c *gin.Context) {
	mediciones, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mediciones)
}

func (h *MedicionHandler) ObtenerPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	medicion, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medicion)
}

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

func (h *MedicionHandler) Eliminar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.uc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *MedicionHandler) ObtenerPorUsuarioID(c *gin.Context) {
	usuarioID, _ := strconv.Atoi(c.Param("usuario_id"))
	mediciones, err := h.uc.GetMesurementsByUserID(uint(usuarioID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mediciones)
}
