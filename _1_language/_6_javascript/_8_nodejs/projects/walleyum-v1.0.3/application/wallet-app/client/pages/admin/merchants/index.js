import Head from 'next/head';
import React from 'react';
import { useTranslation } from 'react-i18next';
import { BiAnalyse, BiErrorCircle, BiLeftTopArrowCircle } from 'react-icons/bi';
import MerchantList from '../../../components/admin/merchant/MerchantList';
import RequestsList from '../../../components/admin/merchant/RequestsList';
import SidebarAdmin from '../../../components/SidebarAdmin';
import TabModule from '../../../components/tabs/TabModule';
import UserTab from '../../../components/tabs/UserTab';
import UserHeaderAdmin from '../../../components/UserHeaderAdmin';
import withAuthAdmin from '../../../hoc/withAuthAdmin';

const Merchants = ({ userData, settings }) => {
  const { t } = useTranslation();
  return (
    <>
      <Head>
        <title>
          {t('Admin Panel')}
        </title>
      </Head>
      <UserHeaderAdmin />
      <SidebarAdmin userData={userData} settings={settings} />
      <div className="content-body">
        <div className="container-fluid">
          <div className="row">
            <div className="col">
              <UserTab
                title={t('Merchants')}
                description={t('Manage your merchants from this panel')}
              >
                <TabModule icon={<BiAnalyse />} name={t('Merchant List')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Merchant List')}</h4>
                    <MerchantList />
                  </div>
                </TabModule>
                <TabModule icon={<BiErrorCircle />} name={t('Pending Verification')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Pending Verification')}</h4>
                    <MerchantList pending />
                  </div>
                </TabModule>
                <TabModule icon={<BiLeftTopArrowCircle />} name={t('Payment Requests')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Payment Requests')}</h4>
                    <RequestsList />
                  </div>
                </TabModule>
              </UserTab>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default withAuthAdmin(Merchants);
