package main

import "net/http"

// JSONMiddleware is the middleware for setting the content-type of a response to JSON.
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			next.ServeHTTP(w, r)
		},
	)
}

