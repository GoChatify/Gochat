package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// 라우터 설정
	router := SetupRouter()

	// 서버 시작
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

func login(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		t, _ := template.ParseFiles("index.html")
		t.Execute(res, nil)
		return
	}

	inputID := req.FormValue("id")
	inputPassword := req.FormValue("password")

	if inputID == "Yoon" && inputPassword == "P@ssw0rd" {
		// 세션 ID 생성
		sessionID := GenerateSessionID()

		// 세션 저장
		Mutex.Lock()
		SessionStore[sessionID] = inputID
		Mutex.Unlock()

		// 쿠키에 세션 ID 저장
		http.SetCookie(res, &http.Cookie{
			Name:  "session_id",
			Value: sessionID,
			Path:  "/",
		})

		http.Redirect(res, req, "/greeting", http.StatusSeeOther)
	} else {
		http.Error(res, "Invalid credentials", http.StatusUnauthorized)
	}
}

func greeting(res http.ResponseWriter, req *http.Request) {
	// 쿠키에서 세션 ID 가져오기
	cookie, _ := req.Cookie("session_id")
	sessionID := cookie.Value

	// 세션 ID로 사용자 이름 확인
	Mutex.Lock()
	username := SessionStore[sessionID]
	Mutex.Unlock()

	t, _ := template.ParseFiles("greeting.html")
	t.Execute(res, map[string]string{"Name": username})
}
