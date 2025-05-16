package repository

import (
	"acl-casbin/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CasbinRepository interface {
	Enforce(sub, obj, act, attr, AllowOrDeny string) (bool, error)
	AddPolicy(sub, obj, act, attr, AllowOrDeny string) (bool, error)
	RemovePolicy(sub, obj, act, attr, AllowOrDeny string) (bool, error)
	GetPolicies() ([]model.CasbinPolicy, error)
	AddGroupingPolicy(child, parent string) (bool, error)
	RemoveGroupingPolicy(child, parent string) (bool, error)
	GetGroupingPolicies() ([][]string, error)
	GetPermissionsGroupedBySubject() ([]model.SubjectWithPermissions, error)
	GetCategorizedPermissionsByUserID(userID primitive.ObjectID) (model.CategorizedPermissions, error)
}
