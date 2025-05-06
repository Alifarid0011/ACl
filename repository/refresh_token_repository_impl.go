package repository

import (
	"acl-casbin/constant"
	"acl-casbin/model"
	"acl-casbin/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		collection: db.Collection(constant.RefreshTokenCollection),
	}
}

func (r *refreshTokenRepositoryImpl) Store(userUid primitive.ObjectID, refreshToken, accessToken string, countOfUsage int, userAgent *utils.UserAgent, creationTime, expiresAt time.Time) error {
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
	err := r.collection.FindOne(context.Background(), bson.M{constant.RefreshTokenType: token}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *refreshTokenRepositoryImpl) DeleteByUID(uid string) error {
	_, err := r.collection.DeleteMany(context.Background(), bson.M{constant.UserUid: uid})
	return err
}

func (r *refreshTokenRepositoryImpl) DeleteByToken(token string) error {
	_, err := r.collection.DeleteMany(context.Background(), bson.M{constant.RefreshTokenType: token})
	return err
}
func (r *refreshTokenRepositoryImpl) FindByTokenWithUser(token string) (*model.RefreshTokenWithUser, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"token": token}}},
		{{
			Key: "$lookup", Value: bson.M{
				"from":         "users",
				"localField":   "uid",
				"foreignField": "uid",
				"as":           "user",
			},
		}},
		{{Key: "$unwind", Value: "$user"}},
	}
	cursor, errCollection := r.collection.Aggregate(context.Background(), pipeline)
	if errCollection != nil {
		return nil, errCollection
	}
	defer cursor.Close(context.Background())
	if cursor.Next(context.Background()) {
		var result model.RefreshTokenWithUser
		if errDecode := cursor.Decode(&result); errDecode != nil {
			return nil, errDecode
		}
		return &result, nil
	}
	return nil, mongo.ErrNoDocuments
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
	accessTokenIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "access_token", Value: 1}},
		Options: options.Index().SetName("access_token_index").SetUnique(true),
	}
	refreshTokenIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "refresh_token", Value: 1}},
		Options: options.Index().SetName("refresh_token_index").SetUnique(true),
	}
	_, err := r.collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{ttlIndex, uidIndex, accessTokenIndex, refreshTokenIndex})
	return err
}
