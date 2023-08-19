package coupon

type CouponType int 
const (
	CATEGORY CouponType = iota
	UNIQUE
)

type Coupon struct {
	IssuerBusiness string `json:"issuerBusiness" validate:"required"`
	SuperCoins int `json:"superCoins" validate:"required"`
	Count int `json:"count" validate:"required"`
	Discount int `json:"discount" validate:"required"`
	ProductCategory string `json:"productCategory"`
    ThresholdValue int `json:"thresholdvalue"`
	ProductId string `json:"productId"`
    Type CouponType `json:"couponType" validate:"required"`
    Expiry int `json:"expiry"`
	CouponID int `json:"couponId" validate:"required"`
}