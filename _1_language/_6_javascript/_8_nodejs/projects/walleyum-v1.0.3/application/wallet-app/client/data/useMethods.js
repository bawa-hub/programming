import useSWR from 'swr';

function useMethods() {
  const { data, error } = useSWR('/methods');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useMethodsAdmin() {
  const { data, error } = useSWR('/methods/admin');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useMethodById(id) {
  const { data, error } = useSWR(`/methods/${id}`, {
    revalidateIfStale: false,
    revalidateOnFocus: false,
    revalidateOnReconnect: false
  });
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useMethodByIdAdmin(id) {
  const { data, error } = useSWR(`/methods/${id}/admin`, {
    revalidateIfStale: false,
    revalidateOnFocus: false,
    revalidateOnReconnect: false
  });
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export default useMethods;
