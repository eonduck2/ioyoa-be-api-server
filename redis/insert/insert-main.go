package insert

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    // Redis 클라이언트 초기화
    rdb := redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379", // Redis 서버 주소
        Password: "",                // 비밀번호 (없으면 비워둠)
        DB:       0,                 // 기본 DB 사용
    })

    // 임의의 데이터 저장
    key := "user:1001"
    value := "Alice, 25, New York" // 예: 사용자 정보

    err := rdb.Set(ctx, key, value, 0).Err() // 0은 만료 시간 없음
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
