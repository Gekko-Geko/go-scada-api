package routes

import (
        "scada-api/internal/handlers"

        "github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
        router.GET("/health", handlers.Health)
        v1 := router.Group("/api/v1") 
	{
                motors := v1.Group("/motors") 
		{
                        motors.GET("", handlers.GetMotors)
                }
                pumps := v1.Group("/pumps") 
		{
                        pumps.GET("", handlers.GetPumps)
                }
                tanks := v1.Group("/tanks") 
		{
                        tanks.GET("", handlers.GetTanks)
                }
        }
}

