package middleware

import (
	"acl-casbin/constant"
	"acl-casbin/dto/response"
	"acl-casbin/utils"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func matchURL(path string, routeMap map[string]string) (string, bool) {
	for pattern := range routeMap {
		if util.KeyMatch2(path, pattern) {
			return pattern, true
		}
	}
	return "", false
}

func trimmer(s string, start, end int) string {
	if len(s) >= start+end {
		return s[start : len(s)-end]
	}
	return s
}

func RolesMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
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
		authorized := false
		for _, role := range roles {
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
			claims, IsClaimsOk := claimsVal.(utils.CustomClaims)
			if !IsClaimsOk {
				response.New(c).Message("خطا در بررسی دسترسی").
					MessageID("casbin.enforce.error").
					Status(http.StatusInternalServerError).
					Errors(err).
					Dispatch()
				return
			}
			AttrMap := claims.AttrMap
			matchURL(obj, AttrMap)
			claims.ParseAttr()
			ok, errEnforce := enforcer.Enforce(role, obj, act, attr)
			if errEnforce != nil {
				response.New(c).Message("خطا در بررسی دسترسی").
					MessageID("casbin.enforce.error").
					Status(http.StatusInternalServerError).
					Errors(err).
					Dispatch()
				return
			}
			if ok {
				authorized = true
				break
			}
		}
		if !authorized {
			response.New(c).Message("شما دسترسی لازم را ندارید").
				MessageID("permission.denied").
				Status(http.StatusForbidden).
				Dispatch()
			return
		}
		c.Set(constant.ContextRolesKey, roles)
		c.Next()
	}
}
