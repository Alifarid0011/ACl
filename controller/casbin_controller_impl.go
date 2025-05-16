package controller

import (
	"acl-casbin/dto"
	"acl-casbin/dto/response"
	"acl-casbin/service"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type ACLController struct {
	service service.CasbinService
}

func NewACLController(srv service.CasbinService) *ACLController {
	return &ACLController{service: srv}
}

// CheckPermission بررسی سطح دسترسی کاربر
// @Summary بررسی سطح دسترسی
// @Description بررسی می‌کند که آیا کاربر اجازه انجام عمل مشخصی را دارد یا خیر.
// @Tags ACL
// @Accept json
// @Produce json
// @Param request body dto.CheckPermissionDTO true "اطلاعات بررسی دسترسی"
// @Success 200 {object} response.Response
// @Failure 400,500 {object} response.Response
// @Router /acl/check [get]
func (ctl *ACLController) CheckPermission(ctx *gin.Context) {
	var req dto.CheckPermissionDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.New(ctx).Message("اطلاعات نامعتبر است").
			MessageID("acl.check.invalid").
			Status(http.StatusBadRequest).
			Errors(err).
			Dispatch()
		return
	}
	allowed, err := ctl.service.IsAllowed(req.Sub, req.Obj, req.Act, req.Attr, req.AllowOrDeny)
	if err != nil {
		response.New(ctx).Message("بررسی دسترسی با خطا مواجه شد").
			MessageID("acl.check.failed").
			Status(http.StatusInternalServerError).
			Errors(err).
			Dispatch()
		return
	}
	response.New(ctx).Message("بررسی انجام شد").
		MessageID("acl.check.success").
		Status(http.StatusOK).
		Data(gin.H{"allowed": allowed}).
		Dispatch()
}

// CreatePolicy تعریف مجوز جدید
// @Summary افزودن مجوز
// @Description مجوزی برای یک کاربر/گروه جهت انجام عملی بر یک شیء اضافه می‌کند.
// @Tags ACL
// @Accept json
// @Produce json
// @Param request body dto.CheckPermissionDTO true "اطلاعات مجوز"
// @Success 200 {object} response.Response
// @Failure 400,500 {object} response.Response
// @Router /acl/policy/create [post]
func (ctl *ACLController) CreatePolicy(ctx *gin.Context) {
	var req dto.CheckPermissionDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.New(ctx).Message("اطلاعات نامعتبر است").
			MessageID("acl.policy.create.invalid").
			Status(http.StatusBadRequest).
			Errors(err).
			Dispatch()
		return
	}
	added, err := ctl.service.GrantPermission(req.Sub, req.Obj, req.Act, req.Attr, req.AllowOrDeny)
	if err != nil || !added {
		response.New(ctx).Message("ثبت مجوز با خطا مواجه شد").
			MessageID("acl.policy.create.failed").
			Status(http.StatusInternalServerError).
			Errors(errors.New("تبت مجوز با خطا مواجه شد")).
			Dispatch()
		return
	}
	response.New(ctx).Message("مجوز با موفقیت ثبت شد").
		MessageID("acl.policy.create.success").
		Status(http.StatusCreated).
		Dispatch()
}

// ListAllCasbinData
// @Summary
// @Description
// @Tags ACL
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400,500 {object} response.Response
// @Router /acl/permission/list [get]
func (ctl *ACLController) ListAllCasbinData(ctx *gin.Context) {
	data, err := ctl.service.GetAllCasbinData()
	if err != nil {
		response.New(ctx).Message("خطا در دریافت داده‌های Casbin").
			MessageID("casbin.data.fetch.failed").
			Status(http.StatusInternalServerError).
			Errors(err).
			Dispatch()
		return
	}

	response.New(ctx).Message("تمامی سیاست‌ها و گروه‌بندی‌ها دریافت شد").
		MessageID("casbin.data.fetch.success").
		Status(http.StatusOK).
		Data(data).
		Dispatch()
}

// RemovePolicy حذف مجوز
// @Summary حذف مجوز
// @Description مجوز مشخص شده را برای کاربر/گروه حذف می‌کند.
// @Tags ACL
// @Accept json
// @Produce json
// @Param request body dto.CheckPermissionDTO true "اطلاعات مجوز برای حذف"
// @Success 200 {object} response.Response
// @Failure 400,500 {object} response.Response
// @Router /acl/policy/remove [delete]
func (ctl *ACLController) RemovePolicy(ctx *gin.Context) {
	var req dto.CheckPermissionDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.New(ctx).Message("اطلاعات نامعتبر است").
			MessageID("acl.policy.remove.invalid").
			Status(http.StatusBadRequest).
			Errors(err).
			Dispatch()
		return
	}
	removed, err := ctl.service.RevokePermission(req.Sub, req.Obj, req.Act, req.Attr, req.AllowOrDeny)
	if err != nil || !removed {
		response.New(ctx).Message("حذف مجوز با خطا مواجه شد").
			MessageID("acl.policy.remove.failed").
			Status(http.StatusInternalServerError).
			Errors(err).
			Dispatch()
		return
	}
	response.New(ctx).Message("مجوز با موفقیت حذف شد").
		MessageID("acl.policy.remove.success").
		Status(http.StatusOK).
		Dispatch()
}

// AddGroupingPolicy افزودن گروه‌بندی
// @Summary افزودن گروه‌بندی
// @Description یک نقش یا گروه به نقش/گروه دیگر تخصیص می‌دهد (گروه‌بندی).
// @Tags ACL
// @Accept json
// @Produce json
// @Param request body dto.GroupingDTO true "اطلاعات گروه‌بندی"
// @Success 200 {object} response.Response
// @Failure 400,500 {object} response.Response
// @Router /acl/grouping/add [post]
func (ctl *ACLController) AddGroupingPolicy(ctx *gin.Context) {
	var req dto.GroupingDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.New(ctx).Message("اطلاعات نامعتبر است").
			MessageID("acl.grouping.add.invalid").
			Status(http.StatusBadRequest).
			Errors(err).
			Dispatch()
		return
	}
	err := ctl.service.AddGrouping(req.Parent, req.Child)
	if err != nil {
		response.New(ctx).Message("افزودن گروه با خطا مواجه شد").
			MessageID("acl.grouping.add.failed").
			Status(http.StatusInternalServerError).
			Errors(err).
			Dispatch()
		return
	}
	response.New(ctx).Message("گروه‌بندی با موفقیت اضافه شد").
		MessageID("acl.grouping.add.success").
		Status(http.StatusOK).
		Dispatch()
}

// ListPermissionsBySubject
// @Summary
// @Description
// @Tags ACL
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400,500 {object} response.Response
// @Router /acl/permission/list [get]
func (ctl *ACLController) ListPermissionsBySubject(ctx *gin.Context) {
	data, err := ctl.service.GetPermissionsBySubject()
	if err != nil {
		response.New(ctx).Message("خطا در دریافت مجوزها").
			MessageID("casbin.permissions.fetch.failed").
			Status(http.StatusInternalServerError).
			Errors(err).
			Dispatch()
		return
	}
	response.New(ctx).Message("مجوزها با موفقیت دریافت شد").
		MessageID("casbin.permissions.fetch.success").
		Status(http.StatusOK).
		Data(data).
		Dispatch()
}

func (ctl *ACLController) GetMyPermissions(ctx *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(ctx.GetString("user_uid"))
	if err != nil {
		response.New(ctx).
			Message("شناسه کاربر نامعتبر است").
			MessageID("user.uid.invalid").
			Status(http.StatusBadRequest).
			Errors(err).
			Dispatch()
		return
	}

	data, err := ctl.service.GetUserCategorizedPermissions(userID)
	if err != nil {
		response.New(ctx).
			Message("خطا در دریافت مجوزهای کاربر").
			MessageID("casbin.user.permissions.failed").
			Status(http.StatusInternalServerError).
			Errors(err).
			Dispatch()
		return
	}

	response.New(ctx).
		Message("مجوزهای کاربر با موفقیت دریافت شد").
		MessageID("casbin.user.permissions.success").
		Status(http.StatusOK).
		Data(data).
		Dispatch()
}
