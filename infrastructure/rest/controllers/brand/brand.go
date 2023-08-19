package brand

import (
	"net/http"

	brand "github.com/aman-singh7/loyalty-blockchain/application/core/admin"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	BrandService *brand.Service
}

func (c *Controller) CreateCoupon(ctx echo.Context) error {
	var request CreateCouponRequest
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := validator.New().Struct(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	// TODO: call service
	return nil
}

func (c *Controller) RedeemTokens(ctx echo.Context) error {
	var request RedeemTokensRequest
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := validator.New().Struct(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	// TODO: call service
	return nil
}
