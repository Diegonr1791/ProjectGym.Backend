package usecase

import (
	"errors"
	"time"

	"github.com/Diegonr1791/GymBro/internal/auth"
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	repository "github.com/Diegonr1791/GymBro/internal/domain/repositories"
	"gorm.io/gorm"
)

type RefreshTokenUsecase struct {
	rtRepo    repository.RefreshTokenRepository
	userRepo  repository.UsuarioRepository
	jwtConfig auth.JWTConfig
}

func NewRefreshTokenUsecase(rtRepo repository.RefreshTokenRepository, userRepo repository.UsuarioRepository, jwtConfig auth.JWTConfig) *RefreshTokenUsecase {
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
		return "", err
	}

	// Revoke all previous tokens for the user for added security
	if err := uc.rtRepo.RevokeByUserID(userID); err != nil {
		// Log the error but don't fail the login process
	}

	refreshToken := &model.RefreshToken{
		UserID:    userID,
		TokenHash: auth.HashToken(tokenString),
		ExpiresAt: time.Now().Add(time.Duration(uc.jwtConfig.GetRefreshMaxAge()) * time.Second),
	}

	if err := uc.rtRepo.Save(refreshToken); err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateAndRefresh validates a refresh token string and returns a new access token if valid.
func (uc *RefreshTokenUsecase) ValidateAndRefresh(refreshTokenString string) (string, error) {
	tokenHash := auth.HashToken(refreshTokenString)
	storedToken, err := uc.rtRepo.FindByTokenHash(tokenHash)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid or revoked refresh token")
		}
		return "", err // Internal server error
	}

	if storedToken.ExpiresAt.Before(time.Now()) {
		return "", errors.New("expired refresh token")
	}

	user, err := uc.userRepo.GetByID(storedToken.UserID)
	if err != nil {
		return "", errors.New("user associated with token not found")
	}

	return auth.GenerateJWT(user.ID, user.Email, uc.jwtConfig)
}

// Revoke revokes a given refresh token string.
func (uc *RefreshTokenUsecase) Revoke(refreshTokenString string) error {
	tokenHash := auth.HashToken(refreshTokenString)
	return uc.rtRepo.Revoke(tokenHash)
}
