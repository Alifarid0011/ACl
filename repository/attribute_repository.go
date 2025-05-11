package repository

import (
	"acl-casbin/model"
)

type AttributeRepository interface {
	GetAll() ([]model.Attribute, error)
	Add(attribute model.Attribute) error
	Update(attr *model.Attribute) error
}
