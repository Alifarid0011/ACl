package main

import (
	"acl-casbin/config"
	_ "acl-casbin/docs"
	"acl-casbin/router"
	"acl-casbin/utils"
	"acl-casbin/validation"
	"acl-casbin/wire"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"os"
)

func init() {
	config.ExposeConfig(os.Getenv("APP_ENV"))
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample API for demonstrating Swagger with Bearer Authentication in Go using Gin
// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	validation.Init()
	binding.Validator = &validation.DefaultValidator{}
	app, errInitializeApp := wire.InitializeApp(config.Get.Token.SecretKey)
	if errInitializeApp != nil {
		log.Fatalf("Failed to initialize app: %v", errInitializeApp)
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		utils.RegisterCustomValidators(v, app.Mongo)
	}
	EnsureIndexes(app)
	r := router.SetupRouter(app)
	if err := r.Run(fmt.Sprintf(":%v", config.Get.App.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func EnsureIndexes(app *wire.App) {
	if err := app.UserRepo.EnsureIndexes(); err != nil {
		log.Fatalf("Failed to create indexes: %v", err)
	}
	if err := app.RefreshTokenRepo.EnsureIndexes(); err != nil {
		log.Fatalf("Failed to create indexes: %v", err)
	}
	if err := app.BlackListRepo.EnsureIndexes(); err != nil {
		log.Fatalf("Failed to create indexes: %v", err)
	}
}
