import React from 'react';
import superCoin from 'app/assets/icons/super-coin.svg';
import backIcon from 'app/assets/icons/back-icon.svg';
import TrsansactionHistory from 'app/types/interfaces/transaction_history';
import TrsansactionInsight from 'app/types/interfaces/transaction_insight';
import transactions, {
  superCoinCount,
  transactionHistory,
} from 'app/services/data/dummy_data/dummy';
import { ChartData, ChartOptions } from 'chart.js';
import 'chart.js/auto';
import { Pie } from 'react-chartjs-2';
import './index.scss';
import { formatDate } from 'app/services/date_service';

const History: React.FC = () => {
  const businessList = transactions.map(
    (transaction: TrsansactionInsight) => transaction.business
  );
  const countList = transactions.map((transaction: TrsansactionInsight) =>
    Number(transaction.transactionCount)
  );
  const colorList = transactions.map(
    (transaction: TrsansactionInsight) => transaction.color
  );
  const chartData: ChartData<'pie'> = {
    labels: businessList,
    datasets: [
      {
        label: 'No of Transactions',
        data: countList,
        backgroundColor: colorList,
        borderColor: colorList,
        borderWidth: 0.5,
      },
    ],
  };

  const chartOptions: ChartOptions<'pie'> = {
    plugins: {
      legend: {
        display: false,
      },
    },
  };

  return (
    <div className='client-history'>
      <div className='client-history-header'>
        <img className='client-history-header-icon' src={backIcon} alt='back' />
        <div className='client-history-header-text'>Token Activity</div>
      </div>
      <div className='client-history-token-container'>
        <div>Tokens Available</div>
        <img
          src={superCoin}
          alt='super coin'
          className='client-history-token-container-supercoin-image'
        />
        <div className='client-history-token-container-supercoin'>
          {superCoinCount.toString()}
        </div>
      </div>
      <div className='client-history-insights-container'>
        <div className='client-history-insights-data'>
          <div className='client-history-insights-data-header'>Insights</div>
          {transactions.map(
            ({ business, transactionCount, color }, index: number) => (
              <div
                key={index}
                className='client-history-insights-data-entry-container'
              >
                <div
                  className='client-history-insights-data-entry-color'
                  style={{ backgroundColor: color }}
                />
                <div className='client-history-data-entry'>
                  {`${business} : ${transactionCount}`}
                </div>
              </div>
            )
          )}
        </div>
        <div className='client-history-insights-chart-container'>
          <Pie
            className='client-history-insights-chart'
            options={chartOptions}
            data={chartData}
          ></Pie>
        </div>
      </div>
      <div className='client-history-transaction-container'>
        <div className='client-history-transaction-header'>
          Transaction History
        </div>
        {transactionHistory.map(
          (
            {
              business,
              cost,
              productText,
              dateTime,
              discount,
            }: TrsansactionHistory,
            index: number
          ) => (
            <div key={index} className='client-history-transaction-entry'>
              <div className='client-history-transaction-entry-photo'> </div>
              <div className='client-history-transaction-entry-data'>
                <div className='client-history-transactiycron-entry-data-product'>
                  {`Get ${discount.toString()}% off on ${business}'s
                  ${productText}`}
                </div>
                <div className='client-history-transaction-entry-data-cost'>
                  <img
                    src={superCoin}
                    className='client-history-transaction-entry-data-cost-icon'
                  />
                  <div className='client-history-transaction-entry-data-cost-value'>
                    {cost.toString()}
                  </div>
                </div>
              </div>
              <div className='client-history-transaction-entry-date'>
                {formatDate(dateTime)}
              </div>
            </div>
          )
        )}
      </div>
    </div>
  );
};

export default History;
