import cogoToast from 'cogo-toast';
import { mutate } from 'swr';
import request from '../config/api';

export default async function pageAdd(params) {
  try {
    const { data } = await request.post('/pages', { ...params });
    mutate('/pages');
    cogoToast.success('Page Added Successfully', { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function pageDelete(slug) {
  try {
    const { data } = await request.delete(`/pages/${slug}`);
    mutate('/pages');
    cogoToast.success('Page Deleted Successfully', { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function pageUpdate(slug, params, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.put(`/pages/${slug}`, { ...params });
    setActionLoader(false);
    mutate(`/pages/${slug}`);
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
