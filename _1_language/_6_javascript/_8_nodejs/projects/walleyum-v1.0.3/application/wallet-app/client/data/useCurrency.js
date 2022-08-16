import useSWR from 'swr';

function useCurrency() {
  const { data, error } = useSWR('/currencies');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useCurrencyList() {
  const { data, error } = useSWR('/currencies/list');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useCurrencyById(id) {
  const { data, error } = useSWR(`/currencies/${id}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export default useCurrency;
