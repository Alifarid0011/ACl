package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`                                   // مثلاً "admin", "editor"
	Description string             `bson:"description,omitempty" json:"description,omitempty"` // توضیح نقش
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	CreatedBy   string             `bson:"created_by" json:"created_by"` // Optional
}
