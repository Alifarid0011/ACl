package router

import (
	"acl-casbin/middleware"
	"acl-casbin/wire"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, app *wire.App) {
	auth := r.Group("/auth", middleware.UserAgentMiddleware())
	{
		auth.POST("/login", app.AuthCtrl.Login)
		auth.POST("/refresh_token", app.AuthCtrl.UseRefreshToken)
		auth.GET("/logout", middleware.AuthMiddleware(app.BlackListRepo, app.TokenManager), app.AuthCtrl.Logout)
		auth.POST("/register", app.AuthCtrl.Register)
	}
}
