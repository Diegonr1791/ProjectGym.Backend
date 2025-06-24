// interfaces/http/role_handler.go
package http

import (
	"net/http"
	"strconv"

	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	usecase "github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

// MessageResponse represents a simple message response
// @Description Simple message response
type MessageResponse struct {
	Message string `json:"message" example:"Operation completed successfully"`
}

// RoleHandler handles HTTP requests for role operations
type RoleHandler struct {
	roleUseCase *usecase.RoleUseCase
}

// NewRoleHandler creates a new role handler
func NewRoleHandler(r gin.IRouter, roleUseCase *usecase.RoleUseCase) {
	handler := &RoleHandler{
		roleUseCase: roleUseCase,
	}

	// Grouping role routes under "roles"
	roleRoutes := r.Group("/roles")
	{
		roleRoutes.GET("", handler.GetAllRoles)
		roleRoutes.GET("/all", handler.GetAllRolesWithDeleted)
		roleRoutes.GET("/system", handler.GetSystemRoles)
		roleRoutes.GET("/active", handler.GetActiveRoles)
		roleRoutes.POST("", handler.CreateRole)
		roleRoutes.GET("/:id", handler.GetRoleByID)
		roleRoutes.GET("/name/:name", handler.GetRoleByName)
		roleRoutes.PUT("/:id", handler.UpdateRole)
		roleRoutes.DELETE("/:id", handler.SoftDeleteRole)
		roleRoutes.DELETE("/:id/hard", handler.HardDeleteRole)
		roleRoutes.POST("/:id/restore", handler.RestoreRole)
	}
}

// CreateRole handles POST /roles
// @Summary Create a new role
// @Description Create a new role in the system
// @Tags roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param role body models.Role true "Role object"
// @Success 201 {object} models.Role
// @Failure 400 {object} errors.ErrorResponse
// @Failure 409 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /roles [post]
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.roleUseCase.CreateRole(c.Request.Context(), &role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

// GetRoleByID handles GET /roles/:id
// @Summary Get a role by ID
// @Description Get a role by its ID
// @Tags roles
// @Produce json
// @Security BearerAuth
// @Param id path int true "Role ID"
// @Success 200 {object} models.Role
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /roles/{id} [get]
func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := h.roleUseCase.GetRoleByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

// GetRoleByName handles GET /roles/name/:name
// @Summary Get a role by name
// @Description Get a role by its name
// @Tags roles
// @Produce json
// @Security BearerAuth
// @Param name path string true "Role name"
// @Success 200 {object} models.Role
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /roles/name/{name} [get]
func (h *RoleHandler) GetRoleByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role name is required"})
		return
	}

	role, err := h.roleUseCase.GetRoleByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

// GetAllRoles handles GET /roles
// @Summary Get all active roles
// @Description Get all active roles in the system
// @Tags roles
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Role
// @Failure 500 {object} errors.ErrorResponse
// @Router /roles [get]
func (h *RoleHandler) GetAllRoles(c *gin.Context) {
	roles, err := h.roleUseCase.GetAllRoles(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

// GetAllRolesWithDeleted handles GET /roles/all
// @Summary Get all roles including deleted ones
// @Description Get all roles in the system including deleted ones
// @Tags roles
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Role
// @Failure 500 {object} errors.ErrorResponse
// @Router /roles/all [get]
func (h *RoleHandler) GetAllRolesWithDeleted(c *gin.Context) {
	roles, err := h.roleUseCase.GetAllRolesWithDeleted(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

// UpdateRole handles PUT /roles/:id
// @Summary Update a role
// @Description Update an existing role
// @Tags roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Role ID"
// @Param role body models.Role true "Role object"
// @Success 200 {object} models.Role
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 409 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /roles/{id} [put]
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	role.ID = uint(id)
	if err := h.roleUseCase.UpdateRole(c.Request.Context(), &role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

// SoftDeleteRole handles DELETE /roles/:id
// @Summary Soft delete a role
// @Description Mark a role as deleted (soft delete)
// @Tags roles
// @Produce json
// @Security BearerAuth
// @Param id path int true "Role ID"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /roles/{id} [delete]
func (h *RoleHandler) SoftDeleteRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	if err := h.roleUseCase.SoftDeleteRole(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, MessageResponse{Message: "Role deleted successfully"})
}

// HardDeleteRole handles DELETE /roles/:id/hard
// @Summary Hard delete a role
// @Description Permanently delete a role from the system
// @Tags roles
// @Produce json
// @Security BearerAuth
// @Param id path int true "Role ID"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /roles/{id}/hard [delete]
func (h *RoleHandler) HardDeleteRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	if err := h.roleUseCase.HardDeleteRole(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, MessageResponse{Message: "Role permanently deleted"})
}

// RestoreRole handles POST /roles/:id/restore
// @Summary Restore a deleted role
// @Description Restore a soft-deleted role
// @Tags roles
// @Produce json
// @Security BearerAuth
// @Param id path int true "Role ID"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /roles/{id}/restore [post]
func (h *RoleHandler) RestoreRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	if err := h.roleUseCase.RestoreRole(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, MessageResponse{Message: "Role restored successfully"})
}

// GetSystemRoles handles GET /roles/system
// @Summary Get all system roles
// @Description Get all system roles in the system
// @Tags roles
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Role
// @Failure 500 {object} errors.ErrorResponse
// @Router /roles/system [get]
func (h *RoleHandler) GetSystemRoles(c *gin.Context) {
	roles, err := h.roleUseCase.GetSystemRoles(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

// GetActiveRoles handles GET /roles/active
// @Summary Get all active roles
// @Description Get all active roles in the system
// @Tags roles
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Role
// @Failure 500 {object} errors.ErrorResponse
// @Router /roles/active [get]
func (h *RoleHandler) GetActiveRoles(c *gin.Context) {
	roles, err := h.roleUseCase.GetActiveRoles(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}
