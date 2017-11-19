package router

import (
	"bitbucket.org/DanielFrag/gestor-de-ponto/controller"
)

//GetUserRoutes return the routes for common users
func GetUserRoutes() []Route {
	return []Route{
		Route{
			Name:        "teste",
			Method:      "POST",
			Path:        "/user-login",
			HandlerFunc: controller.UserLogin,
		},
	}
}
