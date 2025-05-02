package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RefreshToken struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Token     string             `bson:"token"`
	UID       string             `bson:"uid"`
	CreatedAt time.Time          `bson:"created_at"`
	ExpiresAt time.Time          `bson:"expires_at"`
}
