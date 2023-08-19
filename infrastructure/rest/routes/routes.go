package routes

import (
	"github.com/aman-singh7/loyalty-blockchain/application/api"
	"github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/adapter"
	"github.com/labstack/echo/v4"
)

func ApplicationV1Router(e *echo.Echo, api *api.Api) {
	routerV1 := e.Group("/v1")
	{
		UserRoutes(routerV1, adapter.UserAdapter(api))
	}
}
