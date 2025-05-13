package repository

import (
	"acl-casbin/model"
	"github.com/casbin/casbin/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type casbinRepository struct {
	enforcer *casbin.Enforcer
	userRepo UserRepository
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

func (r *casbinRepository) GetPolicies() ([]model.CasbinPolicy, error) {
	rawPolicies, err := r.enforcer.GetPolicy()
	var policies []model.CasbinPolicy
	for _, p := range rawPolicies {
		if len(p) >= 3 {
			policies = append(policies, model.CasbinPolicy{
				Subject: p[0],
				Action:  p[1],
				Object:  p[2],
			})
		}
	}
	return policies, err
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

func (r *casbinRepository) GetPermissionsGroupedBySubject() ([]model.SubjectWithPermissions, error) {
	rawPolicies, err := r.enforcer.GetPolicy()
	grouped := make(map[string][]model.Permission)
	for _, p := range rawPolicies {
		if len(p) >= 3 {
			sub := p[0]
			perm := model.Permission{
				Action: p[1],
				Object: p[2],
			}
			grouped[sub] = append(grouped[sub], perm)
		}
	}
	var result []model.SubjectWithPermissions
	for sub, perms := range grouped {
		result = append(result, model.SubjectWithPermissions{
			Subject:     sub,
			Permissions: perms,
		})
	}
	return result, err
}

func determineCategory(path string) string {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	if len(parts) > 0 {
		return strings.Title(parts[0]) // "acl", "user", ...
	}
	return "Other"
}

func (r *casbinRepository) GetCategorizedPermissionsByUserID(userID primitive.ObjectID) (model.CategorizedPermissions, error) {
	//user, err := r.userRepo.FindByUID(userID)
	//if err != nil {
	//	return model.CategorizedPermissions{}, err
	//}
	//mainRole := user.Roles
	//allRoles, err := r.enforcer.GetImplicitRolesForUser(userID.Hex())
	//if err != nil {
	//	return model.CategorizedPermissions{}, err
	//}
	//roleSet := make(map[string]struct{})
	//roleSet[mainRole[0]] = struct{}{}
	//for _, role := range allRoles {
	//	roleSet[role] = struct{}{}
	//}
	//allPermissions := make(map[string][]model.Permission)
	//for role := range roleSet {
	//	perms, err := r.enforcer.GetImplicitPermissionsForUser(role)
	//	if err != nil {
	//		continue
	//	}
	//	for _, p := range perms {
	//		if len(p) >= 3 {
	//			allPermissions[role] = append(allPermissions[role], model.Permission{
	//				Action: p[1],
	//				Object: p[2],
	//			})
	//		}
	//	}
	//}
	//categorized := make(map[string][]model.Permission)
	//for _, perms := range allPermissions {
	//	for _, perm := range perms {
	//		cat := determineCategory(perm.Object)
	//		categorized[cat] = append(categorized[cat], perm)
	//	}
	//}
	return model.CategorizedPermissions{}, nil
}
