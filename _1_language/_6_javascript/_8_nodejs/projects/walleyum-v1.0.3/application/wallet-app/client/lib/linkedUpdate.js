import cogoToast from 'cogo-toast';
import { mutate } from 'swr';
import request from '../config/api';

export async function createLinked(params, id, setActionLoader) {
  setActionLoader(true);
  try {
    console.log({ params, methodId: id });
    const { data } = await request.post('/linkeds', { params, methodId: id });
    mutate('/linkeds/me');
    setActionLoader(false);
    cogoToast.success('Account Linked Successfully', { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}

export async function removeLinked(id) {
  try {
    const { data } = await request.delete(`/linkeds/${id}`);
    mutate('/linkeds/me');
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
