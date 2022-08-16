import useSWR from 'swr';

function useGateways() {
  const { data, error } = useSWR('/gateways');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useGatewayCurrencies() {
  const { data, error } = useSWR('/gateways/currencies');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useGatewayAdmin() {
  const { data, error } = useSWR('/gateways/admin');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useGatewayByValue(value) {
  const { data, error } = useSWR(`/gateways/${value}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export default useGateways;
