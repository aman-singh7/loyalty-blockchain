package api

import (
	"context"
	"fmt"
	"math/big"
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
	tx   *utils.TransactionGen
}

func NewService(api *Api, opts *utils.TransactOpts, tx *utils.TransactionGen) *Service {
	return &Service{
		api:  api,
		opts: opts,
		tx:   tx,
	}
}

func (s *Service) FetchAccountBalance(address common.Address, from common.Address) (int, error) {
	opts := bind.CallOpts{
		Pending: false,
		From:    from,
		Context: context.Background(),
	}
	transactionId := s.tx.GetTransactionId()
	balance, err := s.api.GetAccountBalance(&opts, transactionId, address)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"message": "fetch balance failed",
		})
	}

	return utils.ToInt(balance), nil
}

func (s *Service) PurchaseProductWithCoupon(coupon coupon.Coupon, address common.Address) error {
	PRIVATE_KEY := viper.GetString("PrivateKey")
	opts := s.opts.GetOpts(PRIVATE_KEY)
	transactionId := s.tx.GetTransactionId()
	_, err := s.api.ApiTransactor.ConsumerCouponPayment(opts, transactionId, &coupon.HoldingID, address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
	}
	return nil
}

func (s *Service) PurchaseProductWithToken(tokens big.Int) error {
	PRIVATE_KEY := viper.GetString("PrivateKey")
	opts := s.opts.GetOpts(PRIVATE_KEY)
	transactionId := s.tx.GetTransactionId()
	_, err := s.api.ApiTransactor.ConsumerTokenPayment(opts, transactionId, &tokens)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
	}
	return nil
}

func (s *Service) PurchaseCoupon(coupon coupon.Coupon, count big.Int, address common.Address) error {
	PRIVATE_KEY := viper.GetString("PrivateKey")
	opts := s.opts.GetOpts(PRIVATE_KEY)
	transactionId := s.tx.GetTransactionId()
	_, err := s.api.ApiTransactor.PurchaseCoupon(opts, transactionId, address, &coupon.CouponID, &count)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
	}
	return nil
}

func (s *Service) RewardToken(userAddress common.Address, amount big.Int) error {
	fmt.Println("Rewarding Tokens")
	PRIVATE_KEY := viper.GetString("PrivateKey")
	opts := s.opts.GetOpts(PRIVATE_KEY)
	transactionId := s.tx.GetTransactionId()
	_, err := s.api.ApiTransactor.Pay(opts, transactionId, userAddress, &amount)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
	}
	return nil
}

func (s *Service) CreateCoupon(coupon coupon.Coupon, cost big.Int) error {
	PRIVATE_KEY := viper.GetString("PrivateKey")
	opts := s.opts.GetOpts(PRIVATE_KEY)
	transactionId := s.tx.GetTransactionId()
	_, err := s.api.ApiTransactor.CreateCoupons(opts, transactionId, coupon.IssuerBusiness, &coupon.Count, &coupon.SuperCoins, &coupon.Discount, &coupon.ProductCategory, &coupon.ThresholdValue, &coupon.ProductId, uint8(coupon.Type), &coupon.ExpiryDate, &cost)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
	}
	return nil
}

func (s *Service) RedeemTokens(businessAddress common.Address, amount big.Int) error {
	// TODO: implementation of redeem tokens
	PRIVATE_KEY := viper.GetString("PrivateKey")
	opts := s.opts.GetOpts(PRIVATE_KEY)
	_, err := s.api.ApiTransactor.RedeemTokens(opts, &amount)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Transaction Failed"})
	}
	return nil
}

func (s *Service) RegisterMember(memberAddr common.Address, accountType uint8) error {
	PRIVATE_KEY := viper.GetString("PrivateKey")
	opts := s.opts.GetOpts(PRIVATE_KEY)
	// 1 -> Customer 2 -> Brand
	_, err := s.api.RegisterMember(opts, utils.BigInt(2), memberAddr, accountType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"message": "register member failed",
		})
	}
	return nil
}

func (s *Service) GetAllCoupons() ([]coupon.Coupon, error) {
	opts := &bind.CallOpts{
		Pending: false,
		From:    common.HexToAddress(viper.GetString("Address")),
		Context: context.Background(),
	}
	transactionId := s.tx.GetTransactionId()
	coupons, err := s.api.ApiCaller.GetAllCoupons(opts, transactionId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"message": "fetch coupons failed",
		})
	}
	results := make([]coupon.Coupon, len(coupons))
	for i := 0; i < len(coupons); i++ {
		results[i] = *fromContractCoupontoCoupon(coupons[i])
	}
	return results, nil
}

func (s *Service) GetBrandCoupons(brandAddress common.Address) ([]coupon.Coupon, error) {
	opts := &bind.CallOpts{
		Pending: false,
		From:    common.HexToAddress(viper.GetString("Address")),
		Context: context.Background(),
	}
	transactionId := s.tx.GetTransactionId()
	coupons, err := s.api.ApiCaller.GetBrandCoupons(opts, transactionId, brandAddress)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"message": "fetch coupons failed",
		})
	}
	results := make([]coupon.Coupon, len(coupons))
	for i := 0; i < len(coupons); i++ {
		results[i] = *fromContractCoupontoCoupon(coupons[i])
	}
	return results, nil
}
