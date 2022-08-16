import Head from 'next/head';
import { useRouter } from 'next/router';
import React from 'react';
import { useTranslation } from 'react-i18next';
import {
  BiAnalyse, BiErrorCircle, BiPlusCircle, BiWallet
} from 'react-icons/bi';
import WithdrawHistoryAdmin from '../../../components/admin/withdraw/WithdrawHistoryAdmin';
import WithdrawMethodsList from '../../../components/admin/withdraw/WithdrawMethodsList';
import SidebarAdmin from '../../../components/SidebarAdmin';
import TabModule from '../../../components/tabs/TabModule';
import UserTab from '../../../components/tabs/UserTab';
import UserHeaderAdmin from '../../../components/UserHeaderAdmin';
import withAuthAdmin from '../../../hoc/withAuthAdmin';

const Withdraws = ({ userData, settings }) => {
  const router = useRouter();
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
                title={t('Withdraws')}
                description={t('Manage all withdraws of your users from this panel')}
              >
                <TabModule icon={<BiAnalyse />} name={t('Pending')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Pending Withdraws')}</h4>
                    <WithdrawHistoryAdmin pending />
                  </div>
                </TabModule>
                <TabModule icon={<BiErrorCircle />} name={t('History')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Withdraw History')}</h4>
                    <WithdrawHistoryAdmin />
                  </div>
                </TabModule>
                <TabModule icon={<BiWallet />} name={t('Methods')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Withdraw Methods')}</h4>
                    <WithdrawMethodsList />
                    <button
                      type="button"
                      className="bttn-mid btn-primary btn-new mt-20"
                      onClick={() => router.push('/admin/withdraws/method/new')}
                    >
                      <BiPlusCircle />
                      {t('Add New Method')}
                    </button>
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

export default withAuthAdmin(Withdraws);
