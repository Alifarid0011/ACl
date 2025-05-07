//go:build wireinject
// +build wireinject

package wire

import (
	"acl-casbin/controller"
	"acl-casbin/repository"
	"acl-casbin/service"
	"acl-casbin/utils"
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
	TokenManager     utils.JwtToken
	CasbinRepo       repository.CasbinRepository
	CasbinCtrl       controller.CasbinController
	CasbinService    service.CasbinService
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
		ProviderCasbinRepository,
		ProviderCasbinController,
		ProviderCasbinService,
		wire.Struct(new(App), "Mongo", "Enforcer", "AuthCtrl", "UserCtrl", "UserRepo", "RefreshTokenRepo", "ApproveCtrl", "ApproveRepo", "BlackListRepo", "TokenManager", "CasbinRepo", "CasbinCtrl", "CasbinService"),
	)
	return &App{}, nil
}
