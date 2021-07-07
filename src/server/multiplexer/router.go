package multiplexer

import (
	"github.com/gorilla/mux"
	"net/http"
	"way/src/server/handler"
	"way/src/server/route"
)

const (
	Index      = "Index"
	CreateUser = "Create User"
	Login      = "Log In"
	Logout     = "Log Out"
	VerifyUser = "Verify user"
	AddContact = "Add contact"
	GetFriends = "Get Friends"
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
	// add contact
	route.Route{
		Name:            AddContact,
		Method:          http.MethodPost,
		Pattern:         "/users/{user_id}/contacts",
		HandlerFunction: handler.AddContact,
	},

	// get friends
	route.Route{
		Name:            GetFriends,
		Method:          http.MethodGet,
		Pattern:         "/users/friends",
		HandlerFunction: handler.GetFriends,
	},
}

// Router creates a new route for https requests to the API
func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		if route.Name == AddContact {
			router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler.UserAuthMiddleware(route.HandlerFunction))
			continue
		}

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunction)

	}

	return router
}
