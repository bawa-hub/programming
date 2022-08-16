import useSWR from 'swr';

export function useSend(page, limit) {
  const { data, error } = useSWR(`/transfers?type=send&sort_by=createdAt.desc&offset=${page - 1}&limit=${limit}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useReceive(page, limit) {
  const { data, error } = useSWR(`/transfers?type=receive&sort_by=createdAt.desc&offset=${page - 1}&limit=${limit}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useTransfer(page, limit, status, id) {
  const idFilter = !(id === '') ? `trxId=like:%${id}%` : '';
  const statusFilter = !(status === 'all') ? `type=${status}` : '';
  const { data, error } = useSWR(`/transfers/admin?${statusFilter}${idFilter}&sort_by=createdAt.desc&offset=${page - 1}&limit=${limit}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
