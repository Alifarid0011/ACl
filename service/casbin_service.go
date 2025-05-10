package service

import (
	"acl-casbin/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CasbinService interface {
	IsAllowed(sub, act, obj string) (bool, error)
	GrantPermission(sub, act, obj string) (bool, error)
	RevokePermission(sub, act, obj string) (bool, error)
	ListPermissions() ([]model.CasbinPolicy, error)
	AddGrouping(parent string, child string, policyType string) error
	GetAllCasbinData() (map[string]interface{}, error)
	GetPermissionsBySubject() ([]model.SubjectWithPermissions, error)
	GetUserCategorizedPermissions(userID primitive.ObjectID) (model.CategorizedPermissions, error)
}
