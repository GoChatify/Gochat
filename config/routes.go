package config

import (
	"GoChat/Controller"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", Controller.Home)
	http.HandleFunc("/login", Controller.Login)
	http.HandleFunc("/greeting", Controller.Greeting)
}
