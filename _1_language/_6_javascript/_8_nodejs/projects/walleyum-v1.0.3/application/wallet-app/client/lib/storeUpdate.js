import cogoToast from 'cogo-toast';
import { mutate } from 'swr';
import request from '../config/api';

export default async function storeUpdate(params, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.put('/merchants/me', { ...params });
    setActionLoader(false);
    mutate('/users/me');
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
