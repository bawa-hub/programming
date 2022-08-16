import cogoToast from 'cogo-toast';
import { mutate } from 'swr';
import request from '../config/api';

export default async function profileUpdate(params, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.put('/users/me', { ...params });
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
export async function profileUpdateAdmin(id, params, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.put(`/users/${id}`, { ...params });
    setActionLoader(false);
    mutate(`/users/${id}`);
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function userDelete(id, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.delete(`/users/${id}`);
    setActionLoader(false);
    window.location.href = '/admin/users';
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
