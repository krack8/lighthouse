package services

import (
	"context"
	"errors"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"github.com/krack8/lighthouse/pkg/log"
	"os"
	"time"
)

func Login(username string, password string) (string, string, error) {
	var user *models.User
	user, err := GetUserByUsername(username)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", "", errors.New("wrong password")
	}

	// Load expiry durations from environment variables
	accessTokenExpiry, err := parseDurationFromEnv("ACCESS_TOKEN_EXPIRY")
	if err != nil {
		return "", "", err
	}

	refreshTokenExpiry, err := parseDurationFromEnv("REFRESH_TOKEN_EXPIRY")
	if err != nil {
		return "", "", err
	}

	// Generate JWT tokens with the username and permissions
	accessToken, err := utils.GenerateToken(username, os.Getenv("JWT_SECRET"), accessTokenExpiry)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.GenerateToken(username, os.Getenv("JWT_REFRESH_SECRET"), refreshTokenExpiry)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// Helper function to parse durations from environment variables
func parseDurationFromEnv(envKey string) (time.Duration, error) {
	value := os.Getenv(envKey)
	if value == "" {
		return 0, errors.New(envKey + " is not set in environment variables")
	}

	parsed, err := time.ParseDuration(value)
	if err != nil {
		return 0, errors.New("invalid duration format for " + envKey + ": " + value)
	}

	return parsed, nil
}

func IsAgentAuthTokenValid(authToken string) bool {
	tokenValidation, err := mongoUpdate.GetToken(context.Background(), authToken)
	if tokenValidation == nil {
		log.Logger.Warnw("Unauthorized !!..Token not found in DB", "err", err.Error())
		return false
	}

	if err != nil {
		log.Logger.Errorw("Error fetching token from database", "err", err.Error())
		return false
	}

	_, clusterID, err := ValidateToken(authToken, tokenValidation)
	if err != nil {
		log.Logger.Errorw("Token not valid", "err", err.Error())
		return false
	}

	err = UpdateClusterStatusToActive(clusterID)
	if err != nil {
		log.Logger.Errorw("Failed to update cluster status", "err", err.Error())
		return false
	}
	return true
}
