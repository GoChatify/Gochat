package main

import (
	"Gochat/routes"
	"html/template"
	"net/http"
)

// CustomRouter는 URL 경로와 핸들러를 매핑하는 간단한 라우터입니다.
type CustomRouter struct {
	routes map[string]http.HandlerFunc
}

// NewRouter는 새로운 라우터를 생성합니다.
func NewRouter() *CustomRouter {
	return &CustomRouter{routes: make(map[string]http.HandlerFunc)}
}

// Handle은 경로와 핸들러를 등록합니다.
func (r *CustomRouter) Handle(path string, handler http.HandlerFunc) {
	r.routes[path] = handler
}

// ServeHTTP는 요청 경로에 따라 적절한 핸들러를 호출합니다.
func (r *CustomRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, exists := r.routes[req.URL.Path]; exists {
		handler(w, req)
	} else {
		http.NotFound(w, req)
	}
}

// SetupRouter는 라우터를 설정하고 반환합니다.
func SetupRouter() *CustomRouter {
	router := NewRouter()

	// 기본 경로 등록
	router.Handle("/", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	})

	// 모듈화된 경로 등록
	routes.RegisterGreetingRoutes(router)
	routes.RegisterLoginRoutes(router)

	return router
}
