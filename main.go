package main

import (
	"log"
	"net/http"

	"github.com/Asprilla24/rest-ums/router"
)

func main() {
	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
