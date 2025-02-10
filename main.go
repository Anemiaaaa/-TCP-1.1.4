package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	routersGroup := map[string] string {
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

	for path, response := range routersGroup{
		r.Get(path, func(responseText string) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(responseText))
			}
		}(response))
	}

	http.ListenAndServe(":8080", r)
}