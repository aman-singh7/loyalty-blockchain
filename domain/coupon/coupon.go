package coupon

import "github.com/ethereum/go-ethereum/common"

type CouponType int

const (
	CATEGORY CouponType = iota
	UNIQUE
)

type Coupon struct {
	IssuerBusiness  common.Address `json:"issuerBusiness" validate:"required"`
	SuperCoins      int            `json:"superCoins" validate:"required"`
	Count           int            `json:"count" validate:"required"`
	Discount        int            `json:"discount" validate:"required"`
	ProductCategory int            `json:"productCategory"`
	ThresholdValue  int            `json:"thresholdvalue"`
	ProductId       int            `json:"productId"`
	Type            CouponType     `json:"couponType" validate:"required"`
	ExpiryDate      int            `json:"expiryDate"`
	CouponID        int            `json:"couponId" validate:"required"`
	HoldingID       int            `json:"holdingId"`
}
