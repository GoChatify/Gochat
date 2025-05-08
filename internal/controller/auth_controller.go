package controller

import (
	"GoChat/internal/model"
	"GoChat/internal/service"
	"encoding/json"
	"net/http"
)

type AuthController struct {
	AuthService *service.AuthService
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	token, err := ac.AuthService.GenerateJWT(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
