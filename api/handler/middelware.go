package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"log"
	"net/http"
	"strings"
	"time"
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

// Set's response Content Type to JSON
func ResponseHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Logs information about the current request
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s\t%s", r.Method, r.URL.Path, )
		next.ServeHTTP(w, r)
		log.Print("Responded after: " + time.Since(start).String() + "\n")
	})
}

// Checks for valid Json Web Tokens for Administrator only routes
func AdminAuthMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("THERE WAS AN ERROR")
					}
					return []byte(webTokens.FolioAuthAdminKey), nil
				})
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(
						response.ErrorResponse("Error", 500, "System error"),
					)
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(
						response.ErrorResponse("Error", 500, "invalid authorization token"),
					)
					return
				}
			}
		} else {
			json.NewEncoder(w).Encode(
				response.ErrorResponse("Error", 500, "authorization is required"),
			)
		}
	})
}

// Checks for valid Json Web Tokens for Service only routes
func ServiceAuthMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("INVALID TOKEN")
					}
					return []byte(webTokens.FolioAuthKey), nil
				})
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(
						response.ErrorResponse("Error", 500, "System error"),
					)
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(
						response.ErrorResponse("Error", 500, "invalid authorization token"),
					)
					return
				}
			}
		} else {
			json.NewEncoder(w).Encode(
				response.ErrorResponse("Error", 500, "authorization is required"),
			)
		}
	})
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
