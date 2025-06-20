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

func NewUsuarioHandler(r *gin.Engine, usecase *usuario.UsuarioUsecase) {
	handler := &UsuarioHandler{usecase}

	r.GET("/usuarios", handler.ObtenerTodos)
	r.GET("/usuarios/:id", handler.Obtener)
	r.GET("/usuarios/email/:email", handler.ObtenerUsuarioPorEmail)
	r.POST("/usuarios", handler.Registrar)
	r.POST("/login", handler.Login)
	r.PUT("/usuarios/:id", handler.Actualizar)
	r.DELETE("/usuarios/:id", handler.Eliminar)
}

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

func (h *UsuarioHandler) Login(c *gin.Context) {
	var cred struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.usecase.Login(cred.Email, cred.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UsuarioHandler) ObtenerTodos(c *gin.Context) {
	usuarios, err := h.usecase.ObtenerTodosUsuarios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}

func (h *UsuarioHandler) Obtener(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	usuario, err := h.usecase.ObtenerUsuario(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (h *UsuarioHandler) ObtenerUsuarioPorEmail(c *gin.Context) {
	email := c.Param("email")

	usuario, err := h.usecase.ObtenerUsuarioPorEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

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

func (h *UsuarioHandler) Eliminar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.usecase.EliminarUsuario(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}
