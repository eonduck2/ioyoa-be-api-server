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
    EP_ACTIVITY := modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.EP_ACTIVITY), GIN_MODE)

    gin.SetMode(GIN_MODE)

    r := gin.New()

    r.SetTrustedProxies([]string{WL_PROXIES})

    r.Use(cors.New(modulesCors.BasicCorsConfig()))

    modulesHttpMethod.GinMethodHandler(r, http.MethodPost, staticSymbols.ForwardSlash, func(c *gin.Context) {
        var requestBody map[string]string
        if err := c.BindJSON(&requestBody); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청 본문입니다."})
            return
        }

        // 요청 본문에서 URL을 가져옵니다.
        url, exists := requestBody["url"]
        if !exists {
            c.JSON(http.StatusBadRequest, gin.H{"error": "URL이 제공되지 않았습니다."})
            return
        }

        // 해당 URL로 GET 요청을 보냅니다.
        resp, err := http.Get(url)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "검색 데이터를 가져오는 데 실패했습니다."})
            return
        }
        defer resp.Body.Close()

        // 응답 본문을 읽습니다.
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "응답 본문을 읽는 데 실패했습니다."})
            return
        }

        // 읽은 데이터를 클라이언트에 반환합니다.
        c.Data(http.StatusOK, "application/json", body)
    })

    r.Run(EP_ACTIVITY)
}
