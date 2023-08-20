import { CouponType } from 'app/types/enums/contractEnums';
import Coupons from 'app/types/interfaces/coupons';
import React, { useRef } from 'react';
import { useDraggable } from 'react-use-draggable-scroll';
import superCoin from 'app/assets/icons/super-coin.svg';
import nextIcon from 'app/assets/icons/next-icon.svg';
import './index.scss';

const couponList: Map<string, Coupons[]> = new Map([
  [
    'amazon',
    [
      {
        couponId: BigInt(0),
        issuerBusiness: 'bussiness',
        superCoins: BigInt(0),
        count: BigInt(0),
        discount: BigInt(0),
        productCategory: BigInt(0),
        thresholdValue: BigInt(0),
        productId: BigInt(0),
        couponType: CouponType.UNIQUE,
        expiryTime: BigInt(0),
        productText: 'fridges',
        active: true,
      },
      {
        couponId: BigInt(0),
        issuerBusiness: 'bussiness',
        superCoins: BigInt(0),
        count: BigInt(0),
        discount: BigInt(0),
        productCategory: BigInt(0),
        thresholdValue: BigInt(0),
        productId: BigInt(0),
        couponType: CouponType.UNIQUE,
        expiryTime: BigInt(0),
        productText: 'ABCD fridges',
        active: true,
      },
      {
        couponId: BigInt(0),
        issuerBusiness: 'bussiness',
        superCoins: BigInt(0),
        count: BigInt(0),
        discount: BigInt(0),
        productCategory: BigInt(0),
        thresholdValue: BigInt(0),
        productId: BigInt(0),
        couponType: CouponType.UNIQUE,
        expiryTime: BigInt(0),
        productText: 'ABCD fridges',
        active: true,
      },
      {
        couponId: BigInt(0),
        issuerBusiness: 'bussiness',
        superCoins: BigInt(0),
        count: BigInt(0),
        discount: BigInt(0),
        productCategory: BigInt(0),
        thresholdValue: BigInt(0),
        productId: BigInt(0),
        couponType: CouponType.UNIQUE,
        expiryTime: BigInt(0),
        productText: 'ABCD fridges',
        active: true,
      },
      {
        couponId: BigInt(0),
        issuerBusiness: 'bussiness',
        superCoins: BigInt(100),
        count: BigInt(0),
        discount: BigInt(0),
        productCategory: BigInt(0),
        thresholdValue: BigInt(0),
        productId: BigInt(0),
        couponType: CouponType.UNIQUE,
        expiryTime: BigInt(0),
        productText: 'ABCD fridges',
        active: true,
      },
    ],
  ],
  [
    'netflix',
    [
      {
        couponId: BigInt(0),
        issuerBusiness: 'bussiness',
        superCoins: BigInt(0),
        count: BigInt(0),
        discount: BigInt(0),
        productCategory: BigInt(0),
        thresholdValue: BigInt(0),
        productId: BigInt(0),
        couponType: CouponType.UNIQUE,
        expiryTime: BigInt(0),
        productText: 'fridges',
        active: true,
      },
      {
        couponId: BigInt(0),
        issuerBusiness: 'bussiness',
        superCoins: BigInt(0),
        count: BigInt(0),
        discount: BigInt(0),
        productCategory: BigInt(0),
        thresholdValue: BigInt(0),
        productId: BigInt(0),
        couponType: CouponType.UNIQUE,
        expiryTime: BigInt(0),
        productText: 'ABCD fridges',
        active: true,
      },
      {
        couponId: BigInt(0),
        issuerBusiness: 'bussiness',
        superCoins: BigInt(100),
        count: BigInt(0),
        discount: BigInt(0),
        productCategory: BigInt(0),
        thresholdValue: BigInt(0),
        productId: BigInt(0),
        couponType: CouponType.UNIQUE,
        expiryTime: BigInt(0),
        productText: 'ABCD fridges',
        active: true,
      },
      {
        couponId: BigInt(0),
        issuerBusiness: 'bussiness',
        superCoins: BigInt(100),
        count: BigInt(0),
        discount: BigInt(0),
        productCategory: BigInt(0),
        thresholdValue: BigInt(0),
        productId: BigInt(0),
        couponType: CouponType.UNIQUE,
        expiryTime: BigInt(0),
        productText: 'ABCD fridges',
        active: true,
      },
      {
        couponId: BigInt(0),
        issuerBusiness: 'bussiness',
        superCoins: BigInt(0),
        count: BigInt(0),
        discount: BigInt(0),
        productCategory: BigInt(0),
        thresholdValue: BigInt(0),
        productId: BigInt(0),
        couponType: CouponType.UNIQUE,
        expiryTime: BigInt(0),
        productText: 'ABCD fridges',
        active: true,
      },
    ],
  ],
]);

const count = BigInt(496);

const Home: React.FC = () => {
  const ref =
    useRef<HTMLDivElement>() as React.MutableRefObject<HTMLInputElement>;
  const { events } = useDraggable(ref);

  return (
    <div className='client-home'>
      <div className='client-home-header'>Claim Exclusive Rewards...</div>
      <div className='client-home-token-container'>
        <div>Tokens Available</div>
        <img
          src={superCoin}
          alt='super coin'
          className='client-home-token-container-supercoin-image'
        />
        <div className='client-home-token-container-supercoin'>
          {count.toString()}
        </div>
      </div>
      <button className='client-home-token-activity'>
        <div className='client-home-token-activity-text'>
          See all token activity
        </div>
        <img
          src={nextIcon}
          alt='token activity'
          className='client-home-token-activity-button'
        />
      </button>
      <div className='client-home-offers-headers'>Available Offers</div>
      <div className='client-home-offers-container'>
        {Array.from(couponList.entries()).map(
          ([name, coupons], businessIndex: number) => (
            <div
              key={businessIndex}
              className='client-home-business-offer-container'
            >
              <div className='client-home-business-offer-information'>
                <div className='client-home-business-offer-information-announcement'>
                  Best Offers from {name}
                </div>
                <button className='client-home-business-offer-information-view-all'>
                  View All
                </button>
              </div>
              <div
                className='client-home-business-offer-entity-container'
                {...events}
                ref={ref}
              >
                {coupons.map((coupon: Coupons, couponIndex) => (
                  <div
                    key={couponIndex}
                    className='client-home-business-offer-entity'
                  >
                    <div className='client-home-business-offer-entity-image'>
                      <div className='client-home-business-offer-entity-cost-container'>
                        <img
                          src={superCoin}
                          alt='super coin'
                          className='client-home-business-offer-entity-cost-icon'
                        />
                        <div className='client-home-business-offer-entity-cost'>
                          {coupon.superCoins.toString()}
                        </div>
                      </div>
                    </div>
                    <div className='client-home-business-offer-entity-text'>
                      Get {coupon.discount.toString()}% discount on{' '}
                      {coupon.productText}
                    </div>
                  </div>
                ))}
              </div>
            </div>
          )
        )}
      </div>
    </div>
  );
};

export default Home;
