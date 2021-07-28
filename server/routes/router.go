package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"way/server/handler"
)

const (
	Index      = "Index"
	CreateUser = "Create User"
	Login      = "Log In"
	Logout     = "Log Out"
	VerifyUser = "Verify user"
	AddContact = "Add contact"
	GetFriends = "Get Friends"
	requestPIN ="Get PIN"
)

// Route is the type for the route handled by this Api
type Route struct {
	Name            string
	Method          string
	Pattern         string
	HandlerFunction http.HandlerFunc
}

// Routes is an array of route
type Routes []Route

var routes = Routes{
	// index
	Route{
		Name:            Index,
		Method:          http.MethodGet,
		Pattern:         "/",
		HandlerFunction: handler.Index,
	},

	// sign up
	Route{
		Name:            CreateUser,
		Method:          http.MethodPost,
		Pattern:         "/users",
		HandlerFunction: handler.CreateUser,
	},

	// verify
	Route{
		Name:            VerifyUser,
		Method:          http.MethodPost,
		Pattern:         "/users/verify",
		HandlerFunction: handler.Verify,
	},
	// log in
	Route{
		Name:            Login,
		Method:          http.MethodPost,
		Pattern:         "/users/login",
		HandlerFunction: handler.Login,
	},
	// add contact
	Route{
		Name:            AddContact,
		Method:          http.MethodPost,
		Pattern:         "/users/{user_id}/contacts",
		HandlerFunction: handler.AddContact,
	},

	// get friends
	Route{
		Name:            GetFriends,
		Method:          http.MethodGet,
		Pattern:         "/users/friends",
		HandlerFunction: handler.GetFriends,
	},
	// request verification PIN
	Route{
		Name:            requestPIN,
		Method:          http.MethodPost,
		Pattern:         "/users/pin",
		HandlerFunction: handler.RequestPIN,
	},
}

// Router creates a new route for https requests to the API
func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, _route := range routes {

		if _route.Name == AddContact {
			router.Methods(_route.Method).Path(_route.Pattern).Name(_route.Name).Handler(handler.UserAuthMiddleware(_route.HandlerFunction))
			continue
		}

		router.Methods(_route.Method).Path(_route.Pattern).Name(_route.Name).Handler(_route.HandlerFunction)

	}

	return router
}
