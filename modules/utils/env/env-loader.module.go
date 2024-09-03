package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// EnvLoader는 환경 변수의 이름을 사용하여 해당 환경 변수의 값을 반환합니다.
// envName: 검색할 환경 변수의 이름
// envType: 환경 타입 (예: "debug" 또는 "release")
// 반환값: 환경 변수의 값 (없을 경우 빈 문자열)
func EnvLoader(envName string, envType string) string {
    // 환경에 따라 다른 .env 파일을 로드
    var envFile string
    switch envType {
    case "debug":
        envFile = ".env.development"
    case "release":
        envFile = ".env.production"
    default:
        return ""
    }

    if err := godotenv.Load(envFile); err != nil {
        log.Fatalf("Error loading %s file", envFile)
    }

    return os.Getenv(envName)
}
