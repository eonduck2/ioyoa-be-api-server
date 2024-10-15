package main

import (
	helperEnv "ioyoa/modules/helper/env"
	modulesCors "ioyoa/modules/helper/middleWare/cors"
	modulesHttpMethod "ioyoa/modules/server/gin/httpMethod"
	modulesEnv "ioyoa/modules/utils/env"
	staticEnv "ioyoa/static/env"
	staticSymbols "ioyoa/static/shared/symbols"
	staticUrl "ioyoa/static/shared/url"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RequestBody 정의
type RequestBody struct {
	ServerType string `json:"serverType"`
}

func main() {
	GIN_MODE := helperEnv.EnvModeChecker()
	WL_PROXIES := modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.WL_PROXIES), GIN_MODE)
	EP_MAIN := modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.EP_MAIN), GIN_MODE)

	gin.SetMode(GIN_MODE)

	r := gin.New()

	r.SetTrustedProxies([]string{WL_PROXIES})

	r.Use(cors.New(modulesCors.BasicCorsConfig()))

	modulesHttpMethod.GinMethodHandler(r, http.MethodPost, staticSymbols.ForwardSlash, func(c *gin.Context) {
		var requestBody RequestBody

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		switch requestBody.ServerType {
		case "video":
			response := gin.H{staticUrl.Route: modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.EP_VIDEO), GIN_MODE)}
			c.JSON(http.StatusOK, response)
		case "search":
			response := gin.H{staticUrl.Route: modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.EP_SEARCH), GIN_MODE)}
			c.JSON(http.StatusOK, response)
		case "channel":
			response := gin.H{staticUrl.Route: modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.EP_CHANNEL), GIN_MODE)}
			c.JSON(http.StatusOK, response)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid serverType"})
		}
	})

	r.Run(EP_MAIN)
}
