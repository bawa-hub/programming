import cogoToast from 'cogo-toast';
import { mutate } from 'swr';
import request from '../config/api';

export default async function menuUpdate(params, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.put('/menu/main', { ...params });
    setActionLoader(false);
    mutate('/menu/main');
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function footerMenuUpdate(params, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.put('/menu/footer', { ...params });
    setActionLoader(false);
    mutate('/menu/footer');
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function repeaterUpdate(value, params, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.put(`/repeater/${value}`, { ...params });
    setActionLoader(false);
    mutate('/info');
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
