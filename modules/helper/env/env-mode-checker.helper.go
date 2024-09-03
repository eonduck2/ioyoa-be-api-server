package env

import (
	staticEnv "ioyoa/static/env"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// 개발 환경과 배포 환경을 구분
func EnvModeChecker() string {

    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }

	return os.Getenv(string(staticEnv.EnvListUsedByServer.GIN_MODE))
}