package utils

import (
	"acl-casbin/constant"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtToken interface {
	GenerateAccessToken(Expiry int64, uid primitive.ObjectID, roles []primitive.ObjectID) (string, error)
	GenerateRefreshToken(Expiry int64, uid primitive.ObjectID) (string, error)
	ParseToken(tokenStr string) (*CustomClaims, error)
}
type CustomClaims struct {
	UID       string             `json:"uid"`
	TokenType constant.TokenType `json:"type"`
	Roles     []string           `json:"roles"`
	jwt.RegisteredClaims
}
