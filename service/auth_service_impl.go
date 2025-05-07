package service

import (
	"acl-casbin/config"
	"acl-casbin/dto"
	"acl-casbin/repository"
	"acl-casbin/utils"
	"errors"
	"fmt"
	"time"
)

type AuthServiceImpl struct {
	userRepo     repository.UserRepository
	tokenManager utils.JwtToken
	refreshRepo  repository.RefreshTokenRepository
}

func NewAuthService(
	userRepo repository.UserRepository,
	tokenManager utils.JwtToken,
	refreshRepo repository.RefreshTokenRepository,
) AuthService {
	return &AuthServiceImpl{
		userRepo:     userRepo,
		tokenManager: tokenManager,
		refreshRepo:  refreshRepo,
	}
}
func (s *AuthServiceImpl) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("user not found: %w", err)
	}
	if errCompareHashAndPassword := utils.CompareHashAndPassword(user.Password, []byte(req.Password)); errCompareHashAndPassword != nil {
		return dto.LoginResponse{}, errors.New("invalid credentials")
	}
	accessTokenExpired := time.Now().Add(config.Get.Token.ExpiryAccessToken * time.Minute)
	refreshTokenExpired := time.Now().Add(config.Get.Token.ExpiryRefreshToken * time.Minute)
	accessToken, errGenerateAccessToken := s.tokenManager.GenerateAccessToken(accessTokenExpired.Unix(), user.UID, user.Roles)
	if errGenerateAccessToken != nil {
		return dto.LoginResponse{}, fmt.Errorf("token generation failed: %w", errGenerateAccessToken)
	}
	refreshToken, errGenerateRefreshToken := s.tokenManager.GenerateRefreshToken(refreshTokenExpired.Unix(), user.UID)
	if errGenerateRefreshToken != nil {
		return dto.LoginResponse{}, fmt.Errorf("refresh token generation failed: %w", errGenerateRefreshToken)
	}
	// ذخیره refresh token در Mongo
	if errRefreshRepo := s.refreshRepo.Store(user.UID, refreshToken, refreshTokenExpired); errRefreshRepo != nil {
		return dto.LoginResponse{}, fmt.Errorf("storing refresh token failed: %w", errRefreshRepo)
	}
	return dto.LoginResponse{
		AccessToken:         accessToken,
		RefreshToken:        refreshToken,
		UserID:              user.UID,
		AccessTokenExpired:  accessTokenExpired.Unix(),
		RefreshTokenExpired: refreshTokenExpired.Unix(),
	}, nil
}
func (s *AuthServiceImpl) hasRole(roles []string, target string) bool {
	for _, r := range roles {
		if r == target {
			return true
		}
	}
	return false
}
