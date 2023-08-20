import React from 'react';
import { AppBarOptions, AppBarProductOptions } from 'app/types/types/options';
import search from 'app/assets/icons/search-icon.svg';
import expand from 'app/assets/icons/expand-icon.svg';
import './index.scss';

const options: AppBarOptions[] = [
  { name: 'My Account' },
  { name: 'Become a Seller' },
  { name: 'More' },
  { name: 'Cart' },
];

const productOptions: AppBarProductOptions[] = [
  { name: 'Electronics' },
  { name: 'TV & Appliances' },
  { name: 'Men' },
  { name: 'Women' },
  { name: 'Baby & Kids' },
  { name: 'Home & Furniture' },
  { name: 'Sports, Books & More' },
  { name: 'Flights' },
  { name: 'Offer Zone' },
];

const AppBar: React.FC = () => (
  <div className='app-bar'>
    <div className='app-bar-header'>
      <div className='app-bar-header-search-bar-wrapper'>
        <input
          className='app-bar-header-search-bar'
          type='text'
          placeholder='Search for products, brands and more'
        />
        <img
          className='app-bar-header-search-bar-icon'
          src={search}
          alt='search'
        />
      </div>
      {options.map((option: AppBarOptions, index: number) => (
        <div className='app-bar-header-options' key={index}>
          {option.name}
        </div>
      ))}
    </div>
    <div className='app-bar-drawer'>
      {productOptions.map(
        (productOption: AppBarProductOptions, index: number) => (
          <div key={index} className='app-bar-drawer-option'>
            <span className='app-bar-drawer-option-text'>
              {productOption.name}
            </span>
            <img className='app-bar-drawer-option-icon' src={expand} />
          </div>
        )
      )}
    </div>
  </div>
);

export default AppBar;
