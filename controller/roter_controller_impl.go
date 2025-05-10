package controller

import (
	"acl-casbin/dto/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type routeControllerImpl struct {
	engine *gin.Engine
}

func NewRouteController(engine *gin.Engine) RouteController {
	return &routeControllerImpl{engine: engine}
}

type RouteInfo struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

// ListGroupedRoutes لیست دسته‌بندی‌شده‌ی اندپوینت‌ها
// @Summary دریافت لیست اندپوینت‌ها
// @Description این متد تمام مسیرها را همراه با متد HTTP آن‌ها دسته‌بندی‌شده بر اساس prefix بازمی‌گرداند.
// @Tags Route
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /routes/list [get]
func (ctl *routeControllerImpl) ListGroupedRoutes(c *gin.Context) {
	routes := ctl.engine.Routes()
	grouped := make(map[string][]RouteInfo)
	for _, r := range routes {
		prefix := getFirstPrefix(r.Path)
		grouped[prefix] = append(grouped[prefix], RouteInfo{
			Method: r.Method,
			Path:   r.Path,
		})
	}
	response.New(c).
		Message("لیست اندپوینت‌ها با موفقیت دریافت شد").
		MessageID("route.list.success").
		Status(http.StatusOK).
		Data(grouped).
		Dispatch()
}
func getFirstPrefix(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) > 1 {
		return strings.ToUpper(parts[1])
	}
	return "/"
}
