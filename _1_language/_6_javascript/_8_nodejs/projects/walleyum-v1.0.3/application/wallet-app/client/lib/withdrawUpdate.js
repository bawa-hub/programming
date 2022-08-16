import cogoToast from 'cogo-toast';
import { mutate } from 'swr';
import request from '../config/api';

export default async function withdrawUpdate(id, type, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.put(`/withdraws/${id}/${type}`);
    setActionLoader(false);
    mutate(`/withdraws/${id}/admin`);
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
