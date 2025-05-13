package utils

import (
	"acl-casbin/constant"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtToken interface {
	GenerateAccessToken(Expiry int64, uid primitive.ObjectID, AttrMap map[string]string) (string, error)
	GenerateRefreshToken(Expiry int64, uid primitive.ObjectID) (string, error)
	ParseToken(tokenStr string) (*CustomClaims, error)
}
type CustomClaims struct {
	UID       string             `json:"uid"`
	TokenType constant.TokenType `json:"type"`
	AttrMap   map[string]string  `json:"attr_map"` // key = sub:object:action, value = attribute (e.g., uid, *)
	jwt.RegisteredClaims
}

func (c *CustomClaims) ParseAttr(sub, obj, action string) (string, bool) {
	key := fmt.Sprintf("%s:%s:%s", sub, obj, action)
	value, exist := c.AttrMap[key]
	return value, exist
}
