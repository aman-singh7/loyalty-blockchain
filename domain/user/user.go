package user

import (
	"math/big"

	"github.com/aman-singh7/loyalty-blockchain/domain/coupon"
	"github.com/ethereum/go-ethereum/common"
)

type User struct {
	UID           string `json:"uid"`
	PlatformUID   string `json:"platformUid"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	WalletAddress string `json:"walletAddress"`
	AccountType   uint8  `json:"accountType"`
}

type DiscountRequest struct {
	UserAddress     common.Address `json:"userAddress"`
	ProductID       big.Int        `json:"productID"`
	ProductCategory big.Int        `json:"productCategory"`
	Coupon          coupon.Coupon  `json:"coupon"`
	Tokens          big.Int        `json:"tokens"`
	TransactionID   big.Int        `json:"transactionID"`
}

type PurchaseProductRequest struct {
	UserAddress     common.Address `json:"userAddress"`
	ProductID       big.Int        `json:"productID"`
	ProductCategory big.Int        `json:"productCategory"`
	Coupon          *coupon.Coupon `json:"coupon"`
	Tokens          big.Int        `json:"tokens"`
	Price           big.Int        `json:"price"`
	TransactionID   big.Int        `json:"transactionID"`
}

type PurchaseCouponRequest struct {
	UserAddress   common.Address `json:"userAddress"`
	Tokens        big.Int        `json:"tokens"`
	Count         big.Int        `json:"count"`
	Coupon        coupon.Coupon  `json:"coupon"`
	TransactionID big.Int        `json:"transactionID"`
}

type ReferralRewardRequest struct {
	UserAddress   common.Address `json:"userAddress"`
	Tokens        big.Int        `json:"tokens"`
	TransactionID big.Int        `json:"transactionID"`
}
