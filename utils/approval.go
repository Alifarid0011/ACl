package utils

import "fmt"

// Map for Approval Status
var approvalStatusMap = map[int]string{
	0: "Pending",
	1: "Approved",
	2: "Rejected",
}

func ApprovalStatusToString(status int) string {
	if s, ok := approvalStatusMap[status]; ok {
		return s
	}
	return fmt.Sprintf("Unknown(%d)", status)
}

// Map for Approval Action
var approvalActionMap = map[int]string{
	1: "Approve",
	2: "Reject",
}

func ApprovalActionToString(action int) string {
	if a, ok := approvalActionMap[action]; ok {
		return a
	}
	return fmt.Sprintf("Unknown(%d)", action)
}

// Map for Dependency Type
var dependencyTypeMap = map[string]string{
	"step":  "Step Dependency",
	"group": "Group Dependency",
	"role":  "Role Dependency",
}

func DependencyTypeToString(depType string) string {
	if d, ok := dependencyTypeMap[depType]; ok {
		return d
	}
	return fmt.Sprintf("Unknown(%s)", depType)
}
