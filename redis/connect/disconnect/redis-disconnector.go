package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)


func main() {

	rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis 서버 주소
        Password: "",            // 비밀번호가 없으면 비워둠
        DB: 0,                   // 기본 DB 사용
    })

    // 연결 해제
    err := rdb.Close()
    if err != nil {
        log.Fatalf("Redis 연결 해제 실패: %v", err)
    }
    fmt.Println("Redis 연결 해제 성공")
}
