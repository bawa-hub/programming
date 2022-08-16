import useSWR from 'swr';

export default function useWithdraws(page, limit) {
  const { data, error } = useSWR(`/withdraws?sort_by=createdAt.desc&offset=${page - 1}&limit=${limit}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useWithdrawsAdmin(page, limit, status, id) {
  const idFilter = !(id === '') ? `id=like:%${id}%` : '';
  const statusFilter = !(status === 'all') ? `status=${status}` : '';
  const { data, error } = useSWR(`/withdraws/admin?${statusFilter}${idFilter}&sort_by=createdAt.desc&offset=${page - 1}&limit=${limit}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useWithdrawById(id) {
  const { data, error } = useSWR(`/withdraws/${id}/admin`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
