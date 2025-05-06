package service

import (
	"acl-casbin/dto"
	"acl-casbin/utils"
)

type AuthService interface {
	Login(req dto.LoginRequest, agent *utils.UserAgent) (dto.LoginResponse, error)
}
