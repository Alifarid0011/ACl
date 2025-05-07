package service

import (
	"acl-casbin/repository"
	"fmt"
)

type casbinService struct {
	repo repository.CasbinRepository
}

func NewCasbinService(repo repository.CasbinRepository) CasbinService {
	return &casbinService{repo: repo}
}

func (s *casbinService) AddGrouping(parent string, child string, policyType string) error {
	added, err := s.repo.AddGroupingPolicy(child, parent, policyType)
	if err != nil {
		return err
	}
	if !added {
		return fmt.Errorf("پالیسی قبلاً وجود داشته")
	}
	return nil
}
func (s *casbinService) IsAllowed(sub, act, obj string) (bool, error) {
	return s.repo.Enforce(sub, act, obj)
}

func (s *casbinService) GrantPermission(sub, act, obj string) (bool, error) {
	return s.repo.AddPolicy(sub, act, obj)
}

func (s *casbinService) RevokePermission(sub, act, obj string) (bool, error) {
	return s.repo.RemovePolicy(sub, act, obj)
}

func (s *casbinService) ListPermissions() ([][]string, error) {
	return s.repo.GetPolicies()
}
