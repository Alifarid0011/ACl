package repository

import (
	"acl-casbin/model"
	"acl-casbin/utils"
	"time"
)

type RefreshTokenRepository interface {
	Store(uid, refreshToken string, accessToken string, countOfUsage int, userAgent *utils.UserAgent, creationTime, expiresAt time.Time) error
	FindByToken(token string) (*model.RefreshToken, error)
	DeleteByUID(uid string) error
	EnsureIndexes() error
}
