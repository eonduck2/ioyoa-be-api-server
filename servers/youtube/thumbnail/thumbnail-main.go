package main

import (
	helperEnv "ioyoa/modules/helper/env"
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

    GIN_MODE :=	helperEnv.EnvModeChecker()
    WL_PROXIES := modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.WL_PROXIES), GIN_MODE)
    EP_THUMBNAIL := modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.EP_THUMBNAIL), GIN_MODE)

    gin.SetMode(GIN_MODE)
    
    r := gin.New()

    r.SetTrustedProxies([]string {WL_PROXIES})

    r.Use(cors.New(modulesCors.BasicCorsConfig()))

    modulesHttpMethod.GinMethodHandler(r, http.MethodGet, staticSymbols.ForwardSlash, func(c *gin.Context) {
		response := gin.H{"message": "Hello, thumb!"}
	    c.JSON(http.StatusOK, response)
	})

    r.Run(EP_THUMBNAIL)
}
