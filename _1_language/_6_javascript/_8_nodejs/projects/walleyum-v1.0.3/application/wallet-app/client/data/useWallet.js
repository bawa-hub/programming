import useSWR from 'swr';

function useWallet() {
  const { data, error } = useSWR('/wallets/me');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useWalletByUserId(id) {
  const { data, error } = useSWR(`/wallets/${id}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export default useWallet;
