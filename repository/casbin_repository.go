package repository

type CasbinRepository interface {
	Enforce(sub, act, obj string) (bool, error)
	AddPolicy(sub, act, obj string) (bool, error)
	RemovePolicy(sub, act, obj string) (bool, error)
	GetPolicies() ([][]string, error)
	AddGroupingPolicy(child, parent, policyType string) (bool, error)
	RemoveGroupingPolicy(child, parent, policyType string) (bool, error)
	GetGroupingPolicies(policyType string) ([][]string, error)
}
