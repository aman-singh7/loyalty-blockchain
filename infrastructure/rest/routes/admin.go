package routes

import (
	customjwt "github.com/aman-singh7/loyalty-blockchain/application/security/jwt"
	"github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/controllers/admin"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func AdminRoutes(g *echo.Group, controller *admin.Controller) {
	routerAdmin := g.Group("/admin")
	// JWT Middleware
	config := customjwt.GetJWTConfig()
	routerAdmin.Use(echojwt.WithConfig(config))
	{
		routerAdmin.POST("/addMember", controller.RegisterMember)
	}
}
