package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ACLGroupingPolicy struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Parent    string             `bson:"parent" json:"parent"` // نقش یا گروه بالاتر (مثلاً "editor")
	Child     string             `bson:"child" json:"child"`   // کاربر یا منبع (مثلاً "user_abc123")
	Type      string             `bson:"type" json:"type"`     // "g" یا "g2"
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	CreatedBy string             `bson:"created_by" json:"created_by"` // Optional: User who added
}
