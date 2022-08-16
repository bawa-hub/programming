import useSWR from 'swr';

function useInfo() {
  const { data, error } = useSWR('/info');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export default useInfo;
