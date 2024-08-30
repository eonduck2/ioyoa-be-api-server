package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	static "ioyoa/static/server"
	types "ioyoa/types/shared"
)

func main() {
	servers := static.ServerList // serverList.ServerList를 사용

	for _, srv := range servers {
		go func(name types.TName, path types.TPath) { // 매개변수 타입을 shared.Name과 shared.Path로 설정
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
