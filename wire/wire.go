//go:build wireinject
// +build wireinject

package wire

import (
	"acl-casbin/controller"
	"acl-casbin/repository"
	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Mongo            *mongo.Client
	Enforcer         *casbin.Enforcer
	AuthCtrl         controller.AuthController
	RefreshTokenRepo repository.RefreshTokenRepository
	UserCtrl         controller.UserController
	UserRepo         repository.UserRepository
	ApproveCtrl      controller.ApprovalController
	ApproveRepo      repository.ApprovalRepository
	BlackListRepo    repository.BlackListTokenRepository
}

func InitializeApp(secret string) (*App, error) {
	wire.Build(
		ProvideMongoClient,
		ProvideCasbinEnforcer,
		ProvideUserRepository,
		ProvideRefreshTokenRepository,
		ProvideJWT,
		ProvideDatabase,
		ProvideAuthService,
		ProvideAuthController,
		ProviderApprovalService,
		ProviderApprovalController,
		ProviderApprovalRepository,
		ProvideUserService,
		ProvideUserController,
		ProviderBlackListRepository,
		wire.Struct(new(App), "Mongo", "Enforcer", "AuthCtrl", "UserCtrl", "UserRepo", "RefreshTokenRepo", "ApproveCtrl", "ApproveRepo", "BlackListRepo"),
	)
	return &App{}, nil
}
