package router

import (
	"bitbucket.org/DanielFrag/gestor-de-ponto/controller"
)

func GetUserRoutes() []Route{
	return []Route {
		Route {
			Name: "teste",
			Method: "POST",
			Path: "/user-login",
			HandlerFunc: controller.UserLogin,
		},
	}
}