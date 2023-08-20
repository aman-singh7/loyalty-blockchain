package routes

import (
	customjwt "github.com/aman-singh7/loyalty-blockchain/application/security/jwt"
	userController "github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/controllers/user"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func UserRoutes(g *echo.Group, controller *userController.Controller) {
	routerUser := g.Group("/user")
	// JWT Middleware
	config := customjwt.GetJWTConfig()
	routerUser.Use(echojwt.WithConfig(config))
	{
		routerUser.GET("/balance", controller.FetchBalance)
		routerUser.GET("/discount", controller.Discount)
		routerUser.POST("/purchase", controller.PurchaseProduct)
		routerUser.POST("/coupon", controller.PurchaseCoupon)
		routerUser.POST("/referral", controller.ReferralReward)
		routerUser.GET("/coupons/all", controller.GetAllCoupons)
		routerUser.GET("/coupons", controller.GetBrandCoupons)
	}
}
