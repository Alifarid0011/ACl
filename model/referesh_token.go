package model

import (
	"acl-casbin/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RefreshToken struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	RefreshToken    string             `bson:"refresh_token"`
	AccessToken     string             `bson:"access_token"`
	RefreshUseCount int                `bson:"refresh_use_count"`
	UserUid         primitive.ObjectID `bson:"user_uid"`
	UserAgent       *utils.UserAgent   `bson:"user_agent"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
	ExpiresAt       time.Time          `bson:"expires_at"`
}

type RefreshTokenWithUser struct {
	RefreshToken `bson:",inline"`
	User         User `bson:"user"`
}
