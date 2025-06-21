package http

import (
	"net/http"

	"github.com/Diegonr1791/GymBro/internal/auth"
	"github.com/Diegonr1791/GymBro/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userUsecase *usecase.UsuarioUsecase
	rtUsecase   *usecase.RefreshTokenUsecase
	config      auth.JWTConfig
}

func NewAuthHandler(r gin.IRouter, userUsecase *usecase.UsuarioUsecase, rtUsecase *usecase.RefreshTokenUsecase, cfg auth.JWTConfig) {
	handler := &AuthHandler{
		userUsecase: userUsecase,
		rtUsecase:   rtUsecase,
		config:      cfg,
	}

	authRoutes := r.Group("/auth")
	authRoutes.POST("/login", handler.Login)
	authRoutes.POST("/refresh", handler.RefreshToken)
	authRoutes.POST("/logout", handler.Logout)
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	User        struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
	} `json:"user"`
}

// @Summary Iniciar sesión
// @Description Autentica un usuario, genera un access token y un refresh token en una cookie segura.
// @Tags autenticación
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "Credenciales de login"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} map[string]interface{} "Datos inválidos"
// @Failure 401 {object} map[string]interface{} "Credenciales inválidas"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	user, err := h.userUsecase.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	accessToken, err := auth.GenerateJWT(user.ID, user.Email, h.config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token de acceso"})
		return
	}

	refreshTokenString, err := h.rtUsecase.CreateAndStore(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar refresh token: " + err.Error()})
		return
	}

	c.SetCookie("refresh_token", refreshTokenString, h.config.GetRefreshMaxAge(), "/api/v1/auth", "localhost", true, true)

	response := LoginResponse{AccessToken: accessToken}
	response.User.ID = user.ID
	response.User.Email = user.Email
	c.JSON(http.StatusOK, response)
}

// @Summary Refrescar token
// @Description Renueva el access token usando un refresh token válido desde una cookie.
// @Tags autenticación
// @Produce json
// @Success 200 {object} LoginResponse "Nuevo access token"
// @Failure 400 {object} map[string]interface{} "Cookie no encontrada"
// @Failure 401 {object} map[string]interface{} "Refresh token inválido"
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshTokenString, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token cookie no encontrada"})
		return
	}

	newAccessToken, err := h.rtUsecase.ValidateAndRefresh(refreshTokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}

// @Summary Cerrar sesión
// @Description Invalida el refresh token del usuario.
// @Tags autenticación
// @Produce json
// @Success 200 {object} map[string]interface{} "Sesión cerrada exitosamente"
// @Failure 400 {object} map[string]interface{} "No hay sesión activa para cerrar"
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	refreshTokenString, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No hay sesión activa para cerrar"})
		return
	}

	_ = h.rtUsecase.Revoke(refreshTokenString)

	c.SetCookie("refresh_token", "", -1, "/api/v1/auth", "localhost", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "Sesión cerrada exitosamente"})
}
