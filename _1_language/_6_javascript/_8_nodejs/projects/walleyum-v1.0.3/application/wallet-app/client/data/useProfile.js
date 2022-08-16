import useSWR from 'swr';

export default function useProfile() {
  const { data, error } = useSWR('/users/me');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useProfileAdmin() {
  const { data, error } = useSWR('/users/me/admin');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
