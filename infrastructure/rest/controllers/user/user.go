package user

import (
	"net/http"

	user "github.com/aman-singh7/loyalty-blockchain/application/core/user"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Controller struct {
	UserService *user.Service
}

// TODO: remove this
var uAddress string = "0x4593C14AE1c7E4e0304614fBa80c94C02495df22"

// TODO: replace validator with echo.validator
// TODO: make a transaction id generator

func (c *Controller) FetchBalance(ctx echo.Context) error {
	// TODO: extract address from jwt
	ADDRESS := viper.GetString("Address")

	bal, err := c.UserService.FetchBalance(common.HexToAddress(ADDRESS))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "success", "balance": bal})
}

func (c *Controller) Discount(ctx echo.Context) error {
	var request DiscountRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	ctx.Validate(request)
	// TODO: user address from jwt
	userAddress := common.HexToAddress(uAddress)
	serviceRequest := toDomainDiscountRequest(&request, userAddress)
	discount, err := c.UserService.Discount(&serviceRequest.Coupon)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "success", "discount": discount})
}

func (c *Controller) PurchaseProduct(ctx echo.Context) error {
	var request PurchaseProductRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	// TODO: user address from jwt
	userAddress := common.HexToAddress(uAddress)
	serviceRequest := toDomainPurchaseProductRequest(&request, userAddress)
	err := c.UserService.PurchaseProduct(serviceRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "success"})
}

func (c *Controller) PurchaseCoupon(ctx echo.Context) error {
	var request PurchaseCouponRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	// TODO: user address from jwt
	userAddress := common.HexToAddress(uAddress)
	serviceRequest := toDomainPurchaseCouponRequest(&request, userAddress)
	err := c.UserService.PurchaseCoupon(serviceRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "success"})
}

func (c *Controller) ReferralReward(ctx echo.Context) error {
	var request ReferralRewardRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	// TODO: user address from jwt
	userAddress := common.HexToAddress(uAddress)
	serviceRequest := toDomainReferralRewardRequest(&request, userAddress)
	err := c.UserService.ReferralReward(serviceRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "success"})
}

func (c *Controller) GetAllCoupons(ctx echo.Context) error {
	coupons, err := c.UserService.GetAllCoupons()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "success", "coupons": coupons})
}

func (c *Controller) GetBrandCoupons(ctx echo.Context) error {
	var request GetBrandCouponsRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	brandAddress := common.HexToAddress(request.BrandID)
	coupons, err := c.UserService.GetBrandCoupons(brandAddress)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "success", "coupons": coupons})
}
	