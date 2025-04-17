package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	formTemplate := getTemplates("form.html")
	greetingTemplate := getTemplates("greeting.html")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			greetingTemplate.Execute(w, map[string]string{"Name": name})
		} else {
			formTemplate.Execute(w, nil)
		}
	})

	http.HandleFunc("/api/getType", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		response := map[string]string{"type": "greeting"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started at http://localhost:8080")
}

func getTemplates(fileName string) *template.Template {
	if fileName == "" {
		panic("File name cannot be empty")
	}
	if fileName == "form.html" {
		return template.Must(template.ParseFiles("form.html"))
	}
	if fileName == "greeting.html" {
		return template.Must(template.ParseFiles("greeting.html"))
	}
	panic("Invalid file name")
}
