package user

import (
	"math/big"
	"net/http"
	"time"

	"github.com/aman-singh7/loyalty-blockchain/application/api"
	"github.com/aman-singh7/loyalty-blockchain/domain/coupon"
	"github.com/aman-singh7/loyalty-blockchain/domain/user"
	userRepo "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/user"
	"github.com/aman-singh7/loyalty-blockchain/utils"
	"github.com/ethereum/go-ethereum/common"
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

func (s *Service) FetchBalance(address common.Address) (int, error) {
	bal, err := s.api.FetchAccountBalance(*big.NewInt(1), address, address)
	if err != nil {
		return 0, err
	}

	return bal, nil
}

func (s *Service) Discount(userCoupon *coupon.Coupon) (big.Int, error) {
	if isActive := userCoupon.ExpiryDate.Cmp(utils.BigInt(time.Now().Second())) == 1; !isActive {
		return *big.NewInt(0), echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "coupon has expired"})
	}
	return userCoupon.Discount, nil
}

// TODO: differentiate between purchase with Coupon/Token + Cash
func (s *Service) PurchaseProduct(request *user.PurchaseProductRequest) error {
	if request.Coupon != nil && request.Tokens.Cmp(big.NewInt(0)) != 0 {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "coupon and tokens cannot be in same transaction"})
	}

	// TODO: validate user
	if request.Coupon != nil {
		_, err := s.Discount(request.Coupon)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "coupon has expired"})
		}
		if request.Coupon.Type == coupon.UNIQUE {
			if &request.Coupon.ProductId != &request.ProductID {
				return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "coupon is not allowed on this product"})
			}
		} else {
			if request.Coupon.ProductCategory.Cmp(&request.ProductCategory) == 0 {
				return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "coupon is not allowed on this product"})
			} else if request.Coupon.ThresholdValue.Cmp(&request.Price) == -1 {
				return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "coupon is not allowed on this product"})
			}
		}

		if request.Coupon.CouponID.String() != "" {
			if err := s.api.PurchaseProductWithCoupon(request.TransactionID, *request.Coupon, request.Coupon.IssuerBusiness); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
			}
		}
	}
	if request.Tokens.Cmp(big.NewInt(0)) != 0 {
		if err := s.api.PurchaseProductWithToken(request.TransactionID, request.Tokens); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
		}
	}
	// TODO: decide a reward amount
	rewardAmount := *(big.NewInt(0)).Div(&request.Price, big.NewInt(100))
	if err := s.api.RewardToken(request.TransactionID, request.UserAddress, rewardAmount); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Reward Generation Failed"})
	}
	return nil
}

func (s *Service) PurchaseCoupon(request *user.PurchaseCouponRequest) error {
	// TODO: validate user
	if err := s.api.PurchaseCoupon(request.TransactionID, request.Coupon, request.Count, request.Coupon.IssuerBusiness); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
	}
	return nil
}

func (s *Service) ReferralReward(request *user.ReferralRewardRequest) error {
	// TODO: validate user1 and user2
	if err := s.api.RewardToken(request.TransactionID, request.UserAddress, request.Tokens); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Reward Generation Failed"})
	}
	return nil
}
