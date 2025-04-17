package routes

import (
	"Gochat/main"
	"html/template"
	"net/http"
)

// GetName은 사용자 이름을 반환하는 핸들러입니다.
func GetName(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/greeting.html")
	t.Execute(w, map[string]string{"Name": "Yoon"})
}

// GetNumber는 사용자 번호를 반환하는 핸들러입니다.
func GetNumber(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/greeting.html")
	t.Execute(w, map[string]string{"Number": "12345"})
}

func RegisterGreetingRoutes(router http.Handler) {
	if mux, ok := router.(*main.CustomRouter); ok {
		mux.Handle("/greet", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})
	}
}
