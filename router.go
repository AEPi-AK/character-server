package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	var routes = Routes{
		Route{
			"CharacterCreate",
			"POST",
			"/characters/create",
			CharacterCreate,
		},
		Route{
			"CharacterUpdate",
			"POST",
			"/characters/update",
			CharacterUpdate,
		},
		Route{
			"CharacterShow",
			"GET",
			"/characters/{identifier}",
			CharacterShow,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
