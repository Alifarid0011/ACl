package service

import "acl-casbin/dto"

type AuthService interface {
	Login(req dto.LoginRequest) (dto.LoginResponse, error)
}
