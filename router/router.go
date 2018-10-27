package router

import (
	"net/http"

	"github.com/Asprilla24/rest-ums/controller"

	"github.com/Asprilla24/rest-ums/repository"

	"github.com/gorilla/mux"
)

var user = &controller.User{Repository: repository.User{}}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
