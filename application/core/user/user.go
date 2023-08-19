package user

import (
	"net/http"
	"time"

	"github.com/aman-singh7/loyalty-blockchain/application/api"
	"github.com/aman-singh7/loyalty-blockchain/domain/coupon"
	"github.com/aman-singh7/loyalty-blockchain/domain/user"
	userRepo "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/user"
	"github.com/labstack/echo/v4"
)

type Service struct {
	repo *userRepo.Repository
	api  *api.Api
}

func NewService(repo *userRepo.Repository, api *api.Api) *Service {
	return &Service{
		repo: repo,
		api:  api,
	}
}

// TODO: make enum for errors

func (s *Service) Discount(userCoupon coupon.Coupon) (int, error) {
	if isActive := (userCoupon.ExpiryDate < time.Now().Second()); !isActive {
		return 0, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "coupon has expired"})
	}
	return userCoupon.Discount, nil
}

// TODO: differentiate between purchase with Coupon/Token + Cash
func (s *Service) PurchaseProduct(request user.PurchaseProductRequest) error {
	// TODO: validate user
	_, err := s.Discount(request.Coupon)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "coupon has expired"})
	}
	if request.Coupon.Type == coupon.UNIQUE {
		if request.Coupon.ProductId != request.ProductID {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "coupon is not allowed on this product"})
		}
	} else {
		if request.Coupon.ProductCategory != request.ProductCategory {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "coupon is not allowed on this product"})
		} else if request.Coupon.ThresholdValue < request.Price {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "coupon is not allowed on this product"})
		}
	}
	// TODO: call sol_purchase_product()
	// TODO: give reward tokens
	return nil
}

func (s *Service) PurchaseCoupon(request user.PurchaseCouponRequest) error {
	// TODO: validate user
	// TODO: call sol_purchase_coupon()
	return nil
}

func (s *Service) ReferralReward(request user.ReferralRewardRequest) error {
	// TODO: validate user1 and user2
	// TODO: give reward tokens
	return nil
}
