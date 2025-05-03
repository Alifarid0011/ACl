package repository

import (
	"acl-casbin/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &UserRepositoryImpl{
		collection: db.Collection("users"),
	}
}
func (r *UserRepositoryImpl) Delete(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Delete the user by UID
	_, err := r.collection.DeleteOne(ctx, bson.M{"uid": user.UID})
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Update(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Update the user by UID
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"uid": user.UID},
		bson.M{
			"$set": user, // Update all fields in the user object
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindByUID(uid string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"uid": uid}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Create(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepositoryImpl) GetAll() ([]model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []model.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
