import useSWR from 'swr';

export default function useKyc() {
  const { data, error } = useSWR('/kyc/me');
  return {
    data,
    loading: !data && !error,
    error,
  };
}
export function useKycAll(page, limit) {
  const { data, error } = useSWR(`/kyc?status=submitted&sort_by=createdAt.desc&offset=${page - 1}&limit=${limit}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}

export function useKycById(id) {
  const { data, error } = useSWR(`/kyc/${id}`);
  return {
    data,
    loading: !data && !error,
    error,
  };
}
