package api

import (
	"net/http"

	"github.com/aman-singh7/loyalty-blockchain/domain/coupon"
	"github.com/aman-singh7/loyalty-blockchain/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
)

type Service struct {
	api *Api
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

func (s *Service) CreateCoupon(transactionId int, coupon coupon.Coupon, address common.Address, cost int) error {
	_, err := s.api.ApiTransactor.CreateCoupons(&bind.TransactOpts{}, utils.BigInt(transactionId), address, utils.BigInt(coupon.Count), utils.BigInt(coupon.SuperCoins), utils.BigInt(coupon.Discount), utils.BigInt(coupon.ProductCategory), utils.BigInt(coupon.ThresholdValue), utils.BigInt((coupon.ProductId)), uint8(coupon.Type), utils.BigInt(coupon.ExpiryDate), utils.BigInt(cost))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
	}
	return nil
}

func (s *Service) RedeemTokens(amount int) error {
	_, err := s.api.ApiTransactor.RedeemTokens(&bind.TransactOpts{}, utils.BigInt(amount))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
	}
	return nil
}
