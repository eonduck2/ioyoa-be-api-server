package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
    servers := []struct {
        name string
        path string
    }{
        {"Main", "servers/main.go"},
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
    fmt.Println("Press Ctrl+C to stop all servers.")

    select {}
}