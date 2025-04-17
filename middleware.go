package main

import (
	"net/http"
)

// isAuthenticated는 사용자가 인증되었는지 확인합니다.
func isAuthenticated(req *http.Request) bool {
	cookie, err := req.Cookie("session_id")
	if err != nil {
		return false
	}

	sessionID := cookie.Value

	Mutex.Lock()
	_, exists := SessionStore[sessionID]
	Mutex.Unlock()

	return exists
}
