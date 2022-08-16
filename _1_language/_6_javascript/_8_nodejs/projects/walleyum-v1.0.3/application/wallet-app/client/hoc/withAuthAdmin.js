import { useRouter } from 'next/router';
import React from 'react';
import Loader from '../components/Loader';
import useCheckAuth from '../data/useCheckAuth';
import { useProfileAdmin } from '../data/useProfile';
import useSettings from '../data/useSettings';

const withAuthAdmin = (WrappedComponent) => (props) => {
  const Router = useRouter();

  const { data, loading } = useCheckAuth(true);
  const { data: userData, loading: userLoading } = useProfileAdmin();
  const { data: settings, loading: settingsLoading } = useSettings();

  if (loading || userLoading || settingsLoading) {
    return <Loader height="100vh" />;
  }

  if (data?.isAdmin) {
    return <WrappedComponent {...props} userData={userData} settings={settings} />;
  }
  Router.replace('/admin/login');
  return null;
};

export default withAuthAdmin;
