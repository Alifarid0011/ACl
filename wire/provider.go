package wire

import (
	"acl-casbin/config"
	"acl-casbin/controller"
	"acl-casbin/repository"
	"acl-casbin/service"
	"acl-casbin/utils"
	"fmt"
	"github.com/casbin/casbin/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

func ProvideMongoClient() *mongo.Client {
	BaseUri := fmt.Sprintf("%s://%s:%s@", config.Get.Mongo.Protocol, config.Get.Mongo.Username, config.Get.Mongo.Password)
	for _, host := range config.Get.Mongo.Hosts {
		BaseUri += host + ":" + config.Get.Mongo.Port + ","
	}
	Uri := fmt.Sprintf(strings.TrimRight(BaseUri, ",")+"/%s?authSource="+config.Get.Mongo.AuthSource, config.Get.Mongo.DbName)
	return config.InitMongoClient(Uri)
}
func ProvideCasbinEnforcer(mongoClient *mongo.Client) *casbin.Enforcer {
	return config.InitCasbin(mongoClient)
}
func ProvideUserRepository(db *mongo.Database) repository.UserRepository {
	return repository.NewUserRepository(db)
}
func ProvideAuthService(
	userRepo repository.UserRepository,
	refreshTokenRepo repository.RefreshTokenRepository,
	tokenManager utils.JwtToken,
	blackListRepo repository.BlackListTokenRepository,
) service.AuthService {
	return service.NewAuthService(userRepo, tokenManager, refreshTokenRepo, blackListRepo)
}

func ProvideAuthController(authService service.AuthService, userService service.UserService) controller.AuthController {
	return controller.NewAuthController(authService, userService)
}

// ProvideUserService wires UserRepository into UserServiceImpl.
func ProvideUserService(userRepo repository.UserRepository) service.UserService {
	return service.NewUserService(userRepo)
}
func ProvideUserController(userService service.UserService) controller.UserController {
	return controller.NewUserController(userService)
}
func ProviderApprovalService(approvalRepo repository.ApprovalRepository) service.ApprovalService {
	return service.NewApprovalService(approvalRepo)
}
func ProviderApprovalController(approvalService service.ApprovalService) controller.ApprovalController {
	return controller.NewApprovalController(approvalService)
}
func ProviderApprovalRepository(db *mongo.Database) repository.ApprovalRepository {
	return repository.NewApprovalMongoRepository(db)
}
func ProviderBlackListRepository(db *mongo.Database) repository.BlackListTokenRepository {
	return repository.NewBlackListRepository(db)
}
func ProvideRefreshTokenRepository(db *mongo.Database) repository.RefreshTokenRepository {
	return repository.NewRefreshTokenRepository(db)
}
func ProvideJWT(secret string) utils.JwtToken {
	return utils.NewJwtToken(secret)
}
func ProvideDatabase(client *mongo.Client) *mongo.Database {
	return client.Database(config.Get.Mongo.DbName)
}
