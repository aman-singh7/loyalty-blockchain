package api

import (
	"context"
	"net/http"

	"github.com/aman-singh7/loyalty-blockchain/domain/coupon"
	"github.com/aman-singh7/loyalty-blockchain/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Service struct {
	api  *Api
	opts *utils.TransactOpts
}

func NewService(api *Api, opts *utils.TransactOpts) *Service {
	return &Service{
		api:  api,
		opts: opts,
	}
}

func (s *Service) FetchAccountBalance(transactionId int, address common.Address, from common.Address) (int, error) {
	opts := bind.CallOpts{
		Pending: false,
		From:    from,
		Context: context.Background(),
	}

	balance, err := s.api.GetAccountBalance(&opts, utils.BigInt(transactionId), address)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"message": "fetch balance failed",
		})
	}

	return utils.ToInt(balance), nil
}

func (s *Service) PurchaseProductWithCoupon(transactionId int, coupon coupon.Coupon, address common.Address) error {
	_, err := s.api.ApiTransactor.ConsumerCouponPayment(&bind.TransactOpts{}, utils.BigInt(transactionId), utils.BigInt(coupon.HoldingID), address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
	}
	return nil
}

func (s *Service) PurchaseProductWithToken(transactionId int, tokens int) error {
	_, err := s.api.ApiTransactor.ConsumerTokenPayment(&bind.TransactOpts{}, utils.BigInt(transactionId), utils.BigInt(tokens))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
	}
	return nil
}

func (s *Service) PurchaseCoupon(transactionId int, coupon coupon.Coupon, count int, address common.Address) error {
	_, err := s.api.ApiTransactor.PurchaseCoupon(&bind.TransactOpts{}, utils.BigInt(transactionId), address, utils.BigInt(coupon.CouponID), utils.BigInt(count))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
	}
	return nil
}

func (s *Service) RewardToken(transactionId int, amount int) error {
	_, err := s.api.ApiTransactor.MintTokens(&bind.TransactOpts{}, utils.BigInt(transactionId), utils.BigInt(amount))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
	}
	return nil
}

func (s *Service) CreateCoupon(transactionId int, coupon coupon.Coupon, cost int) error {
	_, err := s.api.ApiTransactor.CreateCoupons(&bind.TransactOpts{}, utils.BigInt(transactionId), coupon.IssuerBusiness, utils.BigInt(coupon.Count), utils.BigInt(coupon.SuperCoins), utils.BigInt(coupon.Discount), utils.BigInt(coupon.ProductCategory), utils.BigInt(coupon.ThresholdValue), utils.BigInt((coupon.ProductId)), uint8(coupon.Type), utils.BigInt(coupon.ExpiryDate), utils.BigInt(cost))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
	}
	return nil
}

func (s *Service) RedeemTokens(transactionId int, amount int) error {
	_, err := s.api.ApiTransactor.RedeemTokens(&bind.TransactOpts{}, utils.BigInt(amount))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
	}
	return nil
}

func (s *Service) RegisterMember(memberAddr common.Address, accountType uint8) error {
	// 1 -> Customer 2 -> Brand
	PRIVATE_KEY := viper.GetString("PrivateKey")
	opts := s.opts.GetOpts(PRIVATE_KEY)
	_, err := s.api.RegisterMember(opts, utils.BigInt(2), memberAddr, accountType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"message": "register member failed",
		})
	}

	return nil
}
