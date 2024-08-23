package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", helloHandler) // 루트 경로에 핸들러 등록
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
