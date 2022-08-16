import cogoToast from 'cogo-toast';
import { mutate } from 'swr';
import request from '../config/api';

export default async function gatewayUpdate(value, params, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.put(`/payments/${value}`, { ...params });
    setActionLoader(false);
    mutate('/gateways/admin');
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
