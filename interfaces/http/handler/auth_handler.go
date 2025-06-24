package http

import (
	"net/http"
	"regexp"

	"github.com/Diegonr1791/GymBro/internal/auth"
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
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
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	User        struct {
		ID     uint   `json:"id"`
		Email  string `json:"email"`
		Role   string `json:"role"`
		RoleID uint   `json:"role_id"`
	} `json:"user"`
}

// validateLoginRequest validates login request with custom error codes
func (h *AuthHandler) validateLoginRequest(req *LoginRequest) error {
	// Validate email
	if req.Email == "" {
		return domainErrors.ErrEmailRequired
	}

	// Validate email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(req.Email) {
		return domainErrors.ErrInvalidEmailFormat
	}

	// Validate password
	if req.Password == "" {
		return domainErrors.ErrPasswordRequired
	}

	return nil
}

// @Summary      Login
// @Description  Authenticates a user, generates an access token and a refresh token in a secure cookie.
// @Tags         authentication
// @Accept       json
// @Produce      json
// @Param        credentials body LoginRequest true "Login credentials"
// @Success      200  {object}  LoginResponse
// @Failure      400  {object}  errors.ErrorResponse "Invalid data"
// @Failure      401  {object}  errors.ErrorResponse "Invalid credentials"
// @Failure      500  {object}  errors.ErrorResponse "Internal server error"
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// Handle JSON binding errors with custom error codes
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Invalid JSON format", err))
		return
	}

	// Validate request with custom validation
	if err := h.validateLoginRequest(&req); err != nil {
		c.Error(err)
		return
	}

	user, err := h.userUsecase.Login(req.Email, req.Password)
	if err != nil {
		c.Error(err)
		return
	}

	accessToken, err := auth.GenerateJWT(user.ID, user.Email, user.RoleID, h.config)
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusInternalServerError, "JWT_GENERATION_FAILED", "Failed to generate access token", err))
		return
	}

	refreshTokenString, err := h.rtUsecase.CreateAndStore(user.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.SetCookie("refresh_token", refreshTokenString, h.config.GetRefreshMaxAge(), "/api/v1/auth", "localhost", true, true)

	response := LoginResponse{AccessToken: accessToken}
	response.User.ID = user.ID
	response.User.Email = user.Email
	response.User.Role = user.Role.Name
	response.User.RoleID = user.RoleID
	c.JSON(http.StatusOK, response)
}

// @Summary      Refresh token
// @Description  Renews the access token using a valid refresh token from a cookie.
// @Tags         authentication
// @Produce      json
// @Success      200  {object}  map[string]string "New access token"
// @Failure      400  {object}  errors.ErrorResponse "Cookie not found"
// @Failure      401  {object}  errors.ErrorResponse "Invalid refresh token"
// @Router       /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshTokenString, err := c.Cookie("refresh_token")
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "REFRESH_TOKEN_COOKIE_NOT_FOUND", "Refresh token cookie not found", err))
		return
	}

	newAccessToken, err := h.rtUsecase.ValidateAndRefresh(refreshTokenString)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}

// @Summary      Logout
// @Description  Invalidates the user's refresh token.
// @Tags         authentication
// @Produce      json
// @Success      200  {object}  map[string]string "Session closed successfully"
// @Failure      400  {object}  errors.ErrorResponse "No active session to close"
// @Router       /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	refreshTokenString, err := c.Cookie("refresh_token")
	if err != nil {
		c.Error(domainErrors.NewAppError(http.StatusBadRequest, "NO_ACTIVE_SESSION", "No active session to close", err))
		return
	}

	_ = h.rtUsecase.Revoke(refreshTokenString)

	c.SetCookie("refresh_token", "", -1, "/api/v1/auth", "localhost", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "Session closed successfully"})
}
