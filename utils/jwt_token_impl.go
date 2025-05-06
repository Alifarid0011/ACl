package utils

import (
	"acl-casbin/constant"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Jwt struct {
	secretKey string
}

func NewJwtToken(secret string) JwtToken {
	return &Jwt{secretKey: secret}
}
func (j *Jwt) GenerateAccessToken(Expiry int64, uid primitive.ObjectID, roles []string) (string, error) {
	claims := jwt.MapClaims{
		"uid":   uid,
		"roles": roles,
		"exp":   Expiry,
		"type":  constant.AccessToken,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *Jwt) GenerateRefreshToken(Expiry int64, uid primitive.ObjectID) (string, error) {
	claims := jwt.MapClaims{
		"uid":  uid,
		"exp":  Expiry,
		"type": constant.RefreshToken,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *Jwt) ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("invalid claims structure")
	}
	return claims, nil
}
