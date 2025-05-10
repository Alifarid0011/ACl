package router

import (
	"acl-casbin/middleware"
	"acl-casbin/wire"
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *wire.App) *gin.Engine {
	r := app.Engine
	r.Use(middleware.CORSMiddleware())
	RegisterSwaggerRoutes(r)
	RegisterAuthRoutes(r, app)
	RegisterUserRoutes(r, app)
	RegisterApproveRoutes(r, app)
	RegisterAclRoutes(r, app)
	RegisterListRoutes(r, app)
	return r
}
