package dto

import "github.com/krack8/lighthouse/pkg/auth/enum"

// RoleDTO with permissions as string slice
type RoleDTO struct {
	Name        string      `json:"name" binding:"required"`
	Description string      `json:"description"`
	Permissions []string    `json:"permissions" binding:"required"` // String slice
	Status      enum.Status `json:"status"`
	CreatedBy   string      `json:"created_by"`
	UpdatedBy   string      `json:"updated_by"`
}
