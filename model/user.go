package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	UID          primitive.ObjectID `bson:"uid" json:"uid"`
	Username     string             `bson:"username" json:"username"`
	FullName     string             `bson:"full_name" json:"full_name"`
	NationalCode string             `json:"national_code" bson:"national_code"`
	Email        string             `bson:"email" json:"email"`
	Mobile       string             `bson:"mobile" json:"mobile"`
	Password     []byte             `bson:"password,omitempty" json:"-"`
	Roles        string             `bson:"role" json:"role"`
	IsActive     bool               `bson:"is_active" json:"is_active"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}
