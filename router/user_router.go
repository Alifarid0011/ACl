package router

import (
	"acl-casbin/wire"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, app *wire.App) {
	userRouter := r.Group("/user")
	{
		userRouter.POST("/create", app.UserCtrl.Create)
		userRouter.GET("/all", app.UserCtrl.GetAll)
		userRouter.GET("/me", app.UserCtrl.Me)
		userRouter.GET("/uid/:uid", app.UserCtrl.FindByUID)
		userRouter.GET("/username/:username", app.UserCtrl.FindByUsername)
		userRouter.PUT("/:uid", app.UserCtrl.Update)
		userRouter.DELETE("/:uid", app.UserCtrl.Delete)
	}
}
