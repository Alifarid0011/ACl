package router

import (
	"acl-casbin/middleware"
	"acl-casbin/wire"
	"github.com/gin-gonic/gin"
)

func RegisterListRoutes(r *gin.Engine, app *wire.App) {
	router := r.Group("/routes", middleware.UserAgentMiddleware())
	{
		router.GET("/list", app.RouterCtr.ListGroupedRoutes)
	}
}
