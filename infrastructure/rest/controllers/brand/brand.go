package brand

import (
	"net/http"

	brand "github.com/aman-singh7/loyalty-blockchain/application/core/brand"
	"github.com/labstack/echo/v4"
)

// TODO: replace validator with echo.validator

type Controller struct {
	BrandService *brand.Service
}

func (c *Controller) GetBrand(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, echo.Map{"message": "get brand api not implemented"})
}

func (c *Controller) CouponPrice(ctx echo.Context) error {
	var request CouponPriceRequest
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	// TODO: call calculating price
	return nil
}

func (c *Controller) CreateCoupon(ctx echo.Context) error {
	var request CreateCouponRequest
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	// TODO: call service func
	return nil
}

func (c *Controller) RedeemTokens(ctx echo.Context) error {
	var request RedeemTokensRequest
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	// TODO: call service func
	return nil
}
