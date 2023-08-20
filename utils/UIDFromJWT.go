package utils

import (
	"net/http"

	"github.com/aman-singh7/loyalty-blockchain/application/security/jwt"
	"github.com/labstack/echo/v4"
)

func GetUidFromJWT(token string) (string, error) {
	claims, err := jwt.GetClaimsAndVerifyToken(token, "access")
	if err != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return claims["id"].(string), nil
}
