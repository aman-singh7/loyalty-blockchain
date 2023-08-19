package user

import (
	"net/http"
	"strconv"
	"time"

	"github.com/aman-singh7/loyalty-blockchain/application/api"
	"github.com/aman-singh7/loyalty-blockchain/domain/coupon"
	"github.com/aman-singh7/loyalty-blockchain/domain/user"
	userRepo "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/user"
	"github.com/labstack/echo/v4"
)

type Service struct {
	repo *userRepo.Repository
	api  *api.Service
}

func NewService(repo *userRepo.Repository, api *api.Service) *Service {
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
	if(strconv.Itoa(request.Coupon.CouponID) != "" && request.Tokens != 0) {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "coupon and tokens cannot be in same transaction"})
	}
	if request.Tokens != 0 {
		if err := s.api.PurchaseProductWithToken(request.TransactionID, request.Tokens); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
		}
	} else if request.Coupon.CouponID != 0 {
		if err := s.api.PurchaseProductWithCoupon(request.TransactionID, request.Coupon, request.Coupon.IssuerBusiness); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
		}
	}
	// TODO: decide a reward amount
	rewardAmount := request.Price / 100
	if err := s.api.RewardToken(request.TransactionID, rewardAmount); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Reward Generation Failed"})
	}
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
