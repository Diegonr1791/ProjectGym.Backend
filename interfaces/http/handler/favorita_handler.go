package http

import (
	"net/http"
	"strconv"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	"github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct {
	uc *usecase.FavoriteUsecase
}

func NewFavoriteHandler(r gin.IRouter, uc *usecase.FavoriteUsecase) {
	h := &FavoriteHandler{uc}

	favoriteRoutes := r.Group("/favorites")
	{
		favoriteRoutes.POST("", h.Create)
		favoriteRoutes.GET("", h.GetAll)
		favoriteRoutes.GET("/:id", h.GetByID)
		favoriteRoutes.GET("/user/:user_id", h.GetByUserID)
		favoriteRoutes.PUT("/:id", h.Update)
		favoriteRoutes.DELETE("/:id", h.Delete)
	}
}

// @Summary      Create a new favorite
// @Description  Create a new favorite in the system
// @Tags         favorites
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        favorite body models.Favorita true "Favorite data"
// @Success      201  {object}  models.Favorita
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /favorites [post]
func (h *FavoriteHandler) Create(c *gin.Context) {
	var favorita models.Favorita
	if err := c.ShouldBindJSON(&favorita); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}
	if err := h.uc.Create(&favorita); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, favorita)
}

// @Summary      Get all favorites
// @Description  Get a complete list of favorites
// @Tags         favorites
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.Favorita
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /favorites [get]
func (h *FavoriteHandler) GetAll(c *gin.Context) {
	favoritas, err := h.uc.GetAll()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, favoritas)
}

// @Summary      Get favorite by ID
// @Description  Get a specific favorite by its ID
// @Tags         favorites
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Favorite ID"
// @Success      200  {object}  models.Favorita
// @Failure      400  {object}  errors.ErrorResponse "Invalid ID format"
// @Failure      404  {object}  errors.ErrorResponse "Favorite not found"
// @Router       /favorites/{id} [get]
func (h *FavoriteHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Favorite ID must be a valid number", err))
		return
	}
	favorita, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, favorita)
}

// @Summary      Update favorite
// @Description  Update an existing favorite
// @Tags         favorites
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path   int  true  "Favorite ID"
// @Param        favorite body   models.Favorita true "Updated favorite data"
// @Success      200  {object}  models.Favorita
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      404  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /favorites/{id} [put]
func (h *FavoriteHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Favorite ID must be a valid number", err))
		return
	}
	var favorita models.Favorita
	if err := c.ShouldBindJSON(&favorita); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}
	favorita.ID = uint(id)
	if err := h.uc.Update(&favorita); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, favorita)
}

// @Summary      Delete favorite
// @Description  Delete a favorite from the system
// @Tags         favorites
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "Favorite ID"
// @Success      204 "No Content"
// @Failure      400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure      500 {object} errors.ErrorResponse "Internal server error"
// @Router       /favorites/{id} [delete]
func (h *FavoriteHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "Favorite ID must be a valid number", err))
		return
	}
	if err := h.uc.Delete(uint(id)); err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary      Get favorites by user
// @Description  Get all favorites of a specific user
// @Tags         favorites
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user_id  path      int  true  "User ID"
// @Success      200 {array}   models.Favorita
// @Failure      400 {object}  errors.ErrorResponse "Invalid user ID format"
// @Failure      500 {object}  errors.ErrorResponse "Internal server error"
// @Router       /favorites/user/{user_id} [get]
func (h *FavoriteHandler) GetByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_USER_ID", "User ID must be a valid number", err))
		return
	}
	favoritas, err := h.uc.GetByUserID(uint(userID))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, favoritas)
}
