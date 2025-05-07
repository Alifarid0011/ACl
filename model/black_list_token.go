package model

import (
	"acl-casbin/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BlackListToken struct {
	Token     string             `bson:"token" json:"token"`
	UserAgent *utils.UserAgent   `bson:"user_agent" json:"user_agent"`
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	ExpiresAt time.Time          `bson:"expires_at" json:"expires_at"`
}
