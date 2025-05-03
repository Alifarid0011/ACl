package repository

import (
	"acl-casbin/model"
	"time"
)

type RefreshTokenRepository interface {
	Store(uid, token string, expiresAt time.Time) error
	FindByToken(token string) (*model.RefreshToken, error)
	DeleteByUID(uid string) error
	EnsureIndexes() error
}
