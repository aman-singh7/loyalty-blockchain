import React, { useRef, useState } from 'react';
import { monthsList } from 'app/services/date_service';
import {
  brandCouponList,
  productList,
} from 'app/services/data/dummy_data/dummy';
import plus from 'app/assets/icons/plus.svg';
import superCoin from 'app/assets/icons/super-coin.svg';
import './index.scss';
import Coupons from 'app/types/interfaces/coupons';
import { useDraggable } from 'react-use-draggable-scroll';
import Product from 'app/types/interfaces/product';

const Home: React.FC = () => {
  console.log('business definitely rendered');
  const ref1 =
    useRef<HTMLDivElement>() as React.MutableRefObject<HTMLInputElement>;
  const { events: events1 } = useDraggable(ref1);
  const ref2 =
    useRef<HTMLDivElement>() as React.MutableRefObject<HTMLInputElement>;
  const { events: events2 } = useDraggable(ref2);
  return (
    <div className='business-home'>
      <div className='business-home-token-tracker'>
        <div className='business-home-token-tracker-data'>
          <div className='business-home-token-tracker-data-header'>
            Total tokens
          </div>
          <div className='business-home-token-tracker-data-value'>-1000</div>
        </div>
        <button className='business-home-token-tracker-claim-button'>
          Claim Now
        </button>
      </div>
      <div className='business-home-business-offer-container'>
        <div className='business-home-business-offer-container-header'>
          Active Coupons
        </div>
        <div
          className='business-home-business-offer-entity-container'
          {...events1}
          ref={ref1}
        >
          {brandCouponList.map((coupon: Coupons, couponIndex: number) => (
            <div
              key={couponIndex}
              className='business-home-business-offer-entity'
            >
              <div className='business-home-business-offer-entity-image'>
                <div className='business-home-business-offer-entity-cost-container'>
                  <img
                    src={superCoin}
                    alt='super coin'
                    className='business-home-business-offer-entity-cost-icon'
                  />
                  <div className='business-home-business-offer-entity-cost'>
                    {coupon.superCoins.toString()}
                  </div>
                </div>
              </div>
              <div className='business-home-business-offer-entity-text'>
                Get {coupon.discount.toString()}% discount on{' '}
                {coupon.productText}
              </div>
            </div>
          ))}
        </div>
        <div className='business-home-business-offer-create-button'>
          <div className='business-home-business-offer-create-button-text'>
            Create new coupon
          </div>
          <img
            src={plus}
            alt='add coupon'
            className='business-home-business-offer-create-button-icon'
          />
        </div>
      </div>
      {/* <div className='business-home-business-product-container'>
        <div className='business-home-business-product-container-header'>
          Products
        </div>
        <div
          className='business-home-business-product-entity-container'
          {...events2}
          ref={ref2}
        >
          {productList.map((product: Product, productIndex: number) => (
            <div
              key={productIndex}
              className='business-home-business-product-entity'
            >
              <div className='business-home-business-product-entity-image'>
                <div className='business-home-business-product-entity-cost-container'>
                  <img
                    src={superCoin}
                    alt='super coin'
                    className='business-home-business-product-entity-cost-icon'
                  />
                  <div className='business-home-business-product-entity-cost'>
                    {product.cost.toString()}
                  </div>
                </div>
              </div>
              <div className='business-home-business-product-entity-text'>
                {product.productText}
              </div>
            </div>
          ))}
        </div>
        <div className='business-home-business-product-create-button'>
          <div className='business-home-business-product-create-button-text'>
            Create new coupon
          </div>
          <img
            src={plus}
            alt='add coupon'
            className='business-home-business-product-create-button-icon'
          />
        </div>
      </div> */}
      <div className='business-home-business-product-container'>
        <div className='business-home-business-product-container-header'>
          Products
        </div>
        <div
          className='business-home-business-product-entity-container'
          {...events2}
          ref={ref2}
        >
          {productList.map((product: Product, productIndex: number) => (
            <div
              key={productIndex}
              className='business-home-business-offer-entity'
            >
              <div className='business-home-business-offer-entity-image'>
                <div className='business-home-business-offer-entity-cost-container'>
                  <div className='business-home-business-offer-entity-cost-icon'>
                    Rs.
                  </div>
                  <div className='business-home-business-offer-entity-cost'>
                    {product.cost.toString()}
                  </div>
                </div>
              </div>
              <div className='business-home-business-offer-entity-text'>
                {product.productText} is issued at cost{' '}
                {product.cost.toString()}
              </div>
            </div>
          ))}
        </div>
        <div className='business-home-business-offer-create-button'>
          <div className='business-home-business-offer-create-button-text'>
            Create new Product
          </div>
          <img
            src={plus}
            alt='add coupon'
            className='business-home-business-offer-create-button-icon'
          />
        </div>
      </div>
    </div>
  );
};

export default Home;
