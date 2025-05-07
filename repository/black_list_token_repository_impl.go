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

type BlackListTokenRepositoryImpl struct {
	collection *mongo.Collection
}

func NewBlackListRepository(db *mongo.Database) BlackListTokenRepository {
	return &BlackListTokenRepositoryImpl{
		collection: db.Collection(constant.BlackListTokenCollection),
	}
}
func (b BlackListTokenRepositoryImpl) Store(token *model.BlackListToken) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := b.collection.InsertOne(ctx, token)
	return err
}
func (b BlackListTokenRepositoryImpl) FindByToken(token string) (*model.BlackListToken, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var blackToken model.BlackListToken
	err := b.collection.FindOne(ctx, bson.M{"token": token}).Decode(&blackToken)
	if err != nil {
		return nil, err
	}
	return &blackToken, nil
}

func (b BlackListTokenRepositoryImpl) EnsureIndexes() error {
	// TTL index on ExpiresAt
	ttlIndex := mongo.IndexModel{
		Keys: bson.D{{Key: "expires_at", Value: 1}},
		Options: options.Index().
			SetExpireAfterSeconds(0).
			SetName("back_token_expires_at_ttl"),
	}
	_, err := b.collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{ttlIndex})
	return err
}
