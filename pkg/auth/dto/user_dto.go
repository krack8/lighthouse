package dto

import "github.com/krack8/lighthouse/pkg/auth/enum"

type UserDTO struct {
	Username      string      `json:"username" binding:"required,email"`
	FirstName     string      `json:"first_name"`
	LastName      string      `json:"last_name"`
	Password      string      `json:"password" binding:"required,min=6,max=15"`
	UserType      string      `json:"user_type" binding:"required,oneof=ADMIN USER"`
	RoleIds       []string    `json:"role_ids" binding:"required"` // Array of role IDs
	ClusterIdList []string    `json:"cluster_ids"`
	UserIsActive  bool        `json:"user_is_active" binding:"required"`
	IsVerified    bool        `json:"is_verified" binding:"required"`
	Phone         string      `json:"phone,omitempty"`
	Status        enum.Status `json:"status"`
	CreatedBy     string      `json:"created_by"`
	UpdatedBy     string      `json:"updated_by"`
}
