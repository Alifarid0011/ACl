package controller

import "github.com/gin-gonic/gin"

type ApprovalController interface {
	ApplyDecision(c *gin.Context)
	ListFlows(c *gin.Context)
	GetFlowByObject(c *gin.Context)
}
