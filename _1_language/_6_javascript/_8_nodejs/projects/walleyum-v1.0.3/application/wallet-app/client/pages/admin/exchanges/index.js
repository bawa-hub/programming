import Head from 'next/head';
import React from 'react';
import { useTranslation } from 'react-i18next';
import { BiAnalyse, BiErrorCircle } from 'react-icons/bi';
import ExchangeHistoryAdmin from '../../../components/admin/exchange/ExchangeHistoryAdmin';
import SidebarAdmin from '../../../components/SidebarAdmin';
import TabModule from '../../../components/tabs/TabModule';
import UserTab from '../../../components/tabs/UserTab';
import UserHeaderAdmin from '../../../components/UserHeaderAdmin';
import withAuthAdmin from '../../../hoc/withAuthAdmin';

const Exchanges = ({ userData, settings }) => {
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
                title={t('Exchanges')}
                description={t('Manage all currency exchanges of your users from this panel')}
              >
                <TabModule icon={<BiAnalyse />} name={t('Pending')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Pending Exchanges')}</h4>
                    <ExchangeHistoryAdmin pending />
                  </div>
                </TabModule>
                <TabModule icon={<BiErrorCircle />} name={t('History')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Exchange History')}</h4>
                    <ExchangeHistoryAdmin />
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

export default withAuthAdmin(Exchanges);
