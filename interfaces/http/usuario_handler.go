package http

import (
	"net/http"
	"strconv"

	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	usuario "github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UsuarioHandler struct {
	usecase *usuario.UsuarioUsecase
}

func NewUsuarioHandler(r gin.IRouter, usecase *usuario.UsuarioUsecase) {
	handler := &UsuarioHandler{usecase}

	r.GET("/usuarios", handler.ObtenerTodos)
	r.GET("/usuarios/:id", handler.Obtener)
	r.GET("/usuarios/email/:email", handler.ObtenerUsuarioPorEmail)
	r.POST("/usuarios", handler.Registrar)
	r.PUT("/usuarios/:id", handler.Actualizar)
	r.DELETE("/usuarios/:id", handler.Eliminar)
}

// @Summary Registrar nuevo usuario
// @Description Registra un nuevo usuario en el sistema
// @Tags usuarios
// @Accept json
// @Produce json
// @Param usuario body model.Usuario true "Datos del usuario"
// @Success 201 {object} model.Usuario
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /usuarios [post]
func (h *UsuarioHandler) Registrar(c *gin.Context) {
	var u model.Usuario
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.RegistrarUsuario(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar"})
		return
	}

	c.JSON(http.StatusCreated, u)
}

// @Summary Obtener todos los usuarios
// @Description Obtiene la lista completa de usuarios
// @Tags usuarios
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} model.Usuario
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /usuarios [get]
func (h *UsuarioHandler) ObtenerTodos(c *gin.Context) {
	usuarios, err := h.usecase.ObtenerTodosUsuarios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}

// @Summary Obtener usuario por ID
// @Description Obtiene un usuario específico por su ID
// @Tags usuarios
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID del usuario"
// @Success 200 {object} model.Usuario
// @Failure 404 {object} map[string]interface{} "Usuario no encontrado"
// @Router /usuarios/{id} [get]
func (h *UsuarioHandler) Obtener(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	usuario, err := h.usecase.ObtenerUsuario(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

// @Summary Obtener usuario por email
// @Description Obtiene un usuario específico por su email
// @Tags usuarios
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param email path string true "Email del usuario"
// @Success 200 {object} model.Usuario
// @Failure 404 {object} map[string]interface{} "Usuario no encontrado"
// @Router /usuarios/email/{email} [get]
func (h *UsuarioHandler) ObtenerUsuarioPorEmail(c *gin.Context) {
	email := c.Param("email")

	usuario, err := h.usecase.ObtenerUsuarioPorEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

// @Summary Actualizar usuario
// @Description Actualiza los datos de un usuario existente
// @Tags usuarios
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID del usuario"
// @Param usuario body model.Usuario true "Datos actualizados del usuario"
// @Success 200 {object} model.Usuario
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 404 {object} map[string]interface{} "Usuario no encontrado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /usuarios/{id} [put]
func (h *UsuarioHandler) Actualizar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.usecase.ObtenerUsuario(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	var u model.Usuario
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.ActualizarUsuario(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar"})
		return
	}

	c.JSON(http.StatusOK, u)
}

// @Summary Eliminar usuario
// @Description Elimina un usuario del sistema
// @Tags usuarios
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID del usuario"
// @Success 200 {object} map[string]interface{} "Usuario eliminado correctamente"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /usuarios/{id} [delete]
func (h *UsuarioHandler) Eliminar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.usecase.EliminarUsuario(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}
