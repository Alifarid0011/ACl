package service

type CasbinService interface {
	IsAllowed(sub, act, obj string) (bool, error)
	GrantPermission(sub, act, obj string) (bool, error)
	RevokePermission(sub, act, obj string) (bool, error)
	ListPermissions() ([][]string, error)
	AddGrouping(parent string, child string, policyType string) error
}
