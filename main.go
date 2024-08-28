// package main

// import (
// 	"net/http"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
//     router := gin.Default()

//     router.Use(cors.New(cors.Config{
//         AllowOrigins:     []string{"*"},
//         AllowMethods:     []string{"GET", "POST", "OPTIONS"},
//         AllowHeaders:     []string{"Content-Type"},
//         AllowCredentials: true,
//     }))

//     router.GET("/", func(c *gin.Context) {
//         response := gin.H{"message": "Hello, world!"}
//         c.JSON(http.StatusOK, response)
//     })

//     router.Run(":8080")
// }

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type"},
        AllowCredentials: true,
    }))

    router.GET("/", func(c *gin.Context) {
        response := gin.H{"message": "Hello, World!"}
        c.JSON(http.StatusOK, response)
    })

    // 다른 API 서버들 실행
    servers := []struct {
        name string
        path string
    }{
        {"Activity", "servers/activity-api-server/activity-main.go"},
        {"Channel", "servers/channel-api-server/channel-main.go"},
        {"Comment", "servers/comment-api-server/comment-main.go"},
        {"GA", "servers/ga-api-server/ga-main.go"},
        {"Playlist", "servers/playlist-api-server/playlist-main.go"},
        {"S3", "servers/s3-api-server/s3-main.go"},
        {"Thumbnail", "servers/thumbnail-api-server/thumbnail-main.go"},
        {"User", "servers/user-api-server/user-main.go"},
        {"Video", "servers/video-api-server/video-main.go"},
    }

    for _, server := range servers {
        go func(name, path string) {
            cmd := exec.Command("go", "run", path)
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr

            log.Printf("Starting %s server...", name)
            if err := cmd.Start(); err != nil {
                log.Printf("Error starting %s server: %v", name, err)
                return
            }

            if err := cmd.Wait(); err != nil {
                log.Printf("%s server exited with error: %v", name, err)
            }
        }(server.name, server.path)
    }

    fmt.Println("All servers are starting...")

    // 종료 시그널 처리
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        <-quit
        fmt.Println("Shutting down servers...")
        // 여기에 정상적인 종료 로직 추가 가능
        os.Exit(0)
    }()

    // 메인 서버 실행
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Main server error: %v", err)
    }
}