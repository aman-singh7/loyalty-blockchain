package routes

import (
	userController "github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/controllers/auth"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(g *echo.Group, controller *userController.Controller) {
	routerUser := g.Group("/auth")
	{
		routerUser.GET("/signin", controller.RecordUser)
		routerUser.GET("/token", controller.GenerateToken)
	}
}
