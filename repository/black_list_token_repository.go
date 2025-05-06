package repository

import "acl-casbin/model"

type BlackListTokenRepository interface {
	Store(Token *model.BlackListToken) error
	FindByToken(token string) (*model.BlackListToken, error)
	EnsureIndexes() error
}
