package router

import (
	"bitbucket.org/DanielFrag/gestor-de-ponto/controller"
)

//GetAdminRoutes return the routes for manager users
func GetAdminRoutes() []Route {
	return []Route{
		Route{
			Name:        "adminInsertUser",
			Method:      "POST",
			Path:        "/admin/insert-user",
			HandlerFunc: controller.TokenCheckerMiddleware(controller.UserMiddleware(controller.AdminUserValidatorMiddleware(controller.InsertUser))),
		},
	}
}
