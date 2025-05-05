package dto

import (
	"acl-casbin/model"
	"time"
)

// ApplyDecisionRequest تعریف درخواست برای تصمیم‌گیری روی مرحله
type ApplyDecisionRequest struct {
	ObjectID   string              `json:"object_id"`
	ObjectType string              `json:"object_type"`
	StepID     int                 `json:"step_id"`
	Decision   ApprovalDecisionDTO `json:"decision"`
}

// ApprovalDecisionDTO مدل ورودی تصمیم
type ApprovalDecisionDTO struct {
	By      string    `json:"by" validate:"required,uuid4"`
	Action  int       `json:"action" validate:"required,oneof=1 2"` // 1: Approve, 2: Reject
	Comment string    `json:"comment"`
	At      time.Time `json:"at" validate:"required"`
}

type CreateFlowRequest struct {
	ObjectID    string              `json:"object_id" validate:"required,uuid4"`
	ObjectType  string              `json:"object_type" validate:"required"`
	Steps       []ApprovalStepInput `json:"steps" validate:"required,min=1,dive"`
	FinalStepID int                 `json:"final_step_id" validate:"required,min=1"`
}

func (r *CreateFlowRequest) ToModelSteps() []model.ApprovalStep {
	steps := make([]model.ApprovalStep, len(r.Steps))
	for i, s := range r.Steps {
		steps[i] = model.ApprovalStep{
			StepID:       s.StepID,
			Name:         s.Name,
			Assignees:    s.Assignees,
			Dependencies: s.Dependencies,
			Required:     s.Required,
			Status:       0, // پیش‌فرض Pending
		}
	}
	return steps
}

type ApprovalStepInput struct {
	StepID       int                    `json:"step_id" validate:"required,min=1"`
	Name         string                 `json:"name" validate:"required"`
	Assignees    []string               `json:"assignees" validate:"required,min=1,dive,uuid4"`
	Dependencies []model.StepDependency `json:"dependencies"` // ولیدیشن اختیاری
	Required     int                    `json:"required" validate:"required,min=1"`
}

type UpdateStepStatusRequest struct {
	ObjectID   string `json:"object_id" validate:"required,uuid4"`
	ObjectType string `json:"object_type" validate:"required"`
	StepID     int    `json:"step_id" validate:"required,min=1"`
	Status     int    `json:"status" validate:"required,oneof=0 1 2"` // Pending, Approved, Rejected
}
