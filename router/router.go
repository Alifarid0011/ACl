package router

import (
	"acl-casbin/wire"
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *wire.App) *gin.Engine {
	r := gin.Default()
	RegisterSwaggerRoutes(r)
	RegisterAuthRoutes(r, app)
	RegisterUserRoutes(r, app)
	RegisterApproveRoutes(r, app)
	// RegisterAclRoutes(r, app)

	return r
}
