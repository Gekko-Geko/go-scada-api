package handlers

import (
        "net/http"
        "scada-api/internal/mocks"

        "github.com/gin-gonic/gin"
)

func GetMotors(c *gin.Context) {
	motors := mocks.GetMockMotors()
	c.JSON(http.StatusOK, gin.H{
		"motors": motors,
	})
}

