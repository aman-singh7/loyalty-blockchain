import { connectWallet } from 'app/services/auth/web3';
import { AccountType } from 'app/types/enums/contractEnums';
import { User } from 'app/types/interfaces/user';
import React, { ReactNode, createContext, useEffect, useState } from 'react';

type userContextType = {
  user?: User;
  isLoading: boolean;
};

const defaultValue: userContextType = {
  user: {
    accountType: AccountType.BUSINESS,
    name: 'Anand',
    platformUid: 'anosos',
    uid: 'kwkdwd',
    walletId: 'jnqwndqwf',
  },
  isLoading: false,
};

export const UserContext = createContext(defaultValue);

const UserContextProvider = (props: { children: ReactNode }) => {
  const [user, setUser] = useState<User | undefined>(defaultValue.user);
  const [account, setAccount] = useState<string>('');
  const [isLoading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    const connect = async () => {
      const account: string = await connectWallet();
      if (account != '') {
        console.log(account);
        setAccount(account);
      } else {
        console.log('invalid metamask account');
      }
    };

    void (async () => {
      await connect();
    })();
  }, []);

  useEffect(() => {
    setLoading(true);
    // fetch user
    setLoading(false);
  }, []);

  const value: userContextType = { user, isLoading };

  return (
    <UserContext.Provider value={value}>{props.children}</UserContext.Provider>
  );
};

export default UserContextProvider;
