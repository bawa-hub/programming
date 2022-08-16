/* eslint-disable max-len */
import Head from 'next/head';
import { useRouter } from 'next/router';
import React from 'react';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle, BiWallet } from 'react-icons/bi';
import Kyc from '../components/settings/Kyc';
import LinkedAcc from '../components/settings/LinkedAcc';
import ProfileSettings from '../components/settings/ProfileSettings';
import Sidebar from '../components/Sidebar';
import TabModule from '../components/tabs/TabModule';
import UserTab from '../components/tabs/UserTab';
import UserHeader from '../components/UserHeader';
import withAuth from '../hoc/withAuth';

const Settings = ({ userData, settings }) => {
  const router = useRouter();
  const { tab } = router.query;
  const { t } = useTranslation();

  return (
    <>
      <Head>
        <title>
          {t('Settings')}
          {' '}
          -
          {' '}
          {settings?.site?.param1}
        </title>
        <link rel="icon" href={`${settings?.apiUrl?.param1}/public/${settings?.logo?.param2}`} />
      </Head>
      <UserHeader />
      <Sidebar userData={userData} settings={settings} />
      <div className="content-body">
        <div className="container-fluid">
          <div className="row">
            <div className="col">
              <UserTab
                title={t('Settings')}
                description={t('Manage your account settings, KYC, linked accounts')}
                defaultTab={parseInt(tab, 10)}
              >
                <TabModule icon={<BiWallet />} name={t('Account Settings')}>
                  <ProfileSettings userData={userData} />
                </TabModule>
                <TabModule icon={<BiErrorCircle />} name={t('KYC')}>
                  <Kyc />
                </TabModule>
                <TabModule icon={<BiErrorCircle />} name={t('Linked Accounts')}>
                  <LinkedAcc />
                </TabModule>
              </UserTab>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default withAuth(Settings);
