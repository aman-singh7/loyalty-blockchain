package admin

type Admin struct {

}

type GetALLCouponsRequest struct {
	// add filter of cost etc
}

type GetBrandCouponRequest struct {
	// add filter of cost etc
	BrandID int `json:"uid"`
}