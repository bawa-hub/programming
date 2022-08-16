import useSWR from 'swr';

export default function useCheckAuth(admin) {
  const { data, error } = useSWR(admin ? '/checkauth/admin' : '/checkauth', {
    revalidateIfStale: false,
    revalidateOnFocus: false,
    revalidateOnReconnect: true
  });

  return {
    data,
    loading: !data && !error,
    error,
  };
}
