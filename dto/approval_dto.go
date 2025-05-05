package dto

import "time"

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
