package user

import (
	domain "github.com/aman-singh7/loyalty-blockchain/domain/user"
)

func toDomainDiscountRequest(request *DiscountRequest, TransactionID int) *domain.DiscountRequest {
	return &domain.DiscountRequest{
		UserID: request.UserID,
		ProductID: request.ProductID,
		ProductCategory: request.ProductCategory,
		Coupon: request.Coupon,
		Tokens: request.Tokens,
		TransactionID: TransactionID,
	}
} 

func toDomainPurchaseProductRequest(request *PurchaseProductRequest, TransactionID int) *domain.PurchaseProductRequest {
	return &domain.PurchaseProductRequest{
		UserID: request.UserID,
		ProductID: request.ProductID,
		ProductCategory: request.ProductCategory,
		Coupon: request.Coupon,
		Tokens: request.Tokens,
		Price: request.Price,
		TransactionID: TransactionID,
	}
}

func toDomainPurchaseCouponRequest(request *PurchaseCouponRequest, TransactionID int) *domain.PurchaseCouponRequest {
	return &domain.PurchaseCouponRequest{
		UserID: request.UserID,
		Tokens: request.Tokens,
		Coupon: request.Coupon,
		TransactionID: TransactionID,
	}
}

func toDomainReferralRewardRequest(request *ReferralRewardRequest, TransactionID int) *domain.ReferralRewardRequest {
	return &domain.ReferralRewardRequest{
		UserID: request.UserID,
		Tokens: request.Tokens,
		TransactionID: TransactionID,
	}
}