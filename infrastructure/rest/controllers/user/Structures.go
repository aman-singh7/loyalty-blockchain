package user

import (
	"github.com/aman-singh7/loyalty-blockchain/domain/coupon"
)

type User struct{}

type DiscountRequest struct {
	UserID          string        `json:"uid"`
	ProductID       int           `json:"productID"`
	ProductCategory int           `json:"productCategory"`
	Coupon          coupon.Coupon `json:"coupon"`
	Tokens          int           `json:"tokens"`
}

type PurchaseProductRequest struct {
	UserID          string        `json:"uid"`
	ProductID       int           `json:"productID"`
	ProductCategory int           `json:"productCategory"`
	Coupon          coupon.Coupon `json:"coupon"`
	Tokens          int           `json:"tokens"`
	Price           int           `json:"price"`
}

type PurchaseCouponRequest struct {
	UserID string        `json:"uid"`
	Tokens int           `json:"tokens"`
	Coupon coupon.Coupon `json:"coupon"`
}

type ReferralRewardRequest struct {
	UserID string `json:"uid"`
	Tokens int    `json:"tokens"`
}
