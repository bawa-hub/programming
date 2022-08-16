import useSWR from 'swr';

export default function useExchanges(page, limit) {
  const { data, error } = useSWR(`/exchanges?sort_by=createdAt.desc&offset=${page - 1}&limit=${limit}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useExchangeById(id) {
  const { data, error } = useSWR(`/exchanges/${id}/admin`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useExchangesAdmin(page, limit, status, id) {
  const idFilter = !(id === '') ? `id=like:%${id}%` : '';
  const statusFilter = !(status === 'all') ? `status=${status}` : '';
  const { data, error } = useSWR(`/exchanges/admin?${statusFilter}${idFilter}&sort_by=createdAt.desc&offset=${page - 1}&limit=${limit}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
