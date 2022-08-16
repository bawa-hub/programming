import { useRouter } from 'next/router';
import React from 'react';
import Loader from '../components/Loader';
import useCheckAuth from '../data/useCheckAuth';

const checkAuthAccessAdmin = (WrappedComponent) => (props) => {
  const Router = useRouter();
  const { data, loading } = useCheckAuth(true);

  if (loading) {
    return <Loader height="100vh" />;
  }

  if (!data || !(data?.isAdmin)) {
    return <WrappedComponent {...props} />;
  }
  if (data?.isAdmin) {
    Router.replace('/admin/dashboard');
  } else {
    Router.replace('/dashboard');
  }
  return null;
};

export default checkAuthAccessAdmin;
