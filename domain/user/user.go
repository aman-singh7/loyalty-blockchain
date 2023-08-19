package user

import "github.com/aman-singh7/loyalty-blockchain/domain/coupon"

type DomainDiscountRequest struct {
	UserID          int           `json:"uid"`
	ProductID       string        `json:"productID"`
	ProductCategory string        `json:"productCategory"`
	Coupon          coupon.Coupon `json:"coupon"`
	Tokens          int           `json:"tokens"`
	TransactionID   int           `json:"transactionID"`
}

type DomainPurchaseProductRequest struct {
	UserID          int           `json:"uid"`
	ProductID       string        `json:"productID"`
	ProductCategory string        `json:"productCategory"`
	Coupon          coupon.Coupon `json:"coupon"`
	Tokens          int           `json:"tokens"`
	TransactionID   int           `json:"transactionID"`
	Price           int           `json:"price"`
}

type DomainPurchaseCouponRequest struct {
	UserID        int           `json:"uid"`
	Tokens        int           `json:"tokens"`
	Coupon        coupon.Coupon `json:"coupon"`
	TransactionID int           `json:"transactionID"`
}

type DomainReferralRewardRequest struct {
	UserID        int `json:"uid"`
	Tokens        int `json:"tokens"`
	TransactionID int `json:"transactionID"`
}