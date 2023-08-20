package user

import (
	"math/big"
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
var uAddress string = ""

// TODO: replace validator with echo.validator
// TODO: make a transaction id generator

func (c *Controller) FetchBalance(ctx echo.Context) error {
	// TODO: extract address from jwt
	ADDRESS := viper.GetString("Address")

	bal, err := c.UserService.FetchBalance(common.HexToAddress(ADDRESS))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{"balance": bal})
}

func (c *Controller) Discount(ctx echo.Context) error {
	var request DiscountRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	ctx.Validate(&request)
	// TODO: user address from jwt
	userAddress := common.HexToAddress(uAddress);
	// TODO: generate transaction id
	transactionId := *big.NewInt(2)
	serviceRequest := toDomainDiscountRequest(&request, transactionId, userAddress)
	c.UserService.Discount(&serviceRequest.Coupon)
	return nil
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
	userAddress := common.HexToAddress(uAddress);
	// TODO: generate transaction id
	transactionId := *big.NewInt(2)
	serviceRequest := toDomainPurchaseProductRequest(&request, transactionId, userAddress)
	c.UserService.PurchaseProduct(serviceRequest)
	return nil
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
	userAddress := common.HexToAddress(uAddress);
	// TODO: generate transaction id
	transactionId := *big.NewInt(2)
	serviceRequest := toDomainPurchaseCouponRequest(&request, transactionId, userAddress)
	c.UserService.PurchaseCoupon(serviceRequest)
	return nil
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
	userAddress := common.HexToAddress(uAddress);
	// TODO: generate transaction id
	transactionId := *big.NewInt(2)
	serviceRequest := toDomainReferralRewardRequest(&request, transactionId, userAddress)
	c.UserService.ReferralReward(serviceRequest)
	return nil
}
