package router

import (
	"acl-casbin/wire"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, app *wire.App) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", app.AuthCtrl.Login)
	}
}
