package routes

import (
	brandController "github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/controllers/brand"
	"github.com/labstack/echo/v4"
)

func BrandRoutes(g *echo.Group, controller *brandController.Controller) {
	routerUser := g.Group("/brand")
	{
		routerUser.GET("/", controller.GetBrand)
		routerUser.GET("/price", controller.CouponPrice)
		routerUser.POST("/coupon", controller.CreateCoupon)
		routerUser.POST("/redeem", controller.RedeemTokens)
	}
}
