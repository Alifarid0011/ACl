package repository

import (
	"acl-casbin/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type refreshTokenRepositoryImpl struct {
	collection *mongo.Collection
}

func NewRefreshTokenRepository(db *mongo.Database) RefreshTokenRepository {
	return &refreshTokenRepositoryImpl{
		collection: db.Collection("refresh_tokens"),
	}
}

func (r *refreshTokenRepositoryImpl) Store(uid, token string, expiresAt time.Time) error {
	rt := model.RefreshToken{
		Token:     token,
		UID:       uid,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
	}
	_, err := r.collection.InsertOne(context.Background(), rt)
	return err
}
func (r *refreshTokenRepositoryImpl) FindByToken(token string) (*model.RefreshToken, error) {
	var result model.RefreshToken
	err := r.collection.FindOne(context.Background(), bson.M{"token": token}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *refreshTokenRepositoryImpl) DeleteByUID(uid string) error {
	_, err := r.collection.DeleteMany(context.Background(), bson.M{"uid": uid})
	return err
}
