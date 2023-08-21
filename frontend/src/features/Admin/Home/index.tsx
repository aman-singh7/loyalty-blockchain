import React from 'react';
import superCoin from 'app/assets/icons/super-coin.svg';
import { redeemList, superCoinCount } from 'app/services/data/dummy_data/dummy';
import './index.scss';
import RedeemRequest from 'app/types/interfaces/request';

const Home: React.FC = () => (
  <div className='owner-home'>
    <div className='owner-home-token-container'>
      <div className='owner-home-token-container-active'>
        <div>Active Tokens</div>
        <img
          src={superCoin}
          alt='super coin'
          className='owner-home-token-container-supercoin-image'
        />
        <div className='owner-home-token-container-supercoin'>
          {superCoinCount.toString()}
        </div>
      </div>
      <div className='owner-home-token-container-account'>
        <div>Tokens in Account</div>
        <img
          src={superCoin}
          alt='super coin'
          className='owner-home-token-container-supercoin-image'
        />
        <div className='owner-home-token-container-supercoin'>
          {superCoinCount.toString()}
        </div>
      </div>
    </div>
    <div className='owner-home-create-token-container'>
      <div className='owner-home-create-token-header'>Create Tokens</div>
      <div className='owner-home-create-token-form'>
        <div className='owner-home-create-token-form-promt'>
          Number of tokens
        </div>
        <input type='number' className='owner-home-create-token-form-input' />
        <button className='owner-home-create-token-form-button'>
          Create Token
        </button>
      </div>
    </div>
    <div className='owner-home-approval-header'>Businesses</div>
    <div className='owner-home-approval-container'>
      <div className='owner-home-approval-columns-name'>Brand Name</div>
      <div className='owner-home-approval-columns-tokens'>Tokens</div>
      <div className='owner-home-approval-columns-claimed'>Claimed</div>
      {redeemList.map((request: RedeemRequest, index: number) => (
        <>
          <div className='owner-home-approval-columns-entry-name'>
            {request.businessName}
          </div>
          <div className='owner-home-approval-columns-entry-tokens'>
            {request.tokens.toString()}
          </div>
          <div className='owner-home-approval-columns-entry-claimed'>
            {request.claimed.toString()}
          </div>
          <button className='owner-home-approval-columns-entry-approve-button'>
            Approve
          </button>
        </>
      ))}
    </div>
  </div>
);

export default Home;
