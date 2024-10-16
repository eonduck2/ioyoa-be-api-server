package main

import (
	"context"
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
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	GIN_MODE := helperEnv.EnvModeChecker()
	WL_PROXIES := modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.WL_PROXIES), GIN_MODE)
	EP_SEARCH := modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.EP_SEARCH), GIN_MODE)

	REDIS_HOST := modulesEnv.EnvLoader(string(staticEnv.REDIS_HOST), GIN_MODE)
	REDIS_PORT := modulesEnv.EnvLoader(string(staticEnv.REDIS_PORT), GIN_MODE)
	REDIS_PW := modulesEnv.EnvLoader(string(staticEnv.REDIS_PW), GIN_MODE)

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

		searchValue, searchExists := requestBody["search_value"]
		url, urlExists := requestBody["url"]

		if searchExists {
			// Redis 클라이언트 생성
			rdb := redis.NewClient(&redis.Options{
				Addr:     REDIS_HOST +  REDIS_PORT, // 포트는 ':'로 구분
				Password: REDIS_PW, // 비밀번호 (설정되어 있으면 입력)
				DB:       0,       // 사용할 DB 번호
			})

			// 검색어를 Sorted Set에 추가 (빈도수 증가)
			_, err := rdb.ZIncrBy(ctx, "popular_searches", 1, searchValue).Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis에 저장하는 데 실패했습니다."})
				return
			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "searchValue가 제공되지 않았습니다."})
			return
		}

		if !urlExists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "URL이 제공되지 않았습니다."})
			return
		}

		resp, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "검색 데이터를 가져오는 데 실패했습니다."})
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "응답 본문을 읽는 데 실패했습니다."})
			return
		}

		// JSON 형식으로 응답
		c.Data(http.StatusOK, "application/json", body)
	})

	r.Run(EP_SEARCH)
}
