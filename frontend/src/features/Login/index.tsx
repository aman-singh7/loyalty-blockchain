import { UserContext } from 'app/contexts/userContext';
import React, { useContext } from 'react';
import './index.scss';

const Login: React.FC = () => {
  const { user, signIn, isLoading } = useContext(UserContext);

  if (isLoading) {
    return <div className='loading-page'>Loading...</div>;
  }

  if (!user) {
    return (
      <div className='login-page'>
        <div onClick={signIn} className='button'>
          Login as Consumer
        </div>
        <div onClick={signIn} className='button'>
          Login as Business
        </div>
      </div>
    );
  }

  if (user.walletId.length == 0) {
    return (
      <div className='login-page'>
        <div className='button'>Connect With Metmask</div>
      </div>
    );
  }

  return <>UnAuthorized</>;
};

export default Login;
