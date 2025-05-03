package service

import (
	"acl-casbin/config"
	"acl-casbin/constant"
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
	if err := utils.CompareHashAndPassword(user.Password, []byte(req.Password)); err != nil {
		return dto.LoginResponse{}, errors.New("invalid credentials")
	}
	if !s.hasRole(user.Roles, constant.RoleSuperAdmin) {
		return dto.LoginResponse{}, errors.New("unauthorized")
	}
	accessToken, err := s.tokenManager.GenerateAccessToken(time.Now().Add(config.Get.Token.ExpiryAccessToken*time.Minute).Unix(), user.UID, user.Roles)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("token generation failed: %w", err)
	}
	refreshToken, err := s.tokenManager.GenerateRefreshToken(time.Now().Add(config.Get.Token.ExpiryRefreshToken*time.Minute).Unix(), user.UID)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("refresh token generation failed: %w", err)
	}
	// ذخیره refresh token در Mongo
	if err := s.refreshRepo.Store(user.UID, refreshToken, time.Now().Add(config.Get.Token.ExpiryRefreshToken*time.Minute)); err != nil {
		return dto.LoginResponse{}, fmt.Errorf("storing refresh token failed: %w", err)
	}
	return dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       user.UID,
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
