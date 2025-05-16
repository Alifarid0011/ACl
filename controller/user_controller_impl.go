package controller

import (
	"acl-casbin/dto"
	"acl-casbin/dto/response"
	"acl-casbin/service"
	"acl-casbin/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type userControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userControllerImpl{userService: userService}
}

// FindByUsername
// @Summary      Find user by username
// @Description  Get a user object by username
// @Tags         users
// @Security AuthBearer
// @Param        username  path  string  true  "Username"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Success 200 {object} dto.MessageResponse
// @Router       /users/username/{username} [get]
func (u userControllerImpl) FindByUsername(ctx *gin.Context) {
	username := ctx.Param("username")

	userResponse, err := u.userService.FindByUsername(username)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to find user"})
		return
	}
	ctx.JSON(200, userResponse)
}

// FindByUID
// @Summary      Find user by UID
// @Description  Get a user by their UID
// @Tags         users
// @Security AuthBearer
// @Param        uid  path  string  true  "User UID"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Success 200 {object} dto.MessageResponse
// @Router       /users/uid/{uid} [get]
func (u userControllerImpl) FindByUID(ctx *gin.Context) {
	uid := ctx.Param("uid")
	objectID, err := primitive.ObjectIDFromHex(uid)
	userResponse, err := u.userService.FindByUID(objectID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to find user"})
		return
	}
	ctx.JSON(200, userResponse)
}

// Create
// @Summary      Create a new user
// @Description  Register a new user in the system
// @Tags         users
// @Security AuthBearer
// @Accept       json
// @Produce      json
// @Param        data  body  dto.CreateUserRequest  true  "User data"
// @Success      201  {object}  dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Success 200 {object} dto.MessageResponse
// @Router       /users/create [post]
func (u userControllerImpl) Create(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errors := utils.GetValidationErrors(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errors,
		})
		return
	}
	userResponse, err := u.userService.CreateUser(req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.JSON(201, userResponse)
}

// GetAll
// @Summary      Get all users
// @Description  Retrieve a list of all users
// @Tags         users
// @Security AuthBearer
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Success 200 {object} dto.MessageResponse
// @Router       /users/all [get]
func (u userControllerImpl) GetAll(ctx *gin.Context) {
	users, err := u.userService.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}
	ctx.JSON(200, users)
}

// Update
// @Summary      Update a user
// @Description  Update user details by UID (admin or self)
// @Tags         users
// @Security AuthBearer
// @Accept       json
// @Produce      json
// @Param        uid   path  string  true  "User UID"
// @Param        data  body  dto.UpdateUserRequest  true  "User update data"
// @Success      200  {object}  dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router       /users/{uid} [put]
func (u userControllerImpl) Update(ctx *gin.Context) {
	// Check if user is a super admin or update their own data
	uid := ctx.Param("uid") // uid from URL parameter
	// In a production setting, this would be verified from JWT token
	//role := ctx.GetString("role") // assume role is in context after JWT middleware
	var req dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	//if role != "superadmin" && uid != ctx.GetString("userID") {
	//	ctx.JSON(403, gin.H{"error": "You do not have permission to update this user"})
	//	return
	//}
	// Proceed with update
	objectID, err := primitive.ObjectIDFromHex(uid)
	userResponse, err := u.userService.UpdateUser(objectID, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}
	ctx.JSON(200, userResponse)
}

// Delete
// @Summary      Delete a user
// @Description  Delete user by UID (admin only)
// @Tags         users
// @Security AuthBearer
// @Param        uid  path  string  true  "User UID"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Success 200 {object} dto.MessageResponse
// @Router       /users/{uid} [delete]
func (u userControllerImpl) Delete(ctx *gin.Context) {
	uid := ctx.Param("uid")
	objectID, _ := primitive.ObjectIDFromHex(uid)
	errUserService := u.userService.DeleteUser(objectID)
	if errUserService != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}
	ctx.JSON(200, gin.H{"message": "User deleted successfully"})
}

// Me
// @Summary      Get current user info
// @Description  Get authenticated user information from token
// @Tags         users
// @Security AuthBearer
// @Success      200  {object}  dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router       /users/me [get]
func (u userControllerImpl) Me(ctx *gin.Context) {
	// Retrieve current user ID from the JWT or session context
	userID := ctx.GetString("user_uid")
	objectID, err := primitive.ObjectIDFromHex(userID)
	userResponse, err := u.userService.Me(objectID)
	if err != nil {
		response.New(ctx).Errors(err).MessageID("users.me.get.failed").Status(http.StatusInternalServerError).Dispatch()
		return
	}
	response.New(ctx).Message("اطلاعات به درستی دریافت شد.").MessageID("users.me.get.success").Data(userResponse).Status(http.StatusOK).Dispatch()
}
