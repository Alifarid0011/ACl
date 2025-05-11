package repository

import (
	"acl-casbin/constant"
	"acl-casbin/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AttributeRepositoryImp struct {
	collection *mongo.Collection
}

func NewAttributeRepositoryImp(db *mongo.Database) AttributeRepository {
	return &AttributeRepositoryImp{collection: db.Collection(constant.AttributesCollection)}
}
func (r *AttributeRepositoryImp) Add(attribute model.Attribute) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, attribute)
	return err
}

func (r *AttributeRepositoryImp) GetAll() ([]model.Attribute, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []model.Attribute
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *AttributeRepositoryImp) Update(attr *model.Attribute) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Update the user by UID
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"uid": attr.ID},
		bson.M{
			"$set": attr, // Update all fields in the user object
		},
	)
	if err != nil {
		return err
	}
	return nil
}
