package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		id             string
		password       string
		expectedStatus int
		expectedHeader string
	}{
		{
			name:           "Valid credentials",
			method:         http.MethodPost,
			id:             "Yoon",
			password:       "P@ssw0rd",
			expectedStatus: http.StatusSeeOther,
			expectedHeader: "/greeting",
		},
		{
			name:           "Invalid credentials - wrong id and password",
			method:         http.MethodPost,
			id:             "wrong",
			password:       "wrong",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid credentials - wrong id",
			method:         http.MethodPost,
			id:             "wrong",
			password:       "P@ssw0rd",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid credentials - wrong password",
			method:         http.MethodPost,
			id:             "Yoon",
			password:       "wrong",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid method",
			method:         http.MethodGet,
			id:             "",
			password:       "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formData := strings.NewReader("id=" + tt.id + "&password=" + tt.password)
			req := httptest.NewRequest(tt.method, "/login", formData)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			res := httptest.NewRecorder()

			login(res, req)

			if res.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, res.Code)
			}

			if tt.expectedStatus == http.StatusSeeOther {
				location := res.Header().Get("Location")
				if location != tt.expectedHeader {
					t.Errorf("expected redirect to %s, got %s", tt.expectedHeader, location)
				}
			}
		})
	}
}
