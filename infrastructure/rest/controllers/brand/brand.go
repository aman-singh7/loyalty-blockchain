package brand

import (
	"net/http"

	brand "github.com/aman-singh7/loyalty-blockchain/application/core/brand"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	BrandService *brand.Service
}

func (c *Controller) GetBrand(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, echo.Map{"message": "get brand api not implemented"})
}

func (c *Controller) CouponPrice(ctx echo.Context) error {
	var request CouponPriceRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	bAddress, err := c.BrandService.GetAddress(request.JWT)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	brandAddress := common.HexToAddress(bAddress)
	serviceRequest := toDomainCouponPriceRequest(&request, brandAddress)
	price, err := c.BrandService.CouponPrice(serviceRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "success", "price": price})
}

func (c *Controller) CreateCoupon(ctx echo.Context) error {
	var request CreateCouponRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	bAddress, err := c.BrandService.GetAddress(request.JWT)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	brandAddress := common.HexToAddress(bAddress)
	serviceRequest := toDomainCreateCouponRequest(&request, brandAddress)
	if err := c.BrandService.CreateCoupon(serviceRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "success"})
}

func (c *Controller) RedeemTokens(ctx echo.Context) error {
	var request RedeemTokensRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	bAddress, err := c.BrandService.GetAddress(request.JWT)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	brandAddress := common.HexToAddress(bAddress)
	serviceRequest := toDomainRedeemTokensRequest(&request, brandAddress)
	if err := c.BrandService.RedeemTokens(serviceRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "success"})
}
