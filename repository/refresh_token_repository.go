package repository

import (
	"acl-casbin/model"
	"acl-casbin/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RefreshTokenRepository interface {
	Store(uid primitive.ObjectID, refreshToken string, accessToken string, countOfUsage int, userAgent *utils.UserAgent, creationTime, expiresAt time.Time) error
	FindByToken(token string) (*model.RefreshToken, error)
	DeleteByUID(uid string) error
	EnsureIndexes() error
	DeleteByToken(token string) error
	FindByTokenWithUser(token string) (*model.RefreshTokenWithUser, error)
}
