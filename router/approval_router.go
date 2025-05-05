package router

import (
	"acl-casbin/wire"
	"github.com/gin-gonic/gin"
)

func RegisterApproveRoutes(r *gin.Engine, app *wire.App) {
	approval := r.Group("/approvals")
	{
		approval.POST("/decision", app.ApproveCtrl.ApplyDecision)
		approval.GET("/", app.ApproveCtrl.ListFlows)
		approval.GET("/:object_type/:object_id", app.ApproveCtrl.GetFlowByObject)
		approval.POST("/", app.ApproveCtrl.CreateFlow)
		approval.PUT("/step", app.ApproveCtrl.UpdateStepStatus)
	}
}
