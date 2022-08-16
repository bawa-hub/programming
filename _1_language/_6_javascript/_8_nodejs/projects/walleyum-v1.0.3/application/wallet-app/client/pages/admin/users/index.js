import Head from 'next/head';
import React from 'react';
import { useTranslation } from 'react-i18next';
import { BiAnalyse, BiErrorCircle } from 'react-icons/bi';
import KycList from '../../../components/admin/user/KycList';
import UserList from '../../../components/admin/user/UserList';
import SidebarAdmin from '../../../components/SidebarAdmin';
import TabModule from '../../../components/tabs/TabModule';
import UserTab from '../../../components/tabs/UserTab';
import UserHeaderAdmin from '../../../components/UserHeaderAdmin';
import withAuthAdmin from '../../../hoc/withAuthAdmin';

const Users = ({ userData, settings }) => {
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
                title={t('Users')}
                description={t('Manage your users from this panel')}
              >
                <TabModule icon={<BiAnalyse />} name={t('User List')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('User List')}</h4>
                    <UserList />
                  </div>
                </TabModule>
                <TabModule icon={<BiErrorCircle />} name={t('Pending Verification')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Pending Verification')}</h4>
                    <KycList />
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

export default withAuthAdmin(Users);
