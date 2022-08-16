import cogoToast from 'cogo-toast';
import { mutate } from 'swr';
import request from '../config/api';

export default async function generalUpdate(params, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.put('/settings/general', { ...params });
    setActionLoader(false);
    mutate('/settings');
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function settingsUpdate(value, params, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.put(`/settings/byvalue/${value}`, { ...params });
    setActionLoader(false);
    mutate('/settings');
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
export async function logoUpdate(params, setActionLoader) {
  setActionLoader(true);
  try {
    const formData = new FormData();
    formData.append('logo', params.logo[0].file, params.logo[0].file.name);
    formData.append('favicon', params.favicon[0].file, params.favicon[0].file.name);
    const { data } = await request.put('/logo', formData);
    setActionLoader(false);
    mutate('/settings');
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function sendMail(params, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.post('/email', { ...params });
    setActionLoader(false);
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function balanceManage(params, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.post('/balance', { ...params });
    setActionLoader(false);
    mutate(`/wallets/${params?.userId}`);
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function imageUpload(params, setActionLoader) {
  setActionLoader(true);
  try {
    const formData = new FormData();
    formData.append('image', params, params.name);
    const { data } = await request.post('/upload', formData);
    setActionLoader(false);
    cogoToast.success('Image Uploaded', { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message || 'Image Upload Failed', { position: 'bottom-center' });
  }
  return null;
}
