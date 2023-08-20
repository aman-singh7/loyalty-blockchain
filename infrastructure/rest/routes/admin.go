package routes

import (
	"github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/controllers/admin"
	"github.com/labstack/echo/v4"
)

func AdminRoutes(g *echo.Group, controller *admin.Controller) {
	routerAdmin := g.Group("/admin")
	{
		routerAdmin.POST("/addMember", controller.RegisterMember)
	}
}
