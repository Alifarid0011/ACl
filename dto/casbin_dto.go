package dto

type CheckPermissionDTO struct {
	Sub string `json:"sub" validate:"required"` // می‌تواند نقش یا یوزر باشد
	Act string `json:"act" validate:"required"` // متد مانند GET, POST, PUT, DELETE
	Obj string `json:"obj" validate:"required"` // مسیر مانند /user/all یا /approvals/:id
}

type GroupingDTO struct {
	Parent string `json:"parent" validate:"required"`          // نقش
	Child  string `json:"child" validate:"required"`           // یوزر یا منبع
	Type   string `json:"type" validate:"required,oneof=g g2"` // g برای نقش‌ها و g2 برای کاربران
}
