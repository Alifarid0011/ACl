package router

import (
	"acl-casbin/wire"
	"github.com/gin-gonic/gin"
)

func RegisterAclRoutes(r *gin.Engine, app *wire.App) {
	acl_router := r.Group("/acl")
	{
		print(acl_router)
	}
}
