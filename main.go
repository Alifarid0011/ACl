package main

import (
	"acl-casbin/config"
	_ "acl-casbin/docs"
	"acl-casbin/router"
	"acl-casbin/wire"
	"fmt"
	"log"
	"os"
)

func init() {
	config.ExposeConfig(os.Getenv("APP_ENV"))
}
func main() {
	app, err := wire.InitializeApp(config.Get.Token.SecretKey)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}
	r := router.SetupRouter(app)
	if err := r.Run(fmt.Sprintf(":%v", config.Get.App.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
