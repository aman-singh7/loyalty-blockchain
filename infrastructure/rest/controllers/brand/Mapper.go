package brand

import (
	domain "github.com/aman-singh7/loyalty-blockchain/domain/brand"
	"github.com/ethereum/go-ethereum/common"
)

func toDomainCouponPriceRequest(request *CouponPriceRequest, brandAddress common.Address) *domain.CouponPriceRequest {
	return &domain.CouponPriceRequest{
		BrandAddress:  brandAddress,
		Coupon:        request.Coupon,
		Count:         request.Count,
		Discount:      request.Discount,
	}
}

func toDomainCreateCouponRequest(request *CreateCouponRequest, brandAddress common.Address) *domain.CreateCouponRequest {
	return &domain.CreateCouponRequest{
		BrandAddress:  brandAddress,
		Coupon:        request.Coupon,
		Count:         request.Count,
		Discount:      request.Discount,
		Tokens:        request.Tokens,
	}
}

func toDomainRedeemTokensRequest(request *RedeemTokensRequest, brandAddress common.Address) *domain.RedeemTokensRequest {
	return &domain.RedeemTokensRequest{
		BrandAddress:  brandAddress,
		Tokens:        request.Tokens,	
	}
}
