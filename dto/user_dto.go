package dto

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"full_name"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" binding:"omitempty,email"`
	FullName string `json:"full_name"`
}

type UserResponse struct {
	UID       string `json:"uid"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	CreatedAt string `json:"created_at"`
}
