package middleware

import (
	"acl-casbin/constant"
	"acl-casbin/dto/response"
	"acl-casbin/utils"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func matchURL(path string, routeMap utils.AttributeMap) (string, bool) {
	for pattern := range routeMap {
		if util.KeyMatch2(path, pattern) {
			return pattern, true
		}
	}
	return "", false
}

func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, err := enforcer.GetRolesForUser(c.GetString("user_uid"))
		if err != nil {
			response.New(c).Message("عملیات با خطا مواجه شد.").
				MessageID("role.fetch.error").
				Status(http.StatusConflict).
				Errors(err).
				Dispatch()
			return
		}
		act := c.Request.Method
		obj := c.Request.URL.Path
		claimsVal, claimsExists := c.Get("claims")
		if !claimsExists {
			response.New(c).Message("خطا در بررسی دسترسی").
				MessageID("casbin.enforce.error").
				Status(http.StatusInternalServerError).
				Errors(err).
				Dispatch()
			return
		}
		claims, IsClaimsOk := claimsVal.(*utils.CustomClaims)
		if !IsClaimsOk {
			response.New(c).Message("خطا در بررسی دسترسی").
				MessageID("casbin.enforce.error").
				Status(http.StatusUnauthorized).
				Errors(err).
				Dispatch()
			return
		}
		AttrMap := claims.AttrMap
		RealObject, exist := matchURL(obj, AttrMap)
		if !exist {
			response.New(c).Message("خطا در بررسی دسترسی").
				MessageID("casbin.enforce.error").
				Status(http.StatusUnauthorized).
				Errors(errors.New("خطا در برسی سطح دسترسی")).
				Dispatch()
			return
		}
		authorized := false
		//Roll as a Subject
		for _, role := range roles {
			//if action was * attribute  most be *
			valueAttr, _ := claims.ParseAttr(RealObject, role, act)
			ok, errEnforce := enforcer.Enforce(role, obj, act, valueAttr)
			if errEnforce != nil {
				response.New(c).Message("خطا در بررسی دسترسی").
					MessageID("casbin.enforce.error").
					Status(http.StatusUnauthorized).
					Errors(errEnforce).
					Dispatch()
				return
			}
			if ok {
				authorized = true
				break
			}
		}
		if !authorized {
			//User as a Subject
			user := c.GetString("user_uid")
			valueAttr, _ := claims.ParseAttr(RealObject, user, act)
			ok, errEnforce := enforcer.Enforce(user, obj, act, valueAttr)
			if errEnforce != nil {
				response.New(c).Message("خطا در بررسی دسترسی").
					MessageID("casbin.enforce.error").
					Status(http.StatusUnauthorized).
					Errors(errEnforce).
					Dispatch()
				return
			} else if ok {
				c.Set(constant.ContextRolesKey, roles)
				c.Next()
			} else {
				response.New(c).Message("شما دسترسی لازم را ندارید").
					MessageID("casbin.enforce.Unauthorized").
					Status(http.StatusUnauthorized).
					Errors(errors.New("شما درسترسی لازم را ندارید")).
					Dispatch()
				return
			}
		}
		c.Set(constant.ContextRolesKey, roles)
		c.Next()
	}
}
