package main

import (
	"GoChat/config"
	"fmt"
	"net/http"
)

func main() {
	config.RegisterRoutes()
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// // loginHandler는 로그인 폼을 처리하는 핸들러입니다.
// func login(res http.ResponseWriter, req *http.Request) {
// 	if req.Method != http.MethodPost {
// 		t, _ := template.ParseFiles("index.html")
// 		t.Execute(res, nil)
// 		return
// 	}

// 	inputID := req.FormValue("id")
// 	inputPassword := req.FormValue("password")

// 	if inputID == "Yoon" && inputPassword == "P@ssw0rd" {
// 		cookie := &http.Cookie{
// 			Name:  "username",
// 			Value: inputID,
// 			Path:  "/",
// 		}
// 		http.SetCookie(res, cookie)

// 		http.Redirect(res, req, "/greeting", http.StatusSeeOther)
// 	} else {
// 		http.Error(res, "Invalid credentials", http.StatusUnauthorized)
// 	}
// }

// func greeting(res http.ResponseWriter, req *http.Request) {
// 	if req.Method != http.MethodPost {
// 		cookie, err := req.Cookie("username")
// 		var name string
// 		if err != nil {
// 			http.Redirect(res, req, "/", http.StatusNetworkAuthenticationRequired)
// 			return
// 		} else {
// 			name = cookie.Value
// 		}

// 		t, _ := template.ParseFiles("greeting.html")
// 		t.Execute(res, map[string]string{"Name": name})
// 	}
// }
