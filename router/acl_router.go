package router

import (
	"acl-casbin/wire"
	"github.com/gin-gonic/gin"
)

func RegisterAclRoutes(r *gin.Engine, app *wire.App) {
	aclRouter := r.Group("/acl")
	{
		aclRouter.GET("/check", app.CasbinCtrl.CheckPermission)
		aclRouter.GET("/permission/list", app.CasbinCtrl.ListAllCasbinData)
		aclRouter.POST("/grouping/add", app.CasbinCtrl.AddGroupingPolicy)
		aclRouter.POST("/policy/create", app.CasbinCtrl.CreatePolicy)
		aclRouter.DELETE("/policy/remove", app.CasbinCtrl.RemovePolicy)
	}
}
