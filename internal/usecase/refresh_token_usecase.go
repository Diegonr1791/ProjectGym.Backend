package usecase

import (
	"time"

	"github.com/Diegonr1791/GymBro/internal/auth"
	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	models "github.com/Diegonr1791/GymBro/internal/domain/models"
	repositories "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"github.com/pkg/errors"
)

type RefreshTokenUsecase struct {
	rtRepo    repositories.RefreshTokenRepository
	userRepo  repositories.UsuarioRepository
	jwtConfig auth.JWTConfig
}

func NewRefreshTokenUsecase(rtRepo repositories.RefreshTokenRepository, userRepo repositories.UsuarioRepository, jwtConfig auth.JWTConfig) *RefreshTokenUsecase {
	return &RefreshTokenUsecase{
		rtRepo:    rtRepo,
		userRepo:  userRepo,
		jwtConfig: jwtConfig,
	}
}

// CreateAndStore Generates, hashes, and stores a new refresh token for a user.
// It also revokes all previous tokens for that user.
func (uc *RefreshTokenUsecase) CreateAndStore(userID uint) (string, error) {
	tokenString, err := auth.GenerateRefreshToken(userID, uc.jwtConfig)
	if err != nil {
		return "", domainErrors.NewAppError(500, "JWT_GENERATE_REFRESH_TOKEN_FAILED", "Failed to generate refresh token", err)
	}

	// Revoke all previous tokens for the user for added security
	if err := uc.rtRepo.RevokeByUserID(userID); err != nil {
		// Log the error but don't fail the login process
		// This is a security enhancement, not a critical failure
	}

	refreshToken := &models.RefreshToken{
		UserID:    userID,
		TokenHash: auth.HashToken(tokenString),
		ExpiresAt: time.Now().Add(time.Duration(uc.jwtConfig.GetRefreshMaxAge()) * time.Second),
	}

	if err := uc.rtRepo.Save(refreshToken); err != nil {
		return "", domainErrors.NewAppError(500, "DB_SAVE_REFRESH_TOKEN_FAILED", "Failed to save refresh token to database", err)
	}

	return tokenString, nil
}

// ValidateAndRefresh validates a refresh token string and returns a new access token if valid.
func (uc *RefreshTokenUsecase) ValidateAndRefresh(refreshTokenString string) (string, error) {
	tokenHash := auth.HashToken(refreshTokenString)
	storedToken, err := uc.rtRepo.FindByTokenHash(tokenHash)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return "", domainErrors.NewAppError(401, "INVALID_REFRESH_TOKEN", "Invalid or revoked refresh token", err)
		}
		return "", domainErrors.NewAppError(500, "DB_GET_REFRESH_TOKEN_FAILED", "Failed to get refresh token from database", err)
	}

	if storedToken.ExpiresAt.Before(time.Now()) {
		return "", domainErrors.NewAppError(401, "EXPIRED_REFRESH_TOKEN", "Expired refresh token", errors.New("token expired"))
	}

	user, err := uc.userRepo.GetByID(storedToken.UserID)
	if err != nil {
		if errors.Is(err, domainErrors.ErrNotFound) {
			return "", domainErrors.NewAppError(401, "USER_NOT_FOUND", "User associated with token not found", err)
		}
		return "", domainErrors.NewAppError(500, "DB_GET_USER_FAILED", "Failed to get user from database", err)
	}

	accessToken, err := auth.GenerateJWT(user.ID, user.Email, user.RoleID, uc.jwtConfig)
	if err != nil {
		return "", domainErrors.NewAppError(500, "JWT_GENERATE_ACCESS_TOKEN_FAILED", "Failed to generate access token", err)
	}

	return accessToken, nil
}

// Revoke revokes a given refresh token string.
func (uc *RefreshTokenUsecase) Revoke(refreshTokenString string) error {
	tokenHash := auth.HashToken(refreshTokenString)
	if err := uc.rtRepo.Revoke(tokenHash); err != nil {
		return domainErrors.NewAppError(500, "DB_REVOKE_REFRESH_TOKEN_FAILED", "Failed to revoke refresh token from database", err)
	}
	return nil
}
