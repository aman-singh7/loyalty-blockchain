package routes

import (
	customjwt "github.com/aman-singh7/loyalty-blockchain/application/security/jwt"
	brandController "github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/controllers/brand"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func BrandRoutes(g *echo.Group, controller *brandController.Controller) {
	routerBrand := g.Group("/brand")
	// JWT Middleware
	config := customjwt.GetJWTConfig()
	routerBrand.Use(echojwt.WithConfig(config))
	{
		routerBrand.GET("/", controller.GetBrand)
		routerBrand.GET("/price", controller.CouponPrice)
		routerBrand.POST("/coupon", controller.CreateCoupon)
		routerBrand.POST("/redeem", controller.RedeemTokens)
	}
}
