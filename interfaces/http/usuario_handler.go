package http

import (
	"net/http"
	"strconv"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	"github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UsuarioHandler struct {
	usecase *usecase.UsuarioUsecase
}

func NewUsuarioHandler(r gin.IRouter, usecase *usecase.UsuarioUsecase) {
	handler := &UsuarioHandler{usecase}

	// Grouping user routes under "users"
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("", handler.GetAll)
		userRoutes.POST("", handler.Create)
		userRoutes.GET("/:id", handler.GetByID)
		userRoutes.PUT("/:id", handler.Update)
		userRoutes.DELETE("/:id", handler.Delete)
		userRoutes.GET("/email/:email", handler.GetByEmail)
	}
}

// @Summary      Create a new user
// @Description  Register a new user in the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body models.User true "User data"
// @Success      201  {object}  models.User
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      409  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /users [post]
func (h *UsuarioHandler) Create(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	if err := h.usecase.CreateUsuario(&u); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, u)
}

// @Summary      Get all users
// @Description  Get a complete list of users
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.User
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /users [get]
func (h *UsuarioHandler) GetAll(c *gin.Context) {
	usuarios, err := h.usecase.GetAllUsuarios()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, usuarios)
}

// @Summary      Get user by ID
// @Description  Get a specific user by their ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.User
// @Failure      400  {object}  errors.ErrorResponse "Invalid ID format"
// @Failure      404  {object}  errors.ErrorResponse "User not found"
// @Router       /users/{id} [get]
func (h *UsuarioHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "User ID must be a valid number", err))
		return
	}

	usuario, err := h.usecase.GetUsuarioByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, usuario)
}

// @Summary      Get user by email
// @Description  Get a specific user by their email
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        email path     string true "User Email"
// @Success      200   {object} models.User
// @Failure      404   {object} errors.ErrorResponse "User not found"
// @Router       /users/email/{email} [get]
func (h *UsuarioHandler) GetByEmail(c *gin.Context) {
	email := c.Param("email")

	usuario, err := h.usecase.GetUsuarioByEmail(email)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, usuario)
}

// @Summary      Update user
// @Description  Update an existing user's data
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Param        user body      models.User true "Updated user data"
// @Success      200  {object}  models.User
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      404  {object}  errors.ErrorResponse
// @Failure      409  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /users/{id} [put]
func (h *UsuarioHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "User ID must be a valid number", err))
		return
	}

	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body", err))
		return
	}

	u.ID = uint(id)

	if err := h.usecase.UpdateUsuario(&u); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, u)
}

// @Summary      Delete user
// @Description  Delete a user from the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "User ID"
// @Success      204 "No Content"
// @Failure      400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure      500 {object} errors.ErrorResponse "Internal server error"
// @Router       /users/{id} [delete]
func (h *UsuarioHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "User ID must be a valid number", err))
		return
	}

	if err := h.usecase.DeleteUsuario(uint(id)); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
