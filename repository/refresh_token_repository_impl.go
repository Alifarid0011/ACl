package repository

import (
	"acl-casbin/model"
	"acl-casbin/utils"
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *refreshTokenRepositoryImpl) Store(userUid, refreshToken, accessToken string, countOfUsage int, userAgent *utils.UserAgent, creationTime, expiresAt time.Time) error {
	rt := model.RefreshToken{
		RefreshToken:    refreshToken,
		AccessToken:     accessToken,
		UserUid:         userUid,
		UserAgent:       userAgent,
		ExpiresAt:       expiresAt.UTC(),
		RefreshUseCount: countOfUsage,
		CreatedAt:       creationTime.UTC(),
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

func (r *refreshTokenRepositoryImpl) EnsureIndexes() error {
	// TTL index on ExpiresAt
	ttlIndex := mongo.IndexModel{
		Keys: bson.D{{Key: "expires_at", Value: 1}},
		Options: options.Index().
			SetExpireAfterSeconds(0).
			SetName("expires_at_ttl"),
	}
	// Normal index on UID
	uidIndex := mongo.IndexModel{
		Keys: bson.D{{Key: "uid", Value: 1}},
		Options: options.Index().
			SetName("uid_index").
			SetUnique(false),
	}
	_, err := r.collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{ttlIndex, uidIndex})
	return err
}
