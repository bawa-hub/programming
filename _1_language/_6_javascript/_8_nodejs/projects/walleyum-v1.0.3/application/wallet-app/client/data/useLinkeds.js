import useSWR from 'swr';

function useLinkeds() {
  const { data, error } = useSWR('/linkeds/me');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export default useLinkeds;
