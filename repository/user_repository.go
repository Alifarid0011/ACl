package repository

import (
	"acl-casbin/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	FindByUID(uid primitive.ObjectID) (*model.User, error)
	Create(user *model.User) error
	GetAll() ([]model.User, error)
	Update(user *model.User) error // Add Update
	Delete(user *model.User) error // Add Delete
	EnsureIndexes() error
}
