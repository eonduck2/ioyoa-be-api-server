package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }

    ginMode := os.Getenv("GIN_MODE")

    // 환경 모드 설정
    gin.SetMode(ginMode)

    router := gin.New()

    router.SetTrustedProxies([]string {"192.168.100.134"})

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type"},
        AllowCredentials: true,
    }))

    router.GET("/", func(c *gin.Context) {
        response := gin.H{"message": "Hello, maingd!"}
        c.JSON(http.StatusOK, response)
    })

    router.Run(":8080")
}
