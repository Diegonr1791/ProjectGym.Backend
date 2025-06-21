package http

import (
	"net/http"

	"github.com/Diegonr1791/GymBro/internal/auth"
	usuario "github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase *usuario.UsuarioUsecase
	config  auth.JWTConfig
}

func NewAuthHandler(r *gin.Engine, usecase *usuario.UsuarioUsecase, cfg auth.JWTConfig) {
	handler := &AuthHandler{usecase, cfg}

	// Rutas públicas (sin autenticación)
	r.POST("/auth/login", handler.Login)
	r.POST("/auth/refresh", handler.RefreshToken)
	r.POST("/auth/logout", handler.Logout)
}

// LoginRequest representa la solicitud de login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse representa la respuesta de login
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
	} `json:"user"`
}

// RefreshRequest representa la solicitud de refresh
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// RefreshResponse representa la respuesta de refresh
type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// @Summary Iniciar sesión
// @Description Autentica un usuario y genera access token y refresh token
// @Tags autenticación
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "Credenciales de login"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 401 {object} map[string]interface{} "Credenciales inválidas"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Autenticar usuario
	user, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	// Generar access token
	accessToken, err := auth.GenerateJWT(user.ID, user.Email, h.config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	// Generar refresh token
	refreshToken, err := auth.GenerateRefreshToken(user.ID, h.config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar refresh token"})
		return
	}

	// Construir respuesta
	response := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	response.User.ID = user.ID
	response.User.Email = user.Email

	c.JSON(http.StatusOK, response)
}

// @Summary Refrescar token
// @Description Renueva el access token usando un refresh token válido
// @Tags autenticación
// @Accept json
// @Produce json
// @Param refresh body RefreshRequest true "Refresh token"
// @Success 200 {object} RefreshResponse
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 401 {object} map[string]interface{} "Refresh token inválido"
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Validar refresh token
	claims, err := auth.ValidateRefreshToken(req.RefreshToken, h.config)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token inválido: " + err.Error()})
		return
	}

	// Obtener usuario para generar nuevos tokens
	user, err := h.usecase.ObtenerUsuario(uint(claims.UserID))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Generar nuevo access token
	accessToken, err := auth.GenerateJWT(user.ID, user.Email, h.config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	// Generar nuevo refresh token
	refreshToken, err := auth.GenerateRefreshToken(user.ID, h.config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar refresh token"})
		return
	}

	response := RefreshResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Cerrar sesión
// @Description Invalida el refresh token (en producción, agregar a blacklist)
// @Tags autenticación
// @Accept json
// @Produce json
// @Param refresh body RefreshRequest true "Refresh token a invalidar"
// @Success 200 {object} map[string]interface{} "Sesión cerrada exitosamente"
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// En producción, aquí agregarías el refresh token a una blacklist
	// o lo invalidarías en la base de datos/Redis
	// Por ahora, solo validamos que sea un token válido

	_, err := auth.ValidateRefreshToken(req.RefreshToken, h.config)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token inválido"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sesión cerrada exitosamente"})
}
