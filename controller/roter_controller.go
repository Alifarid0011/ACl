package controller

import "github.com/gin-gonic/gin"

type RouteController interface {
	ListGroupedRoutes(c *gin.Context)
}
