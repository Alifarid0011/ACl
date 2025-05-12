package service

import (
	"acl-casbin/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	FindByUsername(username string) (*dto.UserResponse, error)
	FindByUID(uid primitive.ObjectID) (*dto.UserResponse, error)
	CreateUser(req dto.CreateUserRequest) (*dto.UserResponse, error)
	GetAll() ([]dto.UserResponse, error)
	UpdateUser(uid primitive.ObjectID, req dto.UpdateUserRequest) (*dto.UserResponse, error)
	DeleteUser(uid primitive.ObjectID) error
	Me(userID primitive.ObjectID) (*dto.UserResponse, error)
}
