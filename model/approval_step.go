package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ApprovalFlow struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ObjectID    string             `bson:"object_id"`     // شناسه آبجکت اصلی (مثلاً فاکتور، مرخصی)
	ObjectType  string             `bson:"object_type"`   // نوع آبجکت
	Steps       []ApprovalStep     `bson:"steps"`         // لیست مراحل اپروال
	Status      int                `bson:"status"`        // 0: Pending, 1: Approved, 2: Rejected
	FinalStepID int                `bson:"final_step_id"` // شماره مرحله نهایی
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type ApprovalStep struct {
	StepID       int                `bson:"step_id"`      // شماره مرحله
	Name         string             `bson:"name"`         // نام مرحله
	Assignees    []string           `bson:"assignees"`    // شناسه کاربران مجاز به تصمیم‌گیری
	Dependencies []StepDependency   `bson:"dependencies"` // وابستگی به مراحل دیگر
	Decisions    []ApprovalDecision `bson:"decisions"`    // لیست تصمیمات گرفته‌شده
	Required     int                `bson:"required"`     // حداقل تصمیم مثبت موردنیاز
	Status       int                `bson:"status"`       // 0: Pending, 1: Approved, 2: Rejected
}

type ApprovalDecision struct {
	By      string    `bson:"by"`     // user_id تصمیم‌گیرنده
	Action  int       `bson:"action"` // 1: Approve, 2: Reject
	Comment string    `bson:"comment,omitempty"`
	At      time.Time `bson:"at"` // زمان تصمیم
}

type StepDependency struct {
	Type    string `bson:"type"`    // نوع وابستگی: step, group, role
	Targets []int  `bson:"targets"` // لیست شماره مراحل وابسته (همه الزامی هستند)
}
