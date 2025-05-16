package service

import (
	"acl-casbin/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CasbinService interface {
	IsAllowed(sub, obj, act, attr, AllowOrDeny string) (bool, error)
	GrantPermission(sub, obj, act, attr, AllowOrDeny string) (bool, error)
	RevokePermission(sub, obj, act, attr, AllowOrDeny string) (bool, error)
	ListPermissions() ([]model.CasbinPolicy, error)
	AddGrouping(parent string, child string) error
	GetAllCasbinData() (map[string]interface{}, error)
	GetPermissionsBySubject() ([]model.SubjectWithPermissions, error)
	GetUserCategorizedPermissions(userID primitive.ObjectID) (model.CategorizedPermissions, error)
}
