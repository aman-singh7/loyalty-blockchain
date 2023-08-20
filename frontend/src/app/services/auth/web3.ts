export const connectWallet: () => Promise<string> = async () => {
  if (!window.ethereum) {
    console.log('MetaMask is not installed');
    return '';
  }

  console.log('MetaMask is installed');

  const accounts: string[] = await window.ethereum.request({
    method: 'eth_accounts',
  });

  if (accounts.length == 0) {
    console.log('No authorized accout found');
    return '';
  }

  const account: string = accounts[0];
  console.log('Found an authorized account', account);
  return account;
};
