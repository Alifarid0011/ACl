package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBRef struct {
	Ref string             `json:"$ref" bson:"$ref"`                   // نام کالکشن مثلاً "users"
	ID  primitive.ObjectID `json:"$id" bson:"$id"`                     // آیدی مستند هدف
	DB  string             `json:"$db,omitempty" bson:"$db,omitempty"` // نام دیتابیس (اختیاری)
}
