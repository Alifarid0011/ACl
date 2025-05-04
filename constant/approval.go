package constant

// ApprovalStatus وضعیت کلی یک مرحله یا جریان اپروال
const (
	ApprovalStatusPending  = iota // 0
	ApprovalStatusApproved        // 1
	ApprovalStatusRejected        // 2
)

// ApprovalAction عمل انجام‌شده توسط کاربر
const (
	ApprovalActionApprove = 1 // تأیید
	ApprovalActionReject  = 2 // رد
)

// DependencyType نوع وابستگی مرحله‌ها
const (
	DependencyTypeStep  = "step"  // وابستگی به مرحله دیگر
	DependencyTypeGroup = "group" // وابستگی به گروه خاصی
	DependencyTypeRole  = "role"  // وابستگی به نقش خاصی
)
