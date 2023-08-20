import React, { useContext } from 'react';
import ConsumerHome from 'features/Client/Home';
import BusinessHome from 'features/Business/Home';
import OwnerHome from 'features/Admin/Home';

import './index.scss';
import { UserContext } from './contexts/userContext';
import { AccountType } from './types/enums/contractEnums';
import AppBar from './components/AppBar';
import Login from 'features/Login';

function App() {
  const { user } = useContext(UserContext);

  const getHomePage = (accountType: AccountType) => {
    if (accountType === AccountType.CONSUMER) {
      return <ConsumerHome />;
    }

    if (accountType === AccountType.BUSINESS) {
      return <BusinessHome />;
    }

    if (accountType === AccountType.OWNER) {
      return <OwnerHome />;
    }

    return <>Test</>;
  };

  if (user && user.walletId.length != 0) {
    return (
      <>
        <AppBar />
        {getHomePage(user.accountType)}
      </>
    );
  }

  return <Login />;
}

export default App;
