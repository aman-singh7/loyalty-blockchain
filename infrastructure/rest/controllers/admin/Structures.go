package admin

type GetALLCouponsRequest struct {
	// add filter of cost etc
}

type GetBrandCouponRequest struct {
	// add filter of cost etc
	BrandID string `json:"uid"`
}

type RegisterMemberRequest struct {
	Addr string `json:"address" validate:"required"`
	Type uint8  `json:"type" validate:"required"`
}
