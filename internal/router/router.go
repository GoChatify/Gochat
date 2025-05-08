package router

import (
	"GoChat/internal/controller"
	"GoChat/internal/service"
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	authService := &service.AuthService{SecretKey: "DBWrjBFv78AzkggUVxKiPngUEfvR0uECpz4qJut+jLM="}
	authController := &controller.AuthController{AuthService: authService}

	mux.HandleFunc("/api/v1/auth/signin", authController.Login)

	return mux
}
