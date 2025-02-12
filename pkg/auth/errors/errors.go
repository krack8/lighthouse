package errors

import "errors"

var (
	ErrTokenExpired     = errors.New("token has expired")
	ErrTokenRevoked     = errors.New("token has been revoked")
	ErrTokenNotFound    = errors.New("token not found")
	ErrInvalidSignature = errors.New("invalid token signature")
	ErrTokenCollision   = errors.New("token collision detected")
)
