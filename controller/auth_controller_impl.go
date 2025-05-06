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

func (c *authControllerImpl) Logout(ctx *gin.Context) {

}
func (c *authControllerImpl) Register(ctx *gin.Context) {

}

func (c *authControllerImpl) UseRefreshToken(ctx *gin.Context) {

}
