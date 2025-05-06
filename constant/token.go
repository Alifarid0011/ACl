package constant

type TokenType = string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

const (
	RefreshTokenType TokenType = "refresh_token"
	AccessTokenType  TokenType = "access_token"
	UserUid          TokenType = "user_uid"
)
