import React from 'react';
import { useSignInWithGoogle } from 'react-firebase-hooks/auth';
import { auth } from 'app/services/auth/firebase';
import Home from 'features/Client/Home';

const Login: React.FC = () => {
  const [signInWithGoogle, user, loading, error] = useSignInWithGoogle(auth);

  if (error) {
    return <div>Error!</div>;
  }

  if (loading) {
    return <div>loading...</div>;
  }

  if (!user) {
    console.log('not user');
    return (
      <button
        style={{ alignContent: 'center' }}
        onClick={() => {
          signInWithGoogle();
        }}
      >
        Sign in
      </button>
    );
  }

  return <Home />;
};

export default Login;
