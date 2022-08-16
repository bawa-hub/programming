import useSWR from 'swr';

export default function useRequests(page, limit) {
  const { data, error } = useSWR(`/requests?sort_by=createdAt.desc&offset=${page - 1}&limit=${limit}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useRequestByTrx(trxId) {
  const { data, error } = useSWR(`/requests/trx/${trxId}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useRequestsAdmin(page, limit, status, id) {
  const idFilter = !(id === '') ? `trxId=like:%${id}%` : '';
  const statusFilter = !(status === 'all') ? `status=${status}` : '';
  const { data, error } = useSWR(`/requests/admin?${statusFilter}${idFilter}&sort_by=createdAt.desc&offset=${page - 1}&limit=${limit}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
