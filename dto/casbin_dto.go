package dto

import (
	"acl-casbin/config"
	"acl-casbin/constant"
)

type CheckPermissionDTO struct {
	Sub               string `json:"sub" validate:"required"`
	Act               string `json:"act" validate:"required"`
	Obj               string `json:"obj" validate:"required"`
	SubjectCollection string `json:"subject_collection" validate:"required,oneof=User "`
}

type GroupingDTO struct {
	Parent string `json:"parent" validate:"required"` // نقش
	Child  string `json:"child" validate:"required"`  // یوزر یا منبع
	Type   string `json:"type" validate:"required,oneof=g g2"`
}

var SubjectMapCollection = map[string]map[string]string{
	"User": {
		"db":         config.Get.Mongo.DbName,
		"collection": constant.UserCollection,
	},
}
