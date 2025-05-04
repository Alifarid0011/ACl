package repository

import (
	"acl-casbin/model"
	"context"
)

type ApprovalRepository interface {
	ApplyDecisionWithLogic(ctx context.Context, objectID, objectType string, stepID int, decision model.ApprovalDecision) error
	ListFlows(ctx context.Context, objectType string, status *int) ([]model.ApprovalFlow, error)
	EnsureIndexes(ctx context.Context) error
	UpdateFlow(ctx context.Context, flow *model.ApprovalFlow) error
	GetFlowByObject(ctx context.Context, objectID, objectType string) (*model.ApprovalFlow, error)
}
