package user

import (
	"math/big"

	"github.com/aman-singh7/loyalty-blockchain/domain/coupon"
)

type User struct{}

type DiscountRequest struct {
	JWT             string        `json:"jwt"`
	ProductID       big.Int       `json:"productID"`
	ProductCategory big.Int       `json:"productCategory"`
	Coupon          coupon.Coupon `json:"coupon"`
	Tokens          big.Int       `json:"tokens"`
}

type PurchaseProductRequest struct {
	JWT             string         `json:"jwt"`
	ProductID       big.Int        `json:"productID"`
	ProductCategory big.Int        `json:"productCategory"`
	Coupon          *coupon.Coupon `json:"coupon"`
	Tokens          big.Int        `json:"tokens"`
	Price           big.Int        `json:"price"`
}

type PurchaseCouponRequest struct {
	JWT    string        `json:"jwt"`
	Tokens big.Int       `json:"tokens"`
	Coupon coupon.Coupon `json:"coupon"`
}

type ReferralRewardRequest struct {
	JWT    string  `json:"jwt"`
	Tokens big.Int `json:"tokens"`
}

type GetAllCouponsRequest struct {
	JWT string `json:"jwt"`
}

type GetBrandCouponsRequest struct {
	JWT     string `json:"jwt"`
	BrandID string `json:"brandID"`
}
