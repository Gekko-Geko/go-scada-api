package handlers

import (
        "net/http"
        "scada-api/internal/mocks"

        "github.com/gin-gonic/gin"
)

func GetTanks(c *gin.Context) {
	tanks := mocks.GetMockTanks()
	c.JSON(http.StatusOK, gin.H{
		"tanks": tanks,
	})
}

