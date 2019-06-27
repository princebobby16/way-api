package route

import (
	"net/http"
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
