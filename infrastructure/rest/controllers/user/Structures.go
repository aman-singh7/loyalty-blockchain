package user

import (
	"github.com/aman-singh7/loyalty-blockchain/domain/coupon"
)

type User struct {

}

type DiscountRequest struct {
	UserID int `json:"uid"`
	ProductID string `json:"productID"`
	ProductCategory string `json:"productCategory"`
	Coupon coupon.Coupon `json:"coupon"`
	Tokens int `json:"tokens"`
	TransactionID int `json:"transactionID"`
}

type PurchaseProductRequest struct {
	UserID int `json:"uid"`
	ProductID string `json:"productID"`
	ProductCategory string `json:"productCategory"`
	Coupon coupon.Coupon `json:"coupon"`
	Tokens int `json:"tokens"`
	TransactionID int `json:"transactionID"`
}

type PurchaseCouponRequest struct {
	UserID int `json:"uid"`
	Tokens int `json:"tokens"`
	Coupon coupon.Coupon `json:"coupon"`
	TransactionID int `json:"transactionID"`
}

type ReferralRewardRequest struct {
	UserID int `json:"uid"`
	Tokens int `json:"tokens"`
	TransactionID int `json:"transactionID"`
}