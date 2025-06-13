package handlers

import (
        "net/http"
        "scada-api/internal/mocks"

        "github.com/gin-gonic/gin"
)

func GetPumps(c *gin.Context) {
	pumps := mocks.GetMockPumps()
	c.JSON(http.StatusOK, gin.H{
		"pumps": pumps,
	})
}

