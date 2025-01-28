package utils

import (
	"github.com/gin-gonic/gin"
)

// RespondWithJSON is a helper function to send JSON responses
func RespondWithJSON(c *gin.Context, statusCode int, payload interface{}) {
	c.JSON(statusCode, payload)
}

// RespondWithError is a helper function to send error responses
func RespondWithError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}
