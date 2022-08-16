import cogoToast from 'cogo-toast';
import { mutate } from 'swr';
import request from '../config/api';

export default async function currencyDelete(id) {
  try {
    const { data } = await request.delete(`/currencies/${id}`);
    mutate('/currencies');
    cogoToast.success(data.message, { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function fetchCurrencyRates(setLoading) {
  setLoading(true);
  try {
    const { data } = await request.post('/fetchrates');
    mutate('/currencies');
    cogoToast.success(data.message, { position: 'bottom-center' });
    setLoading(false);
    return data;
  } catch (error) {
    const { data } = error.response;
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function currencyAdd(params, setActionLoader) {
  setActionLoader(true);
  try {
    const formData = new FormData();
    if (params?.icon) {
      formData.append('icon', params?.icon[0]?.file, params?.icon[0]?.file?.name);
    }
    formData.append('name', params?.name);
    formData.append('symbol', params?.symbol);
    formData.append('crypto', params?.crypto);
    formData.append('active', params?.active);
    formData.append('rateUsd', params?.rateUsd);

    const { data } = await request.post('/currencies', formData);
    mutate('/currencies');
    setActionLoader(false);
    cogoToast.success('Currency Added', { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
export async function currencyUpdate(id, params, setActionLoader) {
  setActionLoader(true);
  try {
    const formData = new FormData();
    if (params?.icon) {
      formData.append('icon', params?.icon[0]?.file, params?.icon[0]?.file?.name);
    }
    formData.append('active', params?.active);

    const { data } = await request.put(`/currencies/${id}`, formData);
    mutate('/currencies');
    setActionLoader(false);
    cogoToast.success('Currency Updated', { position: 'bottom-center' });
    return data;
  } catch (error) {
    const { data } = error.response;
    setActionLoader(false);
    cogoToast.error(data.message, { position: 'bottom-center' });
  }
  return null;
}
