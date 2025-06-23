package http

import (
	"net/http"
	"strconv"

	dto "github.com/Diegonr1791/GymBro/interfaces/http/dto"
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
		userRoutes.GET("/all", handler.GetAllIncludingDeleted)
		userRoutes.GET("/deleted", handler.GetDeleted)
		userRoutes.POST("", handler.Create)
		userRoutes.GET("/:id", handler.GetByID)
		userRoutes.PUT("/:id", handler.Update)
		userRoutes.DELETE("/:id", handler.Delete)
		userRoutes.POST("/:id/restore", handler.Restore)
		userRoutes.DELETE("/:id/permanent", handler.HardDelete)
		userRoutes.GET("/email/:email", handler.GetByEmail)
	}
}

// @Summary      Create a new user
// @Description  Register a new user in the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user body dto.CreateUserRequest true "User data"
// @Success      201  {object}  dto.UserResponse
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

	userResp := dto.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		RoleID:    u.RoleID,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	c.JSON(http.StatusCreated, userResp)
}

// @Summary      Get all active users
// @Description  Get a complete list of active users (excluding deleted ones)
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   dto.UserResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /users [get]
func (h *UsuarioHandler) GetAll(c *gin.Context) {
	usuarios, err := h.usecase.GetAllUsuarios()
	if err != nil {
		c.Error(err)
		return
	}

	userResps := make([]dto.UserResponse, 0, len(usuarios))
	for _, u := range usuarios {
		userResps = append(userResps, dto.UserResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			RoleID:    u.RoleID,
			IsActive:  u.IsActive,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	c.JSON(http.StatusOK, userResps)
}

// @Summary      Get all users including deleted
// @Description  Get a complete list of all users including deleted ones
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   dto.UserResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /users/all [get]
func (h *UsuarioHandler) GetAllIncludingDeleted(c *gin.Context) {
	usuarios, err := h.usecase.GetAllUsuariosIncludingDeleted()
	if err != nil {
		c.Error(err)
		return
	}

	userResps := make([]dto.UserResponse, 0, len(usuarios))
	for _, u := range usuarios {
		userResps = append(userResps, dto.UserResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			RoleID:    u.RoleID,
			IsActive:  u.IsActive,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	c.JSON(http.StatusOK, userResps)
}

// @Summary      Get deleted users
// @Description  Get a list of all deleted users
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   dto.UserResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /users/deleted [get]
func (h *UsuarioHandler) GetDeleted(c *gin.Context) {
	usuarios, err := h.usecase.GetDeletedUsuarios()
	if err != nil {
		c.Error(err)
		return
	}

	userResps := make([]dto.UserResponse, 0, len(usuarios))
	for _, u := range usuarios {
		userResps = append(userResps, dto.UserResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			RoleID:    u.RoleID,
			IsActive:  u.IsActive,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	c.JSON(http.StatusOK, userResps)
}

// @Summary      Get user by ID
// @Description  Get a specific user by their ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  dto.UserResponse
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

	userResp := dto.UserResponse{
		ID:        usuario.ID,
		Name:      usuario.Name,
		Email:     usuario.Email,
		RoleID:    usuario.RoleID,
		IsActive:  usuario.IsActive,
		CreatedAt: usuario.CreatedAt,
		UpdatedAt: usuario.UpdatedAt,
	}
	c.JSON(http.StatusOK, userResp)
}

// @Summary      Get user by email
// @Description  Get a specific user by their email
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        email path     string true "User Email"
// @Success      200   {object} dto.UserResponse
// @Failure      404   {object} errors.ErrorResponse "User not found"
// @Router       /users/email/{email} [get]
func (h *UsuarioHandler) GetByEmail(c *gin.Context) {
	email := c.Param("email")

	usuario, err := h.usecase.GetUsuarioByEmail(email)
	if err != nil {
		c.Error(err)
		return
	}

	userResp := dto.UserResponse{
		ID:        usuario.ID,
		Name:      usuario.Name,
		Email:     usuario.Email,
		RoleID:    usuario.RoleID,
		IsActive:  usuario.IsActive,
		CreatedAt: usuario.CreatedAt,
		UpdatedAt: usuario.UpdatedAt,
	}
	c.JSON(http.StatusOK, userResp)
}

// @Summary      Update user
// @Description  Update an existing user's data
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Param        user body      dto.UpdateUserRequest true "Updated user data"
// @Success      200  {object}  dto.UserResponse
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

	userResp := dto.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		RoleID:    u.RoleID,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	c.JSON(http.StatusOK, userResp)
}

// @Summary      Soft delete user
// @Description  Soft delete a user from the system (logical deletion)
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "User ID"
// @Success      204 "No Content"
// @Failure      400 {object} errors.ErrorResponse "Invalid ID format or user already deleted"
// @Failure      404 {object} errors.ErrorResponse "User not found"
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

// @Summary      Restore deleted user
// @Description  Restore a soft deleted user
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "User ID"
// @Success      200 {object} dto.UserResponse "User restored successfully"
// @Failure      400 {object} errors.ErrorResponse "Invalid ID format or user not deleted"
// @Failure      404 {object} errors.ErrorResponse "User not found"
// @Failure      500 {object} errors.ErrorResponse "Internal server error"
// @Router       /users/{id}/restore [post]
func (h *UsuarioHandler) Restore(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "User ID must be a valid number", err))
		return
	}

	if err := h.usecase.RestoreUsuario(uint(id)); err != nil {
		c.Error(err)
		return
	}

	// Get the restored user to return in response
	usuario, err := h.usecase.GetUsuarioByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	userResp := dto.UserResponse{
		ID:        usuario.ID,
		Name:      usuario.Name,
		Email:     usuario.Email,
		RoleID:    usuario.RoleID,
		IsActive:  usuario.IsActive,
		CreatedAt: usuario.CreatedAt,
		UpdatedAt: usuario.UpdatedAt,
	}
	c.JSON(http.StatusOK, userResp)
}

// @Summary      Hard delete user
// @Description  Permanently delete a user from the system (physical deletion)
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int  true  "User ID"
// @Success      204 "No Content"
// @Failure      400 {object} errors.ErrorResponse "Invalid ID format"
// @Failure      404 {object} errors.ErrorResponse "User not found"
// @Failure      500 {object} errors.ErrorResponse "Internal server error"
// @Router       /users/{id}/permanent [delete]
func (h *UsuarioHandler) HardDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_ID", "User ID must be a valid number", err))
		return
	}

	if err := h.usecase.HardDeleteUsuario(uint(id)); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
