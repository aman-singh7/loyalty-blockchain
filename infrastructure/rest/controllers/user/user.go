package user

import (
	"net/http"

	"github.com/aman-singh7/loyalty-blockchain/application/core/user"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	UserService *user.Service
}

func (c *Controller) GetUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, echo.Map{"message": "get user api not implemented"})
}
