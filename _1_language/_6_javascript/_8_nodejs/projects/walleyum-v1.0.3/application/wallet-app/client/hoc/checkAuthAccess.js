import { useRouter } from 'next/router';
import React from 'react';
import Loader from '../components/Loader';
import useCheckAuth from '../data/useCheckAuth';

const checkAuthAccess = (WrappedComponent) => (props) => {
  const Router = useRouter();
  const { data, loading } = useCheckAuth();

  if (loading) {
    return <Loader height="100vh" />;
  }

  if (!data) {
    return <WrappedComponent {...props} />;
  }
  if (data?.isAdmin) {
    Router.replace('/admin/dashboard');
  } else {
    Router.replace('/dashboard');
  }
  return null;
};

export default checkAuthAccess;
