/*
struct Coupon {
    uint couponId; // will be indexed onto the vector
    address issuerBusiness;
    uint superCoins; // value of superCoin in each coupon
    uint count;
    uint discount; // multipiled by a factor of `decimalFactor`
    uint productCategory;
    uint thresholdValue;
    uint productId;
    CouponType couponType;
    uint expiryTime;
    bool active;
}
*/

import { CouponType } from 'app/types/enums/contractEnums';

interface Coupons {
  couponId: bigint;
  issuerBusiness: string;
  superCoins: bigint;
  count: bigint;
  discount: bigint;
  productCategory: bigint;
  thresholdValue: bigint;
  productId: bigint;
  couponType: CouponType;
  expiryTime: bigint;
  productText: string;
  active: boolean;
}

export default Coupons;
