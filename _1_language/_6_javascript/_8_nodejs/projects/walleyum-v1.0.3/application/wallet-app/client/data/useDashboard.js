import useSWR from 'swr';

function useDashboard() {
  const { data, error } = useSWR('/dashboard');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export default useDashboard;
