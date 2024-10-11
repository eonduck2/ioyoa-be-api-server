package main

import (
	"io"
	"net/http"

	helperEnv "ioyoa/modules/helper/env"
	modulesCors "ioyoa/modules/helper/middleWare/cors"
	modulesHttpMethod "ioyoa/modules/server/gin/httpMethod"
	modulesEnv "ioyoa/modules/utils/env"
	staticEnv "ioyoa/static/env"
	staticSymbols "ioyoa/static/shared/symbols"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    GIN_MODE := helperEnv.EnvModeChecker()
    WL_PROXIES := modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.WL_PROXIES), GIN_MODE)
    EP_VIDEO := modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.EP_VIDEO), GIN_MODE)
 
    gin.SetMode(GIN_MODE)
    
    r := gin.New()

    r.SetTrustedProxies([]string{WL_PROXIES})

    r.Use(cors.New(modulesCors.BasicCorsConfig()))

    modulesHttpMethod.GinMethodHandler(r, http.MethodGet, staticSymbols.ForwardSlash, func(c *gin.Context){
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello World",
        })
    })

    modulesHttpMethod.GinMethodHandler(r, http.MethodPost, staticSymbols.ForwardSlash, func(c *gin.Context) {
        var requestBody map[string]string
        if err := c.BindJSON(&requestBody); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
            return
        }
        
        url, exists := requestBody["url"]
        if !exists {
            c.JSON(http.StatusBadRequest, gin.H{"error": "URL not provided"})
            return
        }

        resp, err := http.Get(url)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from YouTube"})
            return
        }
        defer resp.Body.Close()

        body, err := io.ReadAll(resp.Body)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
            return
        }

        // here
        
        c.Data(http.StatusOK, "application/json", body)
    })

    r.Run(EP_VIDEO)
}
