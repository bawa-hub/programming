import useSWR from 'swr';

export default function useUsers(page, limit, status, id, email, role) {
  const idFilter = !(id === '') ? `id=like:%${id}%` : '';
  const emailFilter = !(email === '') ? `email=like:%${email}%` : '';
  const statusFilter = !(status === 'all') ? `active=${status}` : '';
  const roleFilter = !(role === 'all') ? `role=${role}` : '';
  const { data, error } = useSWR(`/users?${statusFilter}${idFilter}${emailFilter}${roleFilter}&sort_by=createdAt.desc&offset=${page - 1}&limit=${limit}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useUsersDashboard(page, limit) {
  const { data, error } = useSWR(`/users?sort_by=createdAt.desc&offset=${page - 1}&limit=${limit}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useUserById(id) {
  const { data, error } = useSWR(`/users/${id}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
