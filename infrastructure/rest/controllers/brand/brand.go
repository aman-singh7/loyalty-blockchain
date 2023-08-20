package brand

import (
	"math/big"
	"net/http"

	brand "github.com/aman-singh7/loyalty-blockchain/application/core/brand"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
)

// TODO: remove this
var bAddress string = ""

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
	// TODO: user address from jwt
	brandAddress := common.HexToAddress(bAddress)
	// TODO: generate transaction id
	transactionId := *big.NewInt(2)
	serviceRequest := toDomainCouponPriceRequest(&request, transactionId, brandAddress)
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
	// TODO: user address from jwt
	brandAddress := common.HexToAddress(bAddress)
	// TODO: generate transaction id
	transactionId := *big.NewInt(2)
	serviceRequest := toDomainCreateCouponRequest(&request, transactionId, brandAddress)
	err := c.BrandService.CreateCoupon(serviceRequest)
	if err != nil {
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
	// TODO: user address from jwt
	brandAddress := common.HexToAddress(bAddress)
	// TODO: generate transaction id
	transactionId := *big.NewInt(2)
	serviceRequest := toDomainRedeemTokensRequest(&request, transactionId, brandAddress)
	err := c.BrandService.RedeemTokens(serviceRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "success"})
}
