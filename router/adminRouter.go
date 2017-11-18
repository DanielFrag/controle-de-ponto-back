package router

import (
	"bitbucket.org/DanielFrag/gestor-de-ponto/controller"
)

func GetAdminRoutes() []Route {
	return []Route {
		Route {
			Name: "adminInsertUser",
			Method: "POST",
			Path: "/admin-insert-user",
			HandlerFunc: controller.TokenCheckerMiddleware(controller.InsertUser),
		},
	}
}