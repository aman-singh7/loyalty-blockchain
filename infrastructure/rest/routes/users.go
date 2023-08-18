package routes

import (
	userController "github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/controllers/user"
	"github.com/labstack/echo/v4"
)

func UserRoutes(g *echo.Group, controller *userController.Controller) {
	routerUser := g.Group("/user")
	{
		routerUser.GET("/", controller.GetUser)
	}
}
