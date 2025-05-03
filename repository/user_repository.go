package repository

import "acl-casbin/model"

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	FindByUID(uid string) (*model.User, error)
	Create(user *model.User) error
	GetAll() ([]model.User, error)
	Update(user *model.User) error // Add Update
	Delete(user *model.User) error // Add Delete
	EnsureIndexes() error
}
