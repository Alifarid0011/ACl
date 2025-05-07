package repository

import (
	"acl-casbin/constant"
	"acl-casbin/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type UserRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &UserRepositoryImpl{
		collection: db.Collection(constant.UserCollection),
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
func (r *UserRepositoryImpl) EnsureIndexes() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "username", Value: 1}}, Options: options.Index().SetName("idx_username").SetUnique(true)},
		{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetName("idx_email").SetUnique(true)},
		{Keys: bson.D{{Key: "mobile", Value: 1}}, Options: options.Index().SetName("idx_mobile").SetUnique(true)},
		{Keys: bson.D{{Key: "full_name", Value: 1}}, Options: options.Index().SetName("idx_full_name")},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index().SetName("idx_created_at")},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index().SetName("idx_updated_at")},
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index().SetName("uid_unique").SetUnique(true)},
	}
	_, err := r.collection.Indexes().CreateMany(ctx, indexes)
	return err
}
