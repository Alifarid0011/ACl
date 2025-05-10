package repository

import (
	"acl-casbin/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CasbinRepository interface {
	Enforce(sub, act, obj string) (bool, error)
	AddPolicy(sub, act, obj string) (bool, error)
	RemovePolicy(sub, act, obj string) (bool, error)
	GetPolicies() ([]model.CasbinPolicy, error)
	AddGroupingPolicy(child, parent, policyType string) (bool, error)
	RemoveGroupingPolicy(child, parent, policyType string) (bool, error)
	GetGroupingPolicies(policyType string) ([][]string, error)
	GetPermissionsGroupedBySubject() ([]model.SubjectWithPermissions, error)
	GetCategorizedPermissionsByUserID(userID primitive.ObjectID) (model.CategorizedPermissions, error)
}
