package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	UserID              primitive.ObjectID `json:"user_id"`
	AccessToken         string             `json:"access_token"`
	RefreshToken        string             `json:"refresh_token"`
	AccessTokenExpired  int64              `json:"access_token_expired"`
	RefreshTokenExpired int64              `json:"refresh_token_expired"`
}
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type LogoutRequest struct {
	AccessToken string `json:"access_token" validate:"required"`
}
