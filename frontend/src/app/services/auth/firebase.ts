import { FirebaseApp, FirebaseOptions, initializeApp } from 'firebase/app';
import { Auth, getAuth } from 'firebase/auth';
import { Analytics, getAnalytics } from 'firebase/analytics';

const firebaseConfig: FirebaseOptions = {
  apiKey: process.env.REACT_APP_FIREBASE_API_KEY,
  authDomain: 'grid-5-4039c.firebaseapp.com',
  projectId: 'grid-5-4039c',
  storageBucket: 'grid-5-4039c.appspot.com',
  messagingSenderId: '134472853877',
  appId: '1:134472853877:web:9ae040c331a339508f4b3c',
  measurementId: 'G-GWTPZM9ZQV',
};

// Initialize Firebase
export const app: FirebaseApp = initializeApp(firebaseConfig);
export const auth: Auth = getAuth(app);
export const analytics: Analytics = getAnalytics(app);
