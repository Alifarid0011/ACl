package service

import (
	"acl-casbin/model"
	"context"
)

type ApprovalService interface {
	ApplyDecision(ctx context.Context, objectID, objectType string, stepID int, decision model.ApprovalDecision) error
	ListFlows(ctx context.Context, objectType string, status *int) ([]model.ApprovalFlow, error)
	GetFlowByObject(ctx context.Context, objectID, objectType string) (*model.ApprovalFlow, error)
	CreateFlow(ctx context.Context, flow model.ApprovalFlow) error
	UpdateStepStatus(ctx context.Context, objectID, objectType string, stepID int, status int) error
}
