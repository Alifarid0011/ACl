package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ACLPolicy struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Sub       string             `bson:"sub" json:"sub"` // Subject (نقش یا کاربر)
	Act       string             `bson:"act" json:"act"` // Action (مثل read, write, edit)
	Obj       string             `bson:"obj" json:"obj"` // Object (مثلاً invoice)
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	CreatedBy string             `bson:"created_by" json:"created_by"` // Optional: User who added the rule
}
