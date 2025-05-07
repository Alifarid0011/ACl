package service

import (
	"acl-casbin/config"
	"acl-casbin/dto"
	"acl-casbin/model"
	"acl-casbin/repository"
	"acl-casbin/utils"
	"errors"
	"fmt"
	"time"
)

type AuthServiceImpl struct {
	userRepo      repository.UserRepository
	tokenManager  utils.JwtToken
	refreshRepo   repository.RefreshTokenRepository
	blackListRepo repository.BlackListTokenRepository
}

func NewAuthService(
	userRepo repository.UserRepository,
	tokenManager utils.JwtToken,
	refreshRepo repository.RefreshTokenRepository,
	blackListRepo repository.BlackListTokenRepository,
) AuthService {
	return &AuthServiceImpl{
		userRepo:      userRepo,
		tokenManager:  tokenManager,
		refreshRepo:   refreshRepo,
		blackListRepo: blackListRepo,
	}
}
func (s *AuthServiceImpl) Login(req dto.LoginRequest, userAgent *utils.UserAgent) (dto.LoginResponse, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("user not found: %w", err)
	}
	if errCompareHashAndPassword := utils.CompareHashAndPassword(user.Password, []byte(req.Password)); errCompareHashAndPassword != nil {
		return dto.LoginResponse{}, errors.New("invalid credentials")
	}
	//if !s.hasRole(user.Roles, constant.RoleSuperAdmin) {
	//	return dto.LoginResponse{}, errors.New("unauthorized")
	//}
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
	if errRefreshRepo := s.refreshRepo.Store(user.UID, refreshToken, accessToken, 0, userAgent, time.Now(), refreshTokenExpired); errRefreshRepo != nil {
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
func (s *AuthServiceImpl) Logout(AccessToken string, userAgent *utils.UserAgent) error {
	token, err := s.refreshRepo.FindByAccessToken(AccessToken)
	if err != nil {
		return err
	}
	blackToken := &model.BlackListToken{
		Token:     token.AccessToken,
		UserAgent: userAgent,
		UserId:    token.UserUid,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(config.Get.Token.ExpiryAccessToken * time.Minute),
	}
	errBlackListRepo := s.blackListRepo.Store(blackToken)
	if errBlackListRepo != nil {
		return errBlackListRepo
	}
	errDeleteByRefreshToken := s.refreshRepo.DeleteByRefreshToken(token.RefreshToken)
	if errDeleteByRefreshToken != nil {
		return errDeleteByRefreshToken
	}
	return nil
}
func (s *AuthServiceImpl) UseRefreshToken(req dto.RefreshRequest, userAgent *utils.UserAgent) (dto.LoginResponse, error) {
	OldRefreshToken, errRefreshRepo := s.refreshRepo.FindByRefreshTokenWithUser(req.RefreshToken)
	if errRefreshRepo != nil {
		return dto.LoginResponse{}, fmt.Errorf("refresh token not found: %w", errRefreshRepo)
	}
	accessTokenExpired := time.Now().Add(config.Get.Token.ExpiryAccessToken * time.Minute)
	refreshTokenExpired := time.Now().Add(config.Get.Token.ExpiryRefreshToken * time.Minute)
	accessToken, errGenerateAccessToken := s.tokenManager.GenerateAccessToken(accessTokenExpired.Unix(), OldRefreshToken.UserUid, OldRefreshToken.User.Roles)
	if errGenerateAccessToken != nil {
		return dto.LoginResponse{}, fmt.Errorf("token generation failed: %w", errGenerateAccessToken)
	}
	NewRefreshTokenString, errGenerateRefreshToken := s.tokenManager.GenerateRefreshToken(refreshTokenExpired.Unix(), OldRefreshToken.User.UID)
	if errGenerateRefreshToken != nil {
		return dto.LoginResponse{}, fmt.Errorf("refresh token generation failed: %w", errGenerateRefreshToken)
	}
	if errRefreshRepoStore := s.refreshRepo.Store(OldRefreshToken.User.UID, NewRefreshTokenString, accessToken, OldRefreshToken.RefreshUseCount+1, userAgent, time.Now(), refreshTokenExpired); errRefreshRepoStore != nil {
		return dto.LoginResponse{}, fmt.Errorf("storing refresh token failed: %w", errRefreshRepoStore)
	} else {
		errDeleteByToken := s.refreshRepo.DeleteByRefreshToken(OldRefreshToken.RefreshToken.RefreshToken)
		if errDeleteByToken != nil {
			return dto.LoginResponse{}, errDeleteByToken
		}
	}
	return dto.LoginResponse{
		AccessToken:         accessToken,
		RefreshToken:        NewRefreshTokenString,
		UserID:              OldRefreshToken.User.UID,
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
