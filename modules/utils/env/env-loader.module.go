package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// EnvLoader는 주어진 환경 변수 이름을 사용하여 환경 변수의 값을 반환합니다.
// envName: 검색할 환경 변수의 이름
// 반환값: 환경 변수의 값 (없을 경우 빈 문자열)
func EnvLoader(envName string) string {

    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }

	return os.Getenv(envName)
}