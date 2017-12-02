package router

import (
	"net/http"

	"bitbucket.org/DanielFrag/gestor-de-ponto/controller"
	"github.com/gorilla/mux"
)

//Route struct that contain the details of a single api's resource
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

//NewRouter exports the api routes with his handler functions
func NewRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	adminRoutes := GetAdminRoutes()
	userRoutes := GetUserRoutes()
	commonRoutes := GetCommonRoutes()
	for _, route := range commonRoutes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			HandlerFunc(route.HandlerFunc)
	}
	for _, route := range adminRoutes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			HandlerFunc(route.HandlerFunc)
	}
	for _, route := range userRoutes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			HandlerFunc(route.HandlerFunc)
	}
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./www/")))
	return controller.CorsSetup(router)
}
