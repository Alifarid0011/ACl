package utils

import (
	"acl-casbin/constant"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Jwt struct {
	secretKey string
}

func NewJwtToken(secret string) JwtToken {
	return &Jwt{secretKey: secret}
}
func (j *Jwt) GenerateAccessToken(expiry int64, uid primitive.ObjectID, AttrMap map[string]string) (string, error) {
	claims := CustomClaims{
		UID:       uid.Hex(),
		TokenType: constant.AccessToken,
		AttrMap:   AttrMap,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expiry, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *Jwt) GenerateRefreshToken(Expiry int64, uid primitive.ObjectID) (string, error) {
	claims := CustomClaims{
		UID:       uid.Hex(),
		TokenType: constant.RefreshToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(Expiry, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *Jwt) ParseToken(tokenStr string) (*CustomClaims, error) {
	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil && !token.Valid {
		return nil, fmt.Errorf("token parse error: %w", err)
	}
	return claims, nil
}
