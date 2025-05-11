package dto

type AttributeDTO struct {
	Name       string                 `json:"name" validate:"required"`          // نام اتربیوت
	Database   string                 `json:"database" validate:"required"`      // دیتابیس مورد نظر
	Collection string                 `json:"collection" validate:"required"`    // کالکشن یا جدول هدف
	Filters    map[string]interface{} `json:"filters" validate:"required,min=1"` // فیلترهای اعمالی
}
