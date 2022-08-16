import useSWR from 'swr';

export default function useMerchants(page, limit, status, id) {
  const idFilter = !(id === '') ? `merId=like:%${id}%` : '';
  const statusFilter = !(status === 'all') ? `status=${status}` : '';
  const { data, error } = useSWR(`/merchants?${statusFilter}${idFilter}&sort_by=createdAt.desc&offset=${page - 1}&limit=${limit}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useMerchantById(id) {
  const { data, error } = useSWR(`/merchants/${id}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
