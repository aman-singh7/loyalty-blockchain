package auth

import (
	"net/http"

	"github.com/aman-singh7/loyalty-blockchain/application/core/auth"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	AuthService *auth.Service
}

func (c *Controller) CreateUser(ctx echo.Context) error {
	var request CreateUserRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	serviceRequest := toDomainCreateUserRequest(&request)
	authData, err := c.AuthService.CreateUser(serviceRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, authData)
}

func (c *Controller) GenerateToken(ctx echo.Context) error {
	var request GenerateTokenRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	serviceRequest := toDomainGenerateTokenRequest(&request)
	jwtToken, err := c.AuthService.GenerateAccessTokenFromRefreshToken(serviceRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, jwtToken)
}
