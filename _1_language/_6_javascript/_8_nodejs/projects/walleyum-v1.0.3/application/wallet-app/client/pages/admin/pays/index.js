import Head from 'next/head';
import React from 'react';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle } from 'react-icons/bi';
import PaysHistoryAdmin from '../../../components/admin/pays/PaysHistoryAdmin';
import SidebarAdmin from '../../../components/SidebarAdmin';
import TabModule from '../../../components/tabs/TabModule';
import UserTab from '../../../components/tabs/UserTab';
import UserHeaderAdmin from '../../../components/UserHeaderAdmin';
import withAuthAdmin from '../../../hoc/withAuthAdmin';

const Pays = ({ userData, settings }) => {
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
                title={t('Payments')}
                description="Manage all payments of your users from this panel"
              >
                <TabModule icon={<BiErrorCircle />} name={t('History')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Payment History')}</h4>
                    <PaysHistoryAdmin />
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

export default withAuthAdmin(Pays);
