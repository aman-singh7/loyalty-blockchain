package user

import (
	domain "github.com/aman-singh7/loyalty-blockchain/domain/user"
	"github.com/ethereum/go-ethereum/common"
)

func toDomainDiscountRequest(request *DiscountRequest, userAddress common.Address) *domain.DiscountRequest {
	return &domain.DiscountRequest{
		UserAddress:     userAddress,
		ProductID:       request.ProductID,
		ProductCategory: request.ProductCategory,
		Coupon:          request.Coupon,
		Tokens:          request.Tokens,
	}
}

func toDomainPurchaseProductRequest(request *PurchaseProductRequest, userAddress common.Address) *domain.PurchaseProductRequest {
	return &domain.PurchaseProductRequest{
		UserAddress:     userAddress,
		ProductID:       request.ProductID,
		ProductCategory: request.ProductCategory,
		Coupon:          request.Coupon,
		Tokens:          request.Tokens,
		Price:           request.Price,
	}
}

func toDomainPurchaseCouponRequest(request *PurchaseCouponRequest, userAddress common.Address) *domain.PurchaseCouponRequest {
	return &domain.PurchaseCouponRequest{
		UserAddress:   userAddress,
		Tokens:        request.Tokens,
		Coupon:        request.Coupon,
	}
}

func toDomainReferralRewardRequest(request *ReferralRewardRequest, userAddress common.Address) *domain.ReferralRewardRequest {
	return &domain.ReferralRewardRequest{
		UserAddress:   userAddress,
		Tokens:        request.Tokens,
	}
}
