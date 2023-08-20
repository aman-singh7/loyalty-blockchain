import TrsansactionInsight from 'app/types/interfaces/transaction_insight';

const randomColor = function () {
  const r = Math.floor(Math.random() * 255);
  const g = Math.floor(Math.random() * 255);
  const b = Math.floor(Math.random() * 255);
  return 'rgb(' + r.toString() + ',' + g.toString() + ',' + b.toString() + ')';
};

export const generateRandomColors = (trsansactions: TrsansactionInsight[]) => {
  for (let i = 0; i < trsansactions.length; i++) {
    trsansactions[i].color = randomColor();
  }
};
