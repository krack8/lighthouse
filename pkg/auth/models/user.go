package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

// UserType represents different user types.
type UserType string

const (
	AdminUser   UserType = "ADMIN"
	RegularUser UserType = "USER"
)

// User represents user information and implements user validation and account states.
type User struct {
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	Username            string    `json:"username" validate:"required,email"`
	FirstName           string    `json:"first_name"`
	LastName            string    `json:"last_name"`
	Password            string    `json:"password" validate:"required,min=6,max=15"`
	UserType            UserType  `json:"user_type" validate:"required,oneof=ADMIN USER"`
	UserIsActive        bool      `json:"user_is_active" validate:"required"`
	IsVerified          bool      `json:"is_verified" validate:"required"`
	ForgotPasswordToken string    `json:"forgot_password_token,omitempty"`
	Phone               string    `json:"phone,omitempty"`
}

// Validate validates the UserInfo fields using the validator package.
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// IsAccountNonExpired checks if the account is non-expired.
func (u *User) IsAccountNonExpired() bool {
	return u.UserIsActive
}

// IsAccountNonLocked checks if the account is non-locked.
func (u *User) IsAccountNonLocked() bool {
	return u.UserIsActive
}

// IsCredentialsNonExpired checks if the credentials are non-expired.
func (u *User) IsCredentialsNonExpired() bool {
	return u.UserIsActive
}

// IsEnabled checks if the user is enabled.
func (u *User) IsEnabled() bool {
	return u.UserIsActive
}
