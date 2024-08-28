package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func startServer(name, path string, wg *sync.WaitGroup) {
    defer wg.Done()
    
    cmd := exec.Command("air", "-c", filepath.Join(path, ".air.toml"))
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Dir = path

    log.Printf("Starting %s server...", name)
    if err := cmd.Start(); err != nil {
        log.Printf("Error starting %s server: %v", name, err)
        return
    }

    if err := cmd.Wait(); err != nil {
        log.Printf("%s server exited with error: %v", name, err)
    }
}

func main() {
    var wg sync.WaitGroup

    servers := []struct {
        name string
        path string
    }{
        {"Activity", "../servers/activity-api-server"},
        {"Channel", "../servers/channel-api-server"},
        {"Comment", "../servers/comment-api-server"},
        {"GA", "../servers/ga-api-server"},
        {"Playlist", "../servers/playlist-api-server"},
        {"S3", "../servers/s3-api-server"},
        {"Thumbnail", "../servers/thumbnail-api-server"},
        {"User", "../servers/user-api-server"},
        {"Video", "../servers/video-api-server"},
    }

    for _, server := range servers {
        wg.Add(1)
        go startServer(server.name, server.path, &wg)
    }

    wg.Wait()
    fmt.Println("All servers have started")
}
