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
    MAIN_URL := modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.MAIN_URL))

    r := gin.New()

    gin.SetMode(GIN_MODE)

    r.SetTrustedProxies([]string {WL_PROXIES})

    r.Use(cors.New(modulesCors.BasicCorsConfig()))

    modulesHttpMethod.GinMethodHandler(r, http.MethodGet, staticSymbols.ForwardSlash, func(c *gin.Context) {
		response := gin.H{"message": "Hello, maingd!"}
	    c.JSON(http.StatusOK, response)
	})

    r.Run(MAIN_URL)
}
