package main

import (
	"log"

	"scada-api/internal/routes"
	"github.com/gin-gonic/gin"
)


func main() {
        router := gin.New()
        router.Use(gin.Logger())
        router.Use(gin.Recovery())
        routes.Setup(router)
	addr := "0.0.0.0:8080"
        log.Printf("Server starting on %s", addr)
        if err := router.Run(addr); err != nil {
                log.Fatal("Failed to start server:", err)
        }

}

