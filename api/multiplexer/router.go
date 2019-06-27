package multiplexer

import (
	"github.com/gorilla/mux"
	"net/http"
	"way/api/handler"
	"way/pkg/route"
)

const (
	Index           = "Index"
	CreateUser      = "Create User"
	Login           = "Log In"
	Logout          = "Log Out"
	VerifyUser="Verify user"
)

var routes = route.Routes{
	// index
	route.Route{
		Name:            Index,
		Method:          http.MethodGet,
		Pattern:         "/",
		HandlerFunction: handler.Index,
	},

	// sign up
	route.Route{
		Name:            CreateUser,
		Method:          http.MethodPost,
		Pattern:         "/users",
		HandlerFunction: handler.CreateUser,
	},

	// verify
	route.Route{
		Name:            VerifyUser,
		Method:          http.MethodPost,
		Pattern:         "/users/verify",
		HandlerFunction: handler.Verify,
	},
	// log in
	route.Route{
		Name:            Login,
		Method:          http.MethodPost,
		Pattern:         "/users/login",
		HandlerFunction: handler.Login,
	},
}


// Router creates a new route for https requests to the API
func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)


	for _, route := range routes {

		if false {
			router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler.ServiceAuthMiddleware(route.HandlerFunction))
			continue
		}

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunction)


	}

	return router
}

