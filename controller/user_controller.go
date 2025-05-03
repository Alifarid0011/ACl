package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	FindByUsername(ctx *gin.Context)
	FindByUID(ctx *gin.Context)
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Me(ctx *gin.Context)
}
