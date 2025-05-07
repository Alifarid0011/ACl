package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CreateUserRequest struct {
	Username     string `json:"username" validate:"required,unique_username,min=3,max=50"`
	NationalCode string `json:"national_code" validate:"numeric,len=10"`
	Password     string `json:"password" validate:"required,password,min=6,max=100"`
	Email        string `json:"email" validate:"required,unique_email,email"`
	FullName     string `json:"full_name" validate:"omitempty,max=100"`
	Mobile       string `json:"mobile" validate:"required,unique_mobile,iran_mobile"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" validate:"omitempty,email"`
	FullName string `json:"full_name"`
}

type UserResponse struct {
	UID       primitive.ObjectID `json:"uid"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	FullName  string             `json:"full_name"`
	Mobile    string             `json:"mobile"`
	CreatedAt time.Time          `json:"created_at"`
}
