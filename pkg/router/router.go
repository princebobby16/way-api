package router

import (
	"net/http"
)

// Route is the type for the router handled by this Api
type Route struct {
	Name            string
	Method          string
	Pattern         string
	HandlerFunction http.HandlerFunc
}

// Routes is an array of router
type Routes []Route
