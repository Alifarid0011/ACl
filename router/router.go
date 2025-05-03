package router

import (
	"acl-casbin/wire"
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *wire.App) *gin.Engine {
	r := gin.Default()
	RegisterAuthRoutes(r, app)
	RegisterSwaggerRoutes(r)
	RegisterUserRoutes(r, app)
	// RegisterAclRoutes(r, app)

	return r
}
