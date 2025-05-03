package router

import (
	"acl-casbin/wire"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, app *wire.App) {
	userRouter := r.Group("/user")
	{
		userRouter.POST("/create")
	}
}
