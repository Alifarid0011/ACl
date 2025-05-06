package router

import (
	"acl-casbin/middleware"
	"acl-casbin/wire"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, app *wire.App) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", middleware.UserAgentMiddleware(), app.AuthCtrl.Login)
	}
}
