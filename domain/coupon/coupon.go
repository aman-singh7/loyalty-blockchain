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
	IssuerBusiness  common.Address `json:"issuerBusiness"`
	SuperCoins      big.Int        `json:"superCoins"`
	Count           big.Int        `json:"count"`
	Discount        big.Int        `json:"discount"`
	ProductCategory big.Int        `json:"productCategory"`
	ThresholdValue  big.Int        `json:"thresholdvalue"`
	ProductId       big.Int        `json:"productId"`
	Type            CouponType     `json:"couponType"`
	ExpiryDate      big.Int        `json:"expiryDate"`
	CouponID        big.Int        `json:"couponId"`
	HoldingID       big.Int        `json:"holdingId"`
}

