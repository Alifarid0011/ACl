package dto

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	UserID              string `json:"user_id"`
	AccessToken         string `json:"access_token"`
	RefreshToken        string `json:"refresh_token"`
	AccessTokenExpired  int64  `json:"access_token_expired"`
	RefreshTokenExpired int64  `json:"refresh_token_expired"`
}
