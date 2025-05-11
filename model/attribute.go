package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Attribute struct {
	ID         primitive.ObjectID     `bson:"_id,omitempty"`
	Name       string                 `bson:"name"`
	Database   string                 `bson:"database"`
	Collection string                 `bson:"collection"`
	Filters    map[string]interface{} `bson:"filters"`
}
