import Head from 'next/head';
import { useRouter } from 'next/router';
import React, { useState } from 'react';
import Spinner from 'react-bootstrap/Spinner';
import { useTranslation } from 'react-i18next';
import {
  BiCog, BiGroup, BiPlusCircle, BiPound, BiPurchaseTag, BiSave, BiShuffle
} from 'react-icons/bi';
import AdjSettings from '../../../components/admin/settings/AdjSettings';
import ApiSettings from '../../../components/admin/settings/ApiSettings';
import GeneralSettings from '../../../components/admin/settings/GeneralSettings';
import RefferalSettings from '../../../components/admin/settings/RefferalSettings';
import WalletList from '../../../components/admin/settings/WalletList';
import SidebarAdmin from '../../../components/SidebarAdmin';
import TabModule from '../../../components/tabs/TabModule';
import UserTab from '../../../components/tabs/UserTab';
import UserHeaderAdmin from '../../../components/UserHeaderAdmin';
import withAuthAdmin from '../../../hoc/withAuthAdmin';
import { fetchCurrencyRates } from '../../../lib/currencyRequest';

const Settings = ({ userData, settings }) => {
  const router = useRouter();
  const { t } = useTranslation();
  const [loading, setLoading] = useState(false);
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
                title={t('Settings')}
                description={t('Manage site settings from this panel')}
              >
                <TabModule icon={<BiCog />} name={t('General')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('General Settings')}</h4>
                    <GeneralSettings settings={settings} />
                  </div>
                </TabModule>
                <TabModule icon={<BiPound />} name={t('Wallets')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Wallets')}</h4>
                    <WalletList settings={settings} />
                    <button
                      type="button"
                      className="bttn-mid btn-primary btn-new mt-20"
                      onClick={() => router.push('/admin/settings/wallet/new')}
                    >
                      <BiPlusCircle />
                      {t('Add New Wallet')}
                    </button>
                    <button
                      type="button"
                      className="bttn-mid btn-ylo btn-new"
                      onClick={() => fetchCurrencyRates(setLoading)}
                      disabled={loading}
                    >
                      {loading ? (
                        <>
                          <Spinner animation="border" role="status" size="sm" />
                          {' '}
                          {t('Fetch Currency Rates')}
                        </>
                      ) : (
                        <>
                          <BiSave />
                          {t('Fetch Currency Rates')}
                        </>
                      )}
                    </button>
                  </div>
                </TabModule>
                <TabModule icon={<BiGroup />} name={t('Refferal')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Refferal')}</h4>
                    <RefferalSettings settings={settings} />
                  </div>
                </TabModule>
                <TabModule icon={<BiPurchaseTag />} name={t('Fees')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Fees')}</h4>
                    <AdjSettings settings={settings} />
                  </div>
                </TabModule>
                <TabModule icon={<BiShuffle />} name={t('APIs')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('APIs')}</h4>
                    <ApiSettings settings={settings} />
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

export default withAuthAdmin(Settings);
