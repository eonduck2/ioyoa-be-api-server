package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis 서버 주소
        Password: "",            // 비밀번호가 없으면 비워둠
        DB: 0,                   // 기본 DB 사용
    })

    // Redis 연결 테스트
    pong, err := rdb.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Redis 연결 실패: %v", err)
    }
    fmt.Println("Redis 연결 성공:", pong)
}
