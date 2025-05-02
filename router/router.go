package router

import (
	"acl-casbin/wire"
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *wire.App) *gin.Engine {
	r := gin.Default()
	RegisterAuthRoutes(r, app)
	// بعداً می‌تونی اینا رو اضافه کنی
	// RegisterUserRoutes(r, app)
	// RegisterAclRoutes(r, app)

	return r
}
