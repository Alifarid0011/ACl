package service

import (
	"acl-casbin/constant"
	"acl-casbin/model"
	"acl-casbin/repository"
	"context"
	"time"
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

func (s *ApprovalServiceImpl) CreateFlow(ctx context.Context, flow model.ApprovalFlow) error {
	flow.CreatedAt = time.Now()
	flow.UpdatedAt = time.Now()
	flow.Status = constant.ApprovalStatusPending

	// پیش‌فرض: همه stepها Pending باشن
	for i := range flow.Steps {
		flow.Steps[i].Status = constant.ApprovalStatusPending
	}

	return s.repo.UpdateFlow(ctx, &flow)
}

func (s *ApprovalServiceImpl) UpdateStepStatus(ctx context.Context, objectID, objectType string, stepID int, status int) error {
	flow, err := s.repo.GetFlowByObject(ctx, objectID, objectType)
	if err != nil {
		return err
	}

	for i := range flow.Steps {
		if flow.Steps[i].StepID == stepID {
			flow.Steps[i].Status = status
			break
		}
	}

	flow.UpdatedAt = time.Now()
	return s.repo.UpdateFlow(ctx, flow)
}
