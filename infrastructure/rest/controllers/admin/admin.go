package admin

import (
	"net/http"

	"github.com/aman-singh7/loyalty-blockchain/application/core/admin"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	AdminService *admin.Service
}

func (c *Controller) RegisterMember(ctx echo.Context) error {
	// TODO: add check on account type
	// Consumer -> 1 Brand -> 2

	var request RegisterMemberRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	addr := common.HexToAddress(request.Addr)

	err := c.AdminService.RegisterMember(addr, request.Type)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusAccepted, echo.Map{
		"message": "your request has been submitted",
	})
}
