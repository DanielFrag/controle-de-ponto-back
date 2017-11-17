package router

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
    Name        string
	Method      string
	Path		string
	HandlerFunc http.HandlerFunc
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	adminRoutes := GetAdminRoutes()
	userRoutes := GetUserRoutes()
	for _, route := range adminRoutes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	for _, route := range userRoutes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}