package router

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	Route{
		Name:        "Login",
		Method:      "GET",
		Pattern:     "/api/login",
		HandlerFunc: user.GetUser,
	},
}
