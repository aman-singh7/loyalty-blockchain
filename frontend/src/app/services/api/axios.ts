import axios from 'axios';
import { apiURL } from 'app/constants';

const basicAxios = axios.create({
  baseURL: apiURL,
  withCredentials: true,
});

export default basicAxios;
