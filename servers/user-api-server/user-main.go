package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type"},
        AllowCredentials: true,
    }))

    router.GET("/", func(c *gin.Context) {
        response := gin.H{"message": "Hello, user!"}
        c.JSON(http.StatusOK, response)
    })

    router.Run(":8086")
}
