package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateUID() primitive.ObjectID {
	return primitive.NewObjectID()
}
