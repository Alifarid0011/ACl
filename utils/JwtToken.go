package utils

import (
	"acl-casbin/constant"
	"github.com/golang-jwt/jwt/v5"
)

type JwtToken interface {
	GenerateAccessToken(Expiry int64, uid string, roles []string) (string, error)
	GenerateRefreshToken(Expiry int64, uid string) (string, error)
	ParseToken(tokenStr string) (*CustomClaims, error)
}
type CustomClaims struct {
	UID       string             `json:"uid"`
	TokenType constant.TokenType `json:"type"`
	jwt.RegisteredClaims
}
