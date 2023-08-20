package user

import (
	"math/big"

	domain "github.com/aman-singh7/loyalty-blockchain/domain/user"
	"github.com/ethereum/go-ethereum/common"
)

func toDomainDiscountRequest(request *DiscountRequest, transactionId big.Int, userAddress common.Address) *domain.DiscountRequest {
	return &domain.DiscountRequest{
		UserAddress:     userAddress,
		ProductID:       request.ProductID,
		ProductCategory: request.ProductCategory,
		Coupon:          request.Coupon,
		Tokens:          request.Tokens,
		TransactionID:   transactionId,
	}
}

func toDomainPurchaseProductRequest(request *PurchaseProductRequest, transactionId big.Int, userAddress common.Address) *domain.PurchaseProductRequest {
	return &domain.PurchaseProductRequest{
		UserAddress:     userAddress,
		ProductID:       request.ProductID,
		ProductCategory: request.ProductCategory,
		Coupon:          request.Coupon,
		Tokens:          request.Tokens,
		Price:           request.Price,
		TransactionID:   transactionId,
	}
}

func toDomainPurchaseCouponRequest(request *PurchaseCouponRequest, transactionId big.Int, userAddress common.Address) *domain.PurchaseCouponRequest {
	return &domain.PurchaseCouponRequest{
		UserAddress:   userAddress,
		Tokens:        request.Tokens,
		Coupon:        request.Coupon,
		TransactionID: transactionId,
	}
}

func toDomainReferralRewardRequest(request *ReferralRewardRequest, transactionId big.Int, userAddress common.Address) *domain.ReferralRewardRequest {
	return &domain.ReferralRewardRequest{
		UserAddress:   userAddress,
		Tokens:        request.Tokens,
		TransactionID: transactionId,
	}
}
