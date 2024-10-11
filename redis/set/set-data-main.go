package main

import (
	"context"
	"fmt"
	helperEnv "ioyoa/modules/helper/env"
	utilsEnv "ioyoa/modules/utils/env"
	staticEnv "ioyoa/static/env"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {

    GIN_MODE := helperEnv.EnvModeChecker()
    RDB_HOST := utilsEnv.EnvLoader(staticEnv.REDIS_HOST, GIN_MODE)
    RDB_PORT := utilsEnv.EnvLoader(staticEnv.REDIS_PORT, GIN_MODE)
    RDB_PW := utilsEnv.EnvLoader(staticEnv.REDIS_PW, GIN_MODE)
    rdb := redis.NewClient(&redis.Options{
        Addr:        RDB_HOST + RDB_PORT,
        Password:    RDB_PW,
        DB:          0,    
    })
    

    key := "user:1001"
    value := "Alice, ddd25, New Yoqwdeasdsadsark" 

    err := rdb.Set(ctx, key, value, 0).Err() 
    if err != nil {
        log.Fatalf("데이터 저장 실패: %v", err)
    }
    fmt.Printf("데이터 저장 완료: %s = %s\n", key, value)

    // 데이터 불러오기
    retrievedValue, err := rdb.Get(ctx, key).Result()
    if err == redis.Nil {
        
        fmt.Printf("키 %s에 해당하는 데이터가 없습니다.\n", key)
    } else if err != nil {
        log.Fatalf("데이터 불러오기 실패: %v", err)
    } else {
        fmt.Printf("불러온 데이터: %s = %s\n", key, retrievedValue)
    }

    // 연결 해제
    err = rdb.Close()
    if err != nil {
        log.Fatalf("Redis 연결 해제 실패: %v", err)
    }
    fmt.Println("Redis 연결 해제 성공")
}
