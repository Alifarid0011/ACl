package middleware

import (
	"acl-casbin/constant"
	"acl-casbin/utils"
	"github.com/gin-gonic/gin"
)

func UserAgentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ua := new(utils.UserAgent).Constructor(c)
		c.Set(constant.UserAgentInfo, ua)
		c.Next()
	}
}
