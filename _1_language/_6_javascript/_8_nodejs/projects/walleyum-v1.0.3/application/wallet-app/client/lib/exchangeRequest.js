import cogoToast from 'cogo-toast';
import { mutate } from 'swr';
import request from '../config/api';

export default async function exchangeRequest(params, setActionLoader, setStep) {
  setActionLoader(true);
  try {
    const { data } = await request.post('/exchanges', { ...params });
    setActionLoader(false);
    setStep(3);
    mutate('/users/me');
    mutate('/wallets/me');
    mutate('/exchanges?sort_by=createdAt.desc&offset=0&limit=10');
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
