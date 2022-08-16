import useSWR from 'swr';

function usePages() {
  const { data, error } = useSWR('/pages');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function usePageBySlug(slug) {
  const { data, error } = useSWR(`/pages/${slug}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export default usePages;
