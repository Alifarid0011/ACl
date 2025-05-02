package router

import (
	"acl-casbin/wire"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, app *wire.App) {
	user_router := r.Group("/user")
	{
		print(user_router)
	}
}
