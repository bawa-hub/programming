import useSWR from 'swr';

export default function useSettings() {
  const { data, error } = useSWR('/settings');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
