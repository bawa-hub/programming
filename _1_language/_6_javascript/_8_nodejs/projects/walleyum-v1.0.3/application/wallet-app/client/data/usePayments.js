import useSWR from 'swr';

export default function usePayments(page, limit) {
  const { data, error } = useSWR(`/pays?sort_by=createdAt.desc&offset=${page - 1}&limit=${limit}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function usePaymentsAdmin(page, limit, status, id) {
  const idFilter = !(id === '') ? `trxId=like:%${id}%` : '';
  const statusFilter = !(status === 'all') ? `status=${status}` : '';
  const { data, error } = useSWR(`/pays/admin?${statusFilter}${idFilter}&sort_by=createdAt.desc&offset=${page - 1}&limit=${limit}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
