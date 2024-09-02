package main

import (
	modulesCors "ioyoa/modules/helper/middleWare/cors"
	modulesHttpMethod "ioyoa/modules/server/gin/httpMethod"
	modulesEnv "ioyoa/modules/utils/env"
	staticEnv "ioyoa/static/env"
	staticSymbols "ioyoa/static/shared/symbols"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

    GIN_MODE := modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.GIN_MODE))
    WL_PROXIES := modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.WL_PROXIES))

    router := gin.New()

    gin.SetMode(GIN_MODE)

    router.SetTrustedProxies([]string {WL_PROXIES})

    router.Use(cors.New(modulesCors.BasicCorsConfig()))

    modulesHttpMethod.GinMethodHandler(router, http.MethodGet, staticSymbols.ForwardSlash, func(c *gin.Context) {
		response := gin.H{"message": "Hello, maingd!"}
		c.JSON(http.StatusOK, response)
	})

    router.Run(":8080")
}
