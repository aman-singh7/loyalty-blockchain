package user

import (
	"math/big"

	"github.com/aman-singh7/loyalty-blockchain/domain/coupon"
)

type User struct{}

type DiscountRequest struct {
	UserID          string        `json:"uid"`
	ProductID       big.Int       `json:"productID"`
	ProductCategory big.Int       `json:"productCategory"`
	Coupon          coupon.Coupon `json:"coupon"`
	Tokens          big.Int       `json:"tokens"`
}

type PurchaseProductRequest struct {
	UserID          string        `json:"uid"`
	ProductID       big.Int       `json:"productID"`
	ProductCategory big.Int       `json:"productCategory"`
	Coupon          coupon.Coupon `json:"coupon"`
	Tokens          big.Int       `json:"tokens"`
	Price           big.Int       `json:"price"`
}

type PurchaseCouponRequest struct {
	UserID string        `json:"uid"`
	Tokens big.Int       `json:"tokens"`
	Coupon coupon.Coupon `json:"coupon"`
}

type ReferralRewardRequest struct {
	UserID string  `json:"uid"`
	Tokens big.Int `json:"tokens"`
}
