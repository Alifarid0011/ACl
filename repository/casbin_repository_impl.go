package repository

import "github.com/casbin/casbin/v2"

type casbinRepository struct {
	enforcer *casbin.Enforcer
}

func NewCasbinRepository(enforcer *casbin.Enforcer) CasbinRepository {
	return &casbinRepository{enforcer: enforcer}
}

func (r *casbinRepository) Enforce(sub, act, obj string) (bool, error) {
	return r.enforcer.Enforce(sub, act, obj)
}

func (r *casbinRepository) AddPolicy(sub, act, obj string) (bool, error) {
	added, err := r.enforcer.AddPolicy(sub, act, obj)
	if err == nil && added {
		_ = r.enforcer.SavePolicy()
	}
	return added, err
}

func (r *casbinRepository) RemovePolicy(sub, act, obj string) (bool, error) {
	removed, err := r.enforcer.RemovePolicy(sub, act, obj)
	if err == nil && removed {
		_ = r.enforcer.SavePolicy()
	}
	return removed, err
}

func (r *casbinRepository) GetPolicies() ([][]string, error) {
	return r.enforcer.GetPolicy()
}

func (r *casbinRepository) AddGroupingPolicy(child, parent, policyType string) (bool, error) {
	var added bool
	var err error

	switch policyType {
	case "g":
		added, err = r.enforcer.AddGroupingPolicy(child, parent)
	case "g2":
		added, err = r.enforcer.AddNamedGroupingPolicy("g2", child, parent)
	default:
		return false, nil
	}

	if err == nil && added {
		_ = r.enforcer.SavePolicy()
	}

	return added, err
}

func (r *casbinRepository) RemoveGroupingPolicy(child, parent, policyType string) (bool, error) {
	var removed bool
	var err error

	switch policyType {
	case "g":
		removed, err = r.enforcer.RemoveGroupingPolicy(child, parent)
	case "g2":
		removed, err = r.enforcer.RemoveNamedGroupingPolicy("g2", child, parent)
	default:
		return false, nil
	}

	if err == nil && removed {
		_ = r.enforcer.SavePolicy()
	}

	return removed, err
}

func (r *casbinRepository) GetGroupingPolicies(policyType string) ([][]string, error) {
	switch policyType {
	case "g":
		return r.enforcer.GetGroupingPolicy()
	case "g2":
		return r.enforcer.GetNamedGroupingPolicy("g2")
	default:
		return [][]string{}, nil
	}
}
