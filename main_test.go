package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	r := chi.NewRouter()
	groupRoutes := map[string]string{
		"/group1/1": "Group 1 Привет, мир 1",
		"/group1/2": "Group 1 Привет, мир 2",
		"/group1/3": "Group 1 Привет, мир 3",
		"/group2/1": "Group 2 Привет, мир 1",
		"/group2/2": "Group 2 Привет, мир 2",
		"/group2/3": "Group 2 Привет, мир 3",
		"/group3/1": "Group 3 Привет, мир 1",
		"/group3/2": "Group 3 Привет, мир 2",
		"/group3/3": "Group 3 Привет, мир 3",
	}

	for path, response := range groupRoutes {
		r.Get(path, func(responseText string) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(responseText))
			}
		}(response))
	}

	tests := []struct {
		name         string
		path         string
		expectedCode int
		expectedBody string
	}{
		{"Route /group1/1", "/group1/1", http.StatusOK, "Group 1 Привет, мир 1"},
		{"Route /group1/2", "/group1/2", http.StatusOK, "Group 1 Привет, мир 2"},
		{"Route /group1/3", "/group1/3", http.StatusOK, "Group 1 Привет, мир 3"},
		{"Route /group2/1", "/group2/1", http.StatusOK, "Group 2 Привет, мир 1"},
		{"Route /group2/2", "/group2/2", http.StatusOK, "Group 2 Привет, мир 2"},
		{"Route /group2/3", "/group2/3", http.StatusOK, "Group 2 Привет, мир 3"},
		{"Route /group3/1", "/group3/1", http.StatusOK, "Group 3 Привет, мир 1"},
		{"Route /group3/2", "/group3/2", http.StatusOK, "Group 3 Привет, мир 2"},
		{"Route /group3/3", "/group3/3", http.StatusOK, "Group 3 Привет, мир 3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", tt.path, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("expected %d, got %d", tt.expectedCode, w.Code)
			}

			if w.Body.String() != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, w.Body.String())
			}
		})
	}
}
