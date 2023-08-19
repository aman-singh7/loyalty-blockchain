package brand

import (
	"net/http"

	"github.com/aman-singh7/loyalty-blockchain/application/price"
	"github.com/aman-singh7/loyalty-blockchain/domain/brand"
	brandRepo "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/brand"
	"github.com/labstack/echo/v4"
)

type Service struct {
	repo *brandRepo.Repository
}

func NewService(repo *brandRepo.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s* Service) CouponPrice(request brand.CouponPriceRequest) (int, error) {
	price, err := price.PredictCouponPrice(/* TODO: all the params */)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return price, nil
}

func (s* Service) CreateCoupon(request brand.CreateCouponRequest) error {
	// TODO: validate brand
	// TODO: call sol_create_coupon()
	return nil
}

func (s* Service) RedeemTokens(request brand.RedeemTokensRequest) error {
	// TODO: validate brand
	// TODO: call sol_redeem_tokens()
	return nil
}