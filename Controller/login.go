package Controller

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("Templates/index.html")
	t.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		t, _ := template.ParseFiles("Templates/index.html")
		t.Execute(w, nil)
		return
	}

	inputID := r.FormValue("id")
	inputPassword := r.FormValue("password")

	if inputID == "Yoon" && inputPassword == "P@ssw0rd" {
		cookie := &http.Cookie{
			Name:  "username",
			Value: inputID,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/greeting", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
