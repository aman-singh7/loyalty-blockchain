package user

import (
	"net/http"

	"github.com/aman-singh7/loyalty-blockchain/application/core/user"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	UserService *user.Service
}

// TODO: replace validator with echo.validator
// TODO: make a transaction id generator

func (c *Controller) GetUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, echo.Map{"message": "get user api not implemented"})
}

func (c *Controller) Discount(ctx echo.Context) error {
	var request DiscountRequest
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := validator.New().Struct(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	// TODO: call service func
	return nil
}

func (c *Controller) PurchaseProduct(ctx echo.Context) error {
	var request PurchaseProductRequest
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := validator.New().Struct(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	// TODO: call service func
	return nil
}

func (c *Controller) PurchaseCoupon(ctx echo.Context) error {
	var request PurchaseCouponRequest
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := validator.New().Struct(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	// TODO: call service func
	return nil
}

func (c *Controller) ReferralReward(ctx echo.Context) error {
	var request ReferralRewardRequest
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := validator.New().Struct(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	// TODO: call service func
	return nil
}
