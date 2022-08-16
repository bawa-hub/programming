import useSWR from 'swr';

function useLogs() {
  const { data, error } = useSWR('/logs');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export default useLogs;
