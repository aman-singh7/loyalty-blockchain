package user

import (
	"github.com/aman-singh7/loyalty-blockchain/domain/coupon"
	"github.com/ethereum/go-ethereum/common"
)

type User struct {
	UID           string `json:"uid"`
	PlatformUID   string `json:"platformUid"`
	Name          string `json:"name"`
	WalletAddress string `json:"walletAddress"`
	AccountType   uint8  `json:"accountType"`
}

type DiscountRequest struct {
	UserAddress     common.Address `json:"userAddress"`
	ProductID       int            `json:"productID"`
	ProductCategory int            `json:"productCategory"`
	Coupon          coupon.Coupon  `json:"coupon"`
	Tokens          int            `json:"tokens"`
	TransactionID   int            `json:"transactionID"`
}

type PurchaseProductRequest struct {
	UserAddress     common.Address `json:"userAddress"`
	ProductID       int            `json:"productID"`
	ProductCategory int            `json:"productCategory"`
	Coupon          coupon.Coupon  `json:"coupon"`
	Tokens          int            `json:"tokens"`
	Price           int            `json:"price"`
	TransactionID   int            `json:"transactionID"`
}

type PurchaseCouponRequest struct {
	UserAddress   common.Address `json:"userAddress"`
	Tokens        int            `json:"tokens"`
	Count         int            `json:"count"`
	Coupon        coupon.Coupon  `json:"coupon"`
	TransactionID int            `json:"transactionID"`
}

type ReferralRewardRequest struct {
	UserAddress   common.Address `json:"userAddress"`
	Tokens        int            `json:"tokens"`
	TransactionID int            `json:"transactionID"`
}
