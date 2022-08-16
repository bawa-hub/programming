import cogoToast from 'cogo-toast';
import { mutate } from 'swr';
import request from '../config/api';

export default async function merchantUpdate(id, params, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.put(`/merchants/${id}`, { ...params });
    setActionLoader(false);
    mutate(`/merchants/${id}`);
    cogoToast.success('Merchant Updated', { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function merchantDelete(id, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.delete(`/merchants/${id}`);
    setActionLoader(false);
    window.location.href = '/admin/merchants';
    cogoToast.success('Merchant Removed', { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function merchantDeleteFromUser(userId, id, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.delete(`/merchants/${id}`);
    setActionLoader(false);
    window.location.href = `/admin/users/${userId}`;
    cogoToast.success('Merchant Removed', { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function merchantCreate(params, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.post('/merchants', { ...params });
    setActionLoader(false);
    mutate(`/users/${params?.userId}`);
    cogoToast.success('User Converted To Merchant', { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
