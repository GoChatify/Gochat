package Config

import "net/http"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("username")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusUnauthorized)
			return
		}
		r.Header.Set("username", cookie.Value)
		next(w, r)
	}
}
