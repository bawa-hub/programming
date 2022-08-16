import cogoToast from 'cogo-toast';
import { mutate } from 'swr';
import request from '../config/api';

export default async function kycUpdate(params, setActionLoader) {
  setActionLoader(true);
  try {
    const formData = new FormData();
    formData.append('front', params.front[0].file, params.front[0].file.name);
    if (params.back) {
      formData.append('back', params.back[0].file, params.back[0].file.name);
    }
    formData.append('selfie', params.selfie[0].file, params.selfie[0].file.name);
    formData.append('type', params.type);
    const { data } = await request.post('/kyc', formData);
    setActionLoader(false);
    mutate('/kyc/me');
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function kycResubmit(setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.post('/kyc/resubmit');
    setActionLoader(false);
    mutate('/kyc/me');
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function kycAdminAction(type, id, userId, setActionLoader) {
  setActionLoader(true);
  try {
    const { data } = await request.put(`/kyc/${id}/${type}`);
    setActionLoader(false);
    mutate(`/users/${userId}`);
    mutate(`/kyc/${id}`);
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
