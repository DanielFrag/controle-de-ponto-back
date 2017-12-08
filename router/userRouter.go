package router

import (
	"bitbucket.org/DanielFrag/gestor-de-ponto/controller"
)

//GetUserRoutes return the routes for common users
func GetUserRoutes() []Route {
	return []Route{
		Route{
			Name:        "customTimestamp",
			Method:      "POST",
			Path:        "/user/register-date",
			HandlerFunc: controller.TokenCheckerMiddleware(controller.UserMiddleware(controller.RegisterDate)),
		},
		Route{
			Name:        "registersByDate",
			Method:      "GET",
			Path:        "/user/registers/{start}/{finish}",
			HandlerFunc: controller.TokenCheckerMiddleware(controller.UserMiddleware(controller.GetUserRegistersByDate)),
		},
		Route{
			Name:        "removeRegister",
			Method:      "DELETE",
			Path:        "/user/remove-register/{id}",
			HandlerFunc: controller.TokenCheckerMiddleware(controller.UserMiddleware(controller.RemoveDateRegister)),
		},
	}
}
