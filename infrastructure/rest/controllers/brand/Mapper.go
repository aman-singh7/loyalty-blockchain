package brand

import (
	domain "github.com/aman-singh7/loyalty-blockchain/domain/brand"
	"github.com/ethereum/go-ethereum/common"
)

func toDomainCouponPriceRequest(request *CouponPriceRequest, transactionId int, brandAddress common.Address) *domain.CouponPriceRequest {
	return &domain.CouponPriceRequest{
		BrandAddress:  brandAddress,
		Coupon:        request.Coupon,
		Count:         request.Count,
		Discount:      request.Discount,
		TransactionID: transactionId,
	}
}

func toDomainCreateCouponRequest(request *CreateCouponRequest, transactionId int, brandAddress common.Address) *domain.CreateCouponRequest {
	return &domain.CreateCouponRequest{
		BrandAddress:  brandAddress,
		Coupon:        request.Coupon,
		Count:         request.Count,
		Discount:      request.Discount,
		Tokens:        request.Tokens,
		TransactionID: transactionId,
	}
}

func toDomainRedeemTokensRequest(request *RedeemTokensRequest, transactionId int, brandAddress common.Address) *domain.RedeemTokensRequest {
	return &domain.RedeemTokensRequest{
		BrandAddress:  brandAddress,
		Tokens:        request.Tokens,
		TransactionID: transactionId,
	}
}
