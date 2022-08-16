/* eslint-disable max-len */
import Head from 'next/head';
import { useRouter } from 'next/router';
import React from 'react';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle, BiWallet } from 'react-icons/bi';
import RequestPayment from '../components/merchant/RequestPayment';
import StoreSettings from '../components/merchant/StoreSettings';
import TransactionHistory from '../components/merchant/TransactionHistory';
import Sidebar from '../components/Sidebar';
import TabModule from '../components/tabs/TabModule';
import UserTab from '../components/tabs/UserTab';
import UserHeader from '../components/UserHeader';
import withAuth from '../hoc/withAuth';

const Merchant = ({ userData, settings }) => {
  const router = useRouter();
  const { tab } = router.query;
  const { t } = useTranslation();

  return (
    <>
      <Head>
        <title>
          {t('Merchant')}
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
                title={t('Merchant')}
                description={t('Manage your merchant settings and request payment')}
                defaultTab={parseInt(tab, 10)}
              >
                <TabModule icon={<BiWallet />} name="Request Payment">
                  <RequestPayment />
                </TabModule>
                <TabModule icon={<BiErrorCircle />} name={t('Recent Transactions')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Recent Transactions')}</h4>
                    <TransactionHistory />
                  </div>
                </TabModule>
                <TabModule icon={<BiErrorCircle />} name={t('Store Settings')}>
                  <StoreSettings userData={userData} />
                </TabModule>
              </UserTab>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default withAuth(Merchant);
