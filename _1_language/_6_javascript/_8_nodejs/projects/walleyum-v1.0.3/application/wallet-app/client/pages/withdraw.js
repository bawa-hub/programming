/* eslint-disable max-len */
import Head from 'next/head';
import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle, BiWallet } from 'react-icons/bi';
import Sidebar from '../components/Sidebar';
import TabModule from '../components/tabs/TabModule';
import UserTab from '../components/tabs/UserTab';
import TransactionSteps from '../components/TransactionSteps';
import UserHeader from '../components/UserHeader';
import WithdrawHistory from '../components/withdraws/WithdrawHistory';
import WithdrawStep from '../components/withdraws/WithdrawStep';
import withAuth from '../hoc/withAuth';

const Withdraw = ({ userData, settings }) => {
  const [step, setStep] = useState(1);
  const { t } = useTranslation();
  return (
    <>
      <Head>
        <title>
          {t('Withdraw')}
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
                title={t('Withdraw Money')}
                description={t('Withdraw funds to your desired method')}
              >
                <TabModule icon={<BiWallet />} name={t('Withdraw')}>
                  <div className="deposit-box basic-card">
                    <TransactionSteps step={step} />
                    <WithdrawStep step={step} setStep={setStep} />
                  </div>
                </TabModule>
                <TabModule icon={<BiErrorCircle />} name={t('Withdraw Logs')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Withdraw Logs')}</h4>
                    <WithdrawHistory />
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

export default withAuth(Withdraw);
