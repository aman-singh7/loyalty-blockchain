package coupon

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type CouponType int

const (
	CATEGORY CouponType = iota
	UNIQUE
)

type Coupon struct {
	IssuerBusiness  common.Address `json:"issuerBusiness" validate:"required"`
	SuperCoins      big.Int        `json:"superCoins" validate:"required"`
	Count           big.Int        `json:"count" validate:"required"`
	Discount        big.Int        `json:"discount" validate:"required"`
	ProductCategory big.Int        `json:"productCategory"`
	ThresholdValue  big.Int        `json:"thresholdvalue"`
	ProductId       big.Int        `json:"productId"`
	Type            CouponType     `json:"couponType" validate:"required"`
	ExpiryDate      big.Int        `json:"expiryDate"`
	CouponID        big.Int        `json:"couponId" validate:"required"`
	HoldingID       big.Int        `json:"holdingId"`
}
