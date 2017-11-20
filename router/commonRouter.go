package router

import (
	"bitbucket.org/DanielFrag/gestor-de-ponto/controller"
)

//GetCommonRoutes return the routes for common users and manager users
func GetCommonRoutes() []Route {
	return []Route{
		Route{
			Name:        "login",
			Method:      "POST",
			Path:        "/login",
			HandlerFunc: controller.UserLogin,
		},
		Route{
			Name:        "logout",
			Method:      "POST",
			Path:        "/logout",
			HandlerFunc: controller.TokenCheckerMiddleware(controller.UserMiddleware(controller.UserLogout)),
		},
	}
}
