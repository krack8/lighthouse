package dto

type UserDTO struct {
	Username      string   `json:"username" binding:"required,email"`
	FirstName     string   `json:"first_name"`
	LastName      string   `json:"last_name"`
	Password      string   `json:"password" binding:"required,min=6,max=15"`
	UserType      string   `json:"user_type" binding:"required,oneof=ADMIN USER"`
	RoleIds       []string `json:"role_ids"` // Array of role IDs
	ClusterIdList []string `json:"cluster_ids"`
	UserIsActive  bool     `json:"user_is_active" binding:"required"`
	IsVerified    bool     `json:"is_verified"`
	Phone         string   `json:"phone,omitempty"`
}

// ResetPasswordRequest represents the request payload for resetting password
type ResetPasswordRequest struct {
	CurrentPassword string `json:"currentPassword" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required,min=6,max=15"`
}

// ForgotPasswordRequest represents the request payload for forgot password
type ForgotPasswordRequest struct {
	Username string `json:"username" validate:"required,username"`
}
