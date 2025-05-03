package controller

import (
	"acl-casbin/service"
	"github.com/gin-gonic/gin"
)

type userControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userControllerImpl{userService: userService}
}

func (u userControllerImpl) FindByUsername(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u userControllerImpl) FindByUID(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u userControllerImpl) Create(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u userControllerImpl) GetAll(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u userControllerImpl) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u userControllerImpl) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u userControllerImpl) Me(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
