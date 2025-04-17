package routes

import (
	"main"
	"net/http"
)

// Login은 로그인 처리를 담당하는 핸들러입니다.
func Login(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	t, _ := template.ParseFiles("index.html")
	// 	t.Execute(w, nil)
	// 	return
	// }

	// inputID := r.FormValue("id")
	// inputPassword := r.FormValue("password")

	// if inputID == "Yoon" && inputPassword == "P@ssw0rd" {
	// 	// 세션 ID 생성
	// 	sessionID := GenerateSessionID()

	// 	// 세션 저장
	// 	Mutex.Lock()
	// 	SessionStore[sessionID] = inputID
	// 	Mutex.Unlock()

	// 	// 쿠키에 세션 ID 저장
	// 	http.SetCookie(w, &http.Cookie{
	// 		Name:  "session_id",
	// 		Value: sessionID,
	// 		Path:  "/",
	// 	})

	// 	http.Redirect(w, r, "/greeting", http.StatusSeeOther)
	// } else {
	// 	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	// }
}

// Logout은 로그아웃 처리를 담당하는 핸들러입니다.
func Logout(w http.ResponseWriter, r *http.Request) {
	// // 쿠키에서 세션 ID 가져오기
	// cookie, err := r.Cookie("session_id")
	// if err == nil {
	// 	// 세션 삭제
	// 	Mutex.Lock()
	// 	delete(SessionStore, cookie.Value)
	// 	Mutex.Unlock()

	// 	// 쿠키 삭제
	// 	http.SetCookie(w, &http.Cookie{
	// 		Name:   "session_id",
	// 		Value:  "",
	// 		Path:   "/",
	// 		MaxAge: -1, // 쿠키 만료
	// 	})
	// }

	// // 로그아웃 완료 메시지
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

// RegisterAuthRoutes는 auth 관련 경로를 등록합니다.
func RegisterLoginRoutes(router *main.CustomRouter) {
	router.HandleFunc("/auth/login", Login)
	router.HandleFunc("/auth/logout", Logout)
}
