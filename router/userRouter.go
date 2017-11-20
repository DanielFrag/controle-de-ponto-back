package router

import (
	"bitbucket.org/DanielFrag/gestor-de-ponto/controller"
)

//GetUserRoutes return the routes for common users
func GetUserRoutes() []Route {
	return []Route{
		Route{
			Name:        "nowTimestamp",
			Method:      "POST",
			Path:        "/user/register-now",
			HandlerFunc: controller.TokenCheckerMiddleware(controller.UserMiddleware(controller.RegisterNowTimestamp)),
		},
		Route{
			Name:        "customTimestamp",
			Method:      "POST",
			Path:        "/user/register-date",
			HandlerFunc: controller.TokenCheckerMiddleware(controller.UserMiddleware(controller.RegisterCustomTimestamp)),
		},
		Route{
			Name:        "registersByDate",
			Method:      "POST",
			Path:        "/user/registers-by-date",
			HandlerFunc: controller.TokenCheckerMiddleware(controller.UserMiddleware(controller.GetUserRegisterByDate)),
		},
	}
}
