package model

import (
	"acl-casbin/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BlackListToken struct {
	Token     string             `json:"token"`
	UserAgent *utils.UserAgent   `json:"user_agent"`
	UserId    primitive.ObjectID `json:"user_id"`
	CreatedAt time.Time          `json:"created_at"`
	ExpiresAt time.Time          `json:"expires_at"`
}
