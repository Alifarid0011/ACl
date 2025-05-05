package service

import (
	"acl-casbin/model"
	"acl-casbin/repository"
	"context"
)

type ApprovalServiceImpl struct {
	repo repository.ApprovalRepository
}

func NewApprovalService(repo repository.ApprovalRepository) *ApprovalServiceImpl {
	return &ApprovalServiceImpl{
		repo: repo,
	}
}

func (s *ApprovalServiceImpl) ApplyDecision(ctx context.Context, objectID, objectType string, stepID int, decision model.ApprovalDecision) error {
	// ساده‌سازی شده: فقط ثبت تصمیم
	return s.repo.ApplyDecisionWithLogic(ctx, objectID, objectType, stepID, decision)
}

func (s *ApprovalServiceImpl) ListFlows(ctx context.Context, objectType string, status *int) ([]model.ApprovalFlow, error) {
	return s.repo.ListFlows(ctx, objectType, status)
}

func (s *ApprovalServiceImpl) GetFlowByObject(ctx context.Context, objectID, objectType string) (*model.ApprovalFlow, error) {
	return s.repo.GetFlowByObject(ctx, objectID, objectType)
}
