package controller

import (
	"acl-casbin/constant"
	"acl-casbin/dto"
	"acl-casbin/service"
	"acl-casbin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authControllerImpl struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authControllerImpl{authService: authService}
}

// Login godoc
// @Summary      Authenticate user
// @Description  Takes username and password, returns access and refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginRequest true "Login credentials"
// @Success      200 {object} dto.LoginResponse
// @Router       /auth/login [post]
func (c *authControllerImpl) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	UserAgentInfo, _ := ctx.Get(constant.UserAgentInfo)
	resp, err := c.authService.Login(req, UserAgentInfo.(*utils.UserAgent))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// Logout godoc
// @Summary Logout and invalidate refresh token
// @Description Logs out the user and invalidates the refresh token based on user-agent
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LogoutRequest true "Logout request payload"
// @Router /auth/logout [post]
func (c *authControllerImpl) Logout(ctx *gin.Context) {
	var req dto.LogoutRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	UserAgentInfo, _ := ctx.Get(constant.UserAgentInfo)
	err := c.authService.Logout(req, UserAgentInfo.(*utils.UserAgent))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "Success"})
}
func (c *authControllerImpl) Register(ctx *gin.Context) {

}

// UseRefreshToken godoc
// @Summary Use refresh token to get new access token
// @Description Uses a refresh token and user-agent info to generate a new access token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshRequest true "Refresh token payload"
// @Router /auth/refresh [post]
func (c *authControllerImpl) UseRefreshToken(ctx *gin.Context) {
	var req dto.RefreshRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	UserAgentInfo, _ := ctx.Get(constant.UserAgentInfo)
	resp, err := c.authService.UseRefreshToken(req, UserAgentInfo.(*utils.UserAgent))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
