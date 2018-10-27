package router

import (
	"log"
	"net/http"

	"github.com/Asprilla24/rest-ums/controller"

	"github.com/Asprilla24/rest-ums/repository"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var user = &controller.User{Repository: repository.User{}}

var routes = []Route{
	Route{
		Name:        "Login",
		Method:      "GET",
		Pattern:     "/api/login",
		HandlerFunc: user.Login,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.HandlerFunc
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
