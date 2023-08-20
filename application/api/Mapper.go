package api

import (
	domain "github.com/aman-singh7/loyalty-blockchain/domain/coupon"
)

func fromContractCoupontoCoupon(coupon LoyaltyProgrammeCoupon) *domain.Coupon {
	return &domain.Coupon{
		IssuerBusiness:  coupon.IssuerBusiness,
		Count:           *coupon.Count,
		Discount:        *coupon.Discount,
		ProductCategory: *coupon.ProductCategory,
		ThresholdValue:  *coupon.ThresholdValue,
		ProductId:       *coupon.ProductId,
		Type:            domain.CouponType(coupon.CouponType),
		ExpiryDate:      *coupon.ExpiryTime,
		CouponID:        *coupon.CouponId,
	}
}
