package brand

import (
	domain "github.com/aman-singh7/loyalty-blockchain/domain/brand"
)

func toDomainCouponPriceRequest(request *CouponPriceRequest, TransactionID int) *domain.CouponPriceRequest {
	return &domain.CouponPriceRequest{
		BrandID:       request.BrandID,
		Coupon:        request.Coupon,
		Count:         request.Count,
		Discount:      request.Discount,
		TransactionID: TransactionID,
	}
}

func toDomainCreateCouponRequest(request *CreateCouponRequest, TransactionID int) *domain.CreateCouponRequest {
	return &domain.CreateCouponRequest{
		BrandID:       request.BrandID,
		Coupon:        request.Coupon,
		Count:         request.Count,
		Discount:      request.Discount,
		Tokens:        request.Tokens,
		TransactionID: TransactionID,
	}
}

func toDomainRedeemTokensRequest(request *RedeemTokensRequest, TransactionID int) *domain.RedeemTokensRequest {
	return &domain.RedeemTokensRequest{
		BrandID:       request.BrandID,
		Tokens:        request.Tokens,
		TransactionID: TransactionID,
	}
}
