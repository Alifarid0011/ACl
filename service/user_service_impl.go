package service

import (
	"acl-casbin/dto"
	"acl-casbin/model"
	"acl-casbin/repository"
	"acl-casbin/utils"
	"errors"
)

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: userRepo}
}

func (s *UserServiceImpl) FindByUsername(username string) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	return mapToUserResponse(user), nil
}

func (s *UserServiceImpl) UpdateUser(uid string, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	// Find the user by UID
	user, err := s.userRepo.FindByUID(uid)
	if err != nil {
		return nil, err
	}
	// Update only fields that are non-zero in the request
	if err := utils.UpdateStruct(user, req); err != nil {
		return nil, err
	}
	// Pass the updated user object to the repository for saving
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}
	return mapToUserResponse(user), nil
}

func (s *UserServiceImpl) FindByUID(uid string) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByUID(uid)
	if err != nil {
		return nil, err
	}
	return mapToUserResponse(user), nil
}

func (s *UserServiceImpl) CreateUser(req dto.CreateUserRequest) (*dto.UserResponse, error) {
	_, err := s.userRepo.FindByUsername(req.Username)
	if err == nil {
		return nil, errors.New("username already exists")
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		UID:      utils.GenerateUID(),
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		FullName: req.FullName,
	}
	if errCreate := s.userRepo.Create(user); errCreate != nil {
		return nil, errCreate
	}
	return mapToUserResponse(user), nil
}

func (s *UserServiceImpl) GetAll() ([]dto.UserResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []dto.UserResponse
	for _, u := range users {
		result = append(result, *mapToUserResponse(&u))
	}
	return result, nil
}

func (s *UserServiceImpl) DeleteUser(uid string) error {
	// این متد باید به ریپازیتوری اضافه بشه
	return errors.New("delete not implemented in repository yet")
}

func (s *UserServiceImpl) Me(userID string) (*dto.UserResponse, error) {
	return s.FindByUID(userID)
}

func mapToUserResponse(u *model.User) *dto.UserResponse {
	return &dto.UserResponse{
		UID:      u.UID,
		Username: u.Username,
		Email:    u.Email,
		FullName: u.FullName,
	}
}
