package main

import (
	"GoChat/internal/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()
	log.Println("서버 시작: http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
