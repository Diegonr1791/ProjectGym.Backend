package http

import (
	"net/http"
	"strconv"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	"github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type MeasurementHandler struct {
	uc *usecase.MeasurementUsecase
}

func NewMeasurementHandler(r gin.IRouter, uc *usecase.MeasurementUsecase) {
	h := &MeasurementHandler{uc}

	measurementRoutes := r.Group("/measurements")
	{
		measurementRoutes.POST("", h.Create)
		measurementRoutes.GET("", h.GetAll)
		measurementRoutes.GET("/:id", h.GetByID)
		measurementRoutes.PUT("/:id", h.Update)
		measurementRoutes.DELETE("/:id", h.Delete)
		measurementRoutes.GET("/user/:user_id", h.GetByUserID)
	}
}

// @Summary      Create a new measurement
// @Description  Create a new measurement in the system
// @Tags         measurements
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        measurement body models.Medicion true "Measurement data"
// @Success      201  {object}  models.Medicion
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /measurements [post]
func (h *MeasurementHandler) Create(c *gin.Context) {
	var medicion models.Medicion
	if err := c.ShouldBindJSON(&medicion); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}
	if err := h.uc.Create(&medicion); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, medicion)
}

// @Summary      Get all measurements
// @Description  Get a complete list of measurements
// @Tags         measurements
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.Medicion
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /measurements [get]
func (h *MeasurementHandler) GetAll(c *gin.Context) {
	mediciones, err := h.uc.GetAll()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, mediciones)
}

// @Summary      Get measurement by ID
// @Description  Get a specific measurement by its ID
// @Tags         measurements
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Measurement ID"
// @Success      200  {object}  models.Medicion
// @Failure      400  {object}  errors.ErrorResponse "Invalid ID format"
// @Failure      404  {object}  errors.ErrorResponse "Measurement not found"
// @Router       /measurements/{id} [get]
func (h *MeasurementHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Measurement ID must be a valid number", err))
		return
	}
	medicion, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, medicion)
}

// @Summary      Update measurement
// @Description  Update an existing measurement
// @Tags         measurements
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id          path   int  true  "Measurement ID"
// @Param        measurement body   models.Medicion true "Updated measurement data"
// @Success      200  {object}  models.Medicion
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      404  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /measurements/{id} [put]
func (h *MeasurementHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Measurement ID must be a valid number", err))
		return
	}
	var medicion models.Medicion
	if err := c.ShouldBindJSON(&medicion); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}
	medicion.ID = uint(id)
	if err := h.uc.Update(&medicion); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, medicion)
}

// @Summary      Delete measurement
// @Description  Delete a measurement from the system
// @Tags         measurements
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "Measurement ID"
// @Success      204 "No Content"
// @Failure      400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure      500 {object} errors.ErrorResponse "Internal server error"
// @Router       /measurements/{id} [delete]
func (h *MeasurementHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Measurement ID must be a valid number", err))
		return
	}
	if err := h.uc.Delete(uint(id)); err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary      Get measurements by user
// @Description  Get all measurements of a specific user
// @Tags         measurements
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user_id  path      int  true  "User ID"
// @Success      200 {array}   models.Medicion
// @Failure      400 {object}  errors.ErrorResponse "Invalid user ID format"
// @Failure      500 {object}  errors.ErrorResponse "Internal server error"
// @Router       /measurements/user/{user_id} [get]
func (h *MeasurementHandler) GetByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_USER_ID", "User ID must be a valid number", err))
		return
	}
	mediciones, err := h.uc.GetByUserID(uint(userID))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, mediciones)
}
