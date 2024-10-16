package main

import (
	"context"
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
	EP_REDIS := modulesEnv.EnvLoader(string(staticEnv.EnvListUsedByServer.EP_REDIS), GIN_MODE)

	REDIS_HOST := modulesEnv.EnvLoader(string(staticEnv.REDIS_HOST), GIN_MODE)
	REDIS_PORT := modulesEnv.EnvLoader(string(staticEnv.REDIS_PORT), GIN_MODE)
	REDIS_PW := modulesEnv.EnvLoader(string(staticEnv.REDIS_PW), GIN_MODE)

	gin.SetMode(GIN_MODE)
	r := gin.New()
	r.SetTrustedProxies([]string{WL_PROXIES})
	r.Use(cors.New(modulesCors.BasicCorsConfig()))

	// Redis 클라이언트 생성
	rdb := redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST +  REDIS_PORT,
		Password: REDIS_PW,
		DB:       0,
	})

	// Redis 데이터를 반환하는 엔드포인트 추가
	modulesHttpMethod.GinMethodHandler(r, http.MethodPost, staticSymbols.ForwardSlash, func(c *gin.Context) {
		data, err := rdb.ZRangeWithScores(ctx, "popular_searches", 0, -1).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis에서 데이터를 가져오는 데 실패했습니다."})
			return
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(EP_REDIS)
}
