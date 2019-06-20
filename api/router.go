package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"way/api/handlers"
	"way/pkg/logger"
	"way/pkg/router"
)

// NewRouter creates a new router for https requests to the API
func NewRouter() *mux.Router {
	httpsRouter := mux.NewRouter()
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunction
		handler = logger.Traffic(handler, route.Name)

		httpsRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return httpsRouter
}

// add all router used go here
var routes = router.Routes{
	// index
	router.Route{
		Name:            "Index",
		Method:          http.MethodGet,
		Pattern:         "/",
		HandlerFunction: handlers.Index,
	},

	// sign up
	router.Route{
		Name:            "SignUp",
		Method:          http.MethodPost,
		Pattern:         "/users",
		HandlerFunction: handlers.SignUp,
	},

	// verify
	router.Route{
		Name:            "Verify",
		Method:          http.MethodPost,
		Pattern:         "/users/verify",
		HandlerFunction: handlers.Verify,
	},
	// log in
	router.Route{
		Name:            "Log In",
		Method:          http.MethodPost,
		Pattern:         "/users/login",
		HandlerFunction: handlers.Login,
	},
}
