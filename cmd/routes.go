package main

import (
	"github.com/gojou/goofmongo/pkg/handlers"

	"github.com/gorilla/mux"
)

func routes(r *mux.Router) {
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/person", handlers.Person)

}
