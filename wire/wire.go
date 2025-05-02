//go:build wireinject
// +build wireinject

package wire

import (
	"acl-casbin/controller"
	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Mongo    *mongo.Client
	Enforcer *casbin.Enforcer
	AuthCtrl controller.AuthController
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
		wire.Struct(new(App), "Mongo", "Enforcer", "AuthCtrl"),
	)
	return &App{}, nil
}
