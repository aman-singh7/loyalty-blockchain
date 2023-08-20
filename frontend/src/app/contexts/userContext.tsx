import { login } from 'app/services/api/login';
import { auth } from 'app/services/auth/firebase';
import { connectWallet } from 'app/services/auth/web3';
import { AccountType } from 'app/types/enums/contractEnums';
import { User } from 'app/types/interfaces/user';
import { UserCredential } from 'firebase/auth';
import React, { ReactNode, createContext, useEffect, useState } from 'react';
import { useSignInWithGoogle } from 'react-firebase-hooks/auth';

type userContextType = {
  user?: User;
  error?: string;
  signIn?: VoidFunction;
  isLoading: boolean;
};

const defaultValue: userContextType = {
  isLoading: false,
};

export const UserContext = createContext(defaultValue);

const UserContextProvider = (props: { children: ReactNode }) => {
  const [user, setUser] = useState<User | undefined>(defaultValue.user);
  const [signInWithGoogle, googleUser, loading, googleError] =
    useSignInWithGoogle(auth);
  const [account, setAccount] = useState<string>('');
  const [isLoading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | undefined>('');

  const connect = async () => {
    const account: string = await connectWallet();
    if (account != '') {
      console.log(account);
      setAccount(account);
    } else {
      console.log('invalid metamask account');
    }
  };

  const signIn = async () => {
    if (!user) {
      let creds;
      try {
        creds = await signInWithGoogle();
      } catch (e) {
        console.log(e);
        setError(googleError?.message);
      }

      const uid = creds?.user.uid;
      const token = await creds?.user.getIdToken();

      let res;
      try {
        res = await login(uid!, token!);
      } catch (e) {
        console.log(e);
        setError('Some error occured');
      }
    }
  };

  useEffect(() => {
    setLoading(true);
    // fetch user
    setLoading(false);
  }, []);

  const value: userContextType = { user, signIn, isLoading };

  return (
    <UserContext.Provider value={value}>{props.children}</UserContext.Provider>
  );
};

export default UserContextProvider;
