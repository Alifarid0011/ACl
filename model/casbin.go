package model

type CheckPermissionDTO struct {
	Sub               DBRef  `json:"sub" validate:"required"`
	Act               string `json:"act" validate:"required"`
	Obj               string `json:"obj" validate:"required"`
	SubjectCollection string `json:"subject_collection" validate:"required,oneof=User "`
}

type GroupingDTO struct {
	Parent DBRef  `json:"parent" validate:"required"` // نقش
	Child  DBRef  `json:"child" validate:"required"`  // یوزر یا منبع
	Type   string `json:"type" validate:"required,oneof=g g2"`
}
