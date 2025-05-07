package middleware

import (
	"acl-casbin/repository"
	"acl-casbin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(blackRepo repository.BlackListTokenRepository, tokenManager utils.JwtToken) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString, err := utils.ExtractBearerToken(authHeader)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: " + err.Error()})
			return
		}
		token, err := tokenManager.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: " + err.Error()})
			return
		}
		blackToken, errBlackRepo := blackRepo.FindByToken(tokenString)
		if blackToken != nil && errBlackRepo != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is in black list"})
			return
		}
		c.Set("user_uid", token.UID)
		c.Set("access_token", tokenString)
		c.Next()
	}
}
