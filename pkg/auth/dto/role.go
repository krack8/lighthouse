package dto

// RoleDTO with permissions as string slice
type RoleDTO struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions" binding:"required"`
}
