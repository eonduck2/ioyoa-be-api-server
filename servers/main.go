package main

import (
	module "ioyoa/modules/utils/env"
	static "ioyoa/static/env"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	envVars := static.EnvListUsedByServer

	envValues := make(map[string]string)

	for _, envName := range envVars {
		value := module.EnvLoader(envName)
		envValues[envName] = value
	}

    GIN_MODE := module.EnvLoader("GIN_MODE")
    WL_PROXY := module.EnvLoader("WL_PROXIES")

    router := gin.New()

    gin.SetMode(GIN_MODE)

    router.SetTrustedProxies([]string {WL_PROXY})

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
