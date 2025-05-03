package service

import "acl-casbin/dto"

type UserService interface {
	FindByUsername(username string) (*dto.UserResponse, error)
	FindByUID(uid string) (*dto.UserResponse, error)
	CreateUser(req dto.CreateUserRequest) (*dto.UserResponse, error)
	GetAll() ([]dto.UserResponse, error)
	UpdateUser(uid string, req dto.UpdateUserRequest) (*dto.UserResponse, error)
	DeleteUser(uid string) error
	Me(userID string) (*dto.UserResponse, error)
}
