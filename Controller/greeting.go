package Controller

import (
	"html/template"
	"net/http"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("username")
	t, _ := template.ParseFiles("templates/greeting.html")
	t.Execute(w, map[string]string{"Name": name})
}
