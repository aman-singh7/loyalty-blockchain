import React, { useContext } from 'react';
import AppBar from 'app/components/AppBar';
import Login from 'features/Login';
import ConsumerHome from 'features/Client/Home';
import BusinessHome from 'features/Business/Home';

import './index.scss';
import { UserContext } from './contexts/userContext';
import { AccountType } from './types/enums/contractEnums';

function App() {
  const { user } = useContext(UserContext);

  const getHomePage = (accountType: AccountType) => {
    if (accountType === AccountType.CONSUMER) {
      console.log('consumer home rendered');
      return <ConsumerHome />;
    }

    if (accountType === AccountType.BUSINESS) {
      console.log('business home rendered');
      return <BusinessHome />;
    }

    return <>Test</>;
  };

  return (
    <>
      <AppBar />
      {user ? getHomePage(user.accountType) : <Login />}
    </>
  );
}

export default App;
