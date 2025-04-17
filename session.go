package main

import (
	"math/rand"
	"sync"
	"time"
)

// SessionStore는 세션 ID와 사용자 이름을 저장하는 맵입니다.
var SessionStore = make(map[string]string)

// Mutex는 세션 접근을 동기화하기 위한 뮤텍스입니다.
var Mutex = &sync.Mutex{}

// GenerateSessionID는 고유한 세션 ID를 생성합니다.
func GenerateSessionID() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	sessionID := make([]byte, 32)
	for i := range sessionID {
		sessionID[i] = charset[rand.Intn(len(charset))]
	}
	return string(sessionID)
}
