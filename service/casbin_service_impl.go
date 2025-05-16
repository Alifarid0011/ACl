package service

import (
	"acl-casbin/model"
	"acl-casbin/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type casbinService struct {
	repo repository.CasbinRepository
}

func NewCasbinService(repo repository.CasbinRepository) CasbinService {
	return &casbinService{repo: repo}
}

func (s *casbinService) AddGrouping(parent string, child string) error {
	added, err := s.repo.AddGroupingPolicy(child, parent)
	if err != nil {
		return err
	}
	if !added {
		return fmt.Errorf("پالیسی قبلاً وجود داشته")
	}
	return nil
}
func (s *casbinService) IsAllowed(sub, act, obj, attr, AllowOrDeny string) (bool, error) {
	return s.repo.Enforce(sub, obj, act, attr, AllowOrDeny)
}

func (s *casbinService) GrantPermission(sub, obj, act, attr, AllowOrDeny string) (bool, error) {
	return s.repo.AddPolicy(sub, obj, act, attr, AllowOrDeny)
}

func (s *casbinService) RevokePermission(sub, act, obj, attr, AllowOrDeny string) (bool, error) {
	return s.repo.RemovePolicy(sub, obj, act, attr, AllowOrDeny)
}
func (s *casbinService) GetAllCasbinData() (map[string]interface{}, error) {
	policies, err := s.repo.GetPolicies()
	if err != nil {
		return nil, err
	}
	groupingPoliciesG, err := s.repo.GetGroupingPolicies()
	return map[string]interface{}{
		"policies":            policies,
		"grouping_policies_g": groupingPoliciesG,
	}, nil
}
func (s *casbinService) ListPermissions() ([]model.CasbinPolicy, error) {
	return s.repo.GetPolicies()
}

func (s *casbinService) GetPermissionsBySubject() ([]model.SubjectWithPermissions, error) {
	return s.repo.GetPermissionsGroupedBySubject()
}

func (s *casbinService) GetUserCategorizedPermissions(userID primitive.ObjectID) (model.CategorizedPermissions, error) {
	return s.repo.GetCategorizedPermissionsByUserID(userID)
}
