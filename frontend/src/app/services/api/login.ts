import axios from './axios';

export const login = async (uid: string, token: string) => {
  const data = { uid, token };
  const res = await axios.post('/v1/auth', data);
  console.log(res);
  return res;
};
