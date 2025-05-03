package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwaggerRoutes(r *gin.Engine) {
	swaggerRouter := r.Group("/swagger")
	{
		swaggerRouter.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
