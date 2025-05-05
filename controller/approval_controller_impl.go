package controller

import (
	"acl-casbin/dto"
	"acl-casbin/model"
	"acl-casbin/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type approvalController struct {
	service service.ApprovalService
}

func NewApprovalController(s service.ApprovalService) ApprovalController {
	return &approvalController{
		service: s,
	}
}

// ApplyDecision godoc
// @Summary Apply approval decision
// @Description Apply a decision to a step in an approval flow
// @Tags Approval
// @Accept json
// @Produce json
// @Param request body dto.ApplyDecisionRequest true "Apply Decision Request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /approval/apply-decision [post]
func (ctrl *approvalController) ApplyDecision(c *gin.Context) {
	var req dto.ApplyDecisionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json", "details": err.Error()})
		return
	}
	decision := model.ApprovalDecision{
		By:      req.Decision.By,
		Action:  req.Decision.Action,
		Comment: req.Decision.Comment,
		At:      req.Decision.At,
	}

	if err := ctrl.service.ApplyDecision(context.Background(), req.ObjectID, req.ObjectType, req.StepID, decision); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not apply decision", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "decision applied successfully"})
}

// ListFlows godoc
// @Summary List approval flows
// @Description List all approval flows filtered by object type and status
// @Tags Approval
// @Accept json
// @Produce json
// @Param object_type query string false "Object Type"
// @Param status query int false "Flow Status"
// @Success 200 {array} model.ApprovalFlow
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /approval/list [get]
func (ctrl *approvalController) ListFlows(c *gin.Context) {
	objectType := c.Query("object_type")
	statusStr := c.Query("status")
	var status *int
	if statusStr != "" {
		parsed, err := strconv.Atoi(statusStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
			return
		}
		status = &parsed
	}
	flows, err := ctrl.service.ListFlows(context.Background(), objectType, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not list flows", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, flows)
}

// GetFlowByObject godoc
// @Summary Get approval flow by object ID and type
// @Description Retrieve a flow using object type and object ID
// @Tags Approval
// @Accept json
// @Produce json
// @Param object_type path string true "Object Type"
// @Param object_id path string true "Object ID"
// @Success 200 {object} model.ApprovalFlow
// @Failure 404 {object} map[string]string
// @Router /approval/flow/{object_type}/{object_id} [get]
func (ctrl *approvalController) GetFlowByObject(c *gin.Context) {
	objectID := c.Param("object_id")
	objectType := c.Param("object_type")

	flow, err := ctrl.service.GetFlowByObject(context.Background(), objectID, objectType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "flow not found", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, flow)
}

// CreateFlow godoc
// @Summary Create a new approval flow
// @Description Create an approval flow for a specific object
// @Tags Approval
// @Accept json
// @Produce json
// @Param request body dto.CreateFlowRequest true "Create Flow Request"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /approval/create-flow [post]
func (ctrl *approvalController) CreateFlow(c *gin.Context) {
	var req dto.CreateFlowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json", "details": err.Error()})
		return
	}
	flow := model.ApprovalFlow{
		ObjectID:    req.ObjectID,
		ObjectType:  req.ObjectType,
		Steps:       req.ToModelSteps(),
		Status:      0,
		FinalStepID: req.FinalStepID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := ctrl.service.CreateFlow(context.Background(), flow); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create flow", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "flow created successfully"})
}

// UpdateStepStatus godoc
// @Summary Update step status
// @Description Update the status of a step in the approval flow
// @Tags Approval
// @Accept json
// @Produce json
// @Param request body dto.UpdateStepStatusRequest true "Update Step Status Request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /approval/update-step [post]
func (ctrl *approvalController) UpdateStepStatus(c *gin.Context) {
	var req dto.UpdateStepStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json", "details": err.Error()})
		return
	}
	err := ctrl.service.UpdateStepStatus(
		context.Background(),
		req.ObjectID,
		req.ObjectType,
		req.StepID,
		req.Status,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update step", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "step status updated"})
}
