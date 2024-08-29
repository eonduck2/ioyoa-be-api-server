package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	serverList "ioyoa/static/server" // 서버 목록을 가져오기 위한 패키지 import
	"ioyoa/types/shared"             // shared.Name과 shared.Path를 가져오기 위한 패키지 import
)

func main() {
	servers := serverList.ServerList // serverList.ServerList를 사용

	for _, srv := range servers {
		go func(name shared.Name, path shared.Path) { // 매개변수 타입을 shared.Name과 shared.Path로 설정
			cmd := exec.Command("go", "run", string(path)) // shared.Path를 string으로 변환
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
		}(srv.Name, srv.Path) // shared.Name과 shared.Path를 사용
	}

	fmt.Println("All servers are starting...")
	fmt.Println("Press Ctrl+C to stop all servers.")

	select {} // 무한 대기
}
