package router

import (
	"bitbucket.org/DanielFrag/gestor-de-ponto/controller"
)

func GetCommonRoutes() []Router {
	return []Route {
		Route {
			Name: "login",
			Method: "POST",
			Path: "/login",
			HandlerFunc: controller.UserLogin,
		},
	}
}