package controller

import "github.com/gin-gonic/gin"

type ApprovalController interface {
	ApplyDecision(c *gin.Context)
	ListFlows(c *gin.Context)
	GetFlowByObject(c *gin.Context)
	CreateFlow(c *gin.Context)
	UpdateStepStatus(c *gin.Context)
}
