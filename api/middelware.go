package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"log"
	"net/http"
	"strings"
	"way/pkg/response"
)

// JSONMiddleware is the middleware for setting the content-type of a response to JSON.
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			next.ServeHTTP(w, r)
		},
	)
}

// ValidateMiddleware checks the request for valid JSON web tokens
func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		errorResponse:= response.Error{}

		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("THERE WAS AN ERROR")
					}
					return []byte("way"), nil
				})
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					_ = json.NewEncoder(w).Encode(
						errorResponse.ErrorResponse("Error", 500, "System error"),
					)
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					w.WriteHeader(http.StatusBadRequest)
					_ = json.NewEncoder(w).Encode(
						errorResponse.ErrorResponse("Error", 500, "Invalid authorization token"),
					)
					return
				}
			}
		} else {
			_ = json.NewEncoder(w).Encode(
				errorResponse.ErrorResponse("Error", 500, "An authorization header is required"),
			)
			return
		}
	})
}
