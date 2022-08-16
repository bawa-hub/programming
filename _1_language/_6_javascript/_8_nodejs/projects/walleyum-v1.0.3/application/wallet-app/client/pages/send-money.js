import Head from 'next/head';
import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle, BiWallet } from 'react-icons/bi';
import Sidebar from '../components/Sidebar';
import TabModule from '../components/tabs/TabModule';
import UserTab from '../components/tabs/UserTab';
import TransactionSteps from '../components/TransactionSteps';
import ReceiveHistory from '../components/transfers/ReceiveHistory';
import SendHistory from '../components/transfers/SendHistory';
import TransferStep from '../components/transfers/TransferStep';
import UserHeader from '../components/UserHeader';
import withAuth from '../hoc/withAuth';

const Send = ({ userData, settings }) => {
  const [step, setStep] = useState(1);
  const { t } = useTranslation();
  return (
    <>
      <Head>
        <title>
          {t('Send Money')}
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
                title={t('Send Money')}
                description={t('Send money to other users from your wallet')}
              >
                <TabModule icon={<BiWallet />} name={t('Send Money')}>
                  <div className="deposit-box basic-card">
                    <TransactionSteps step={step} />
                    <TransferStep step={step} setStep={setStep} />
                  </div>
                </TabModule>
                <TabModule icon={<BiErrorCircle />} name={t('Send Logs')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Send Logs')}</h4>
                    <SendHistory />
                  </div>
                </TabModule>
                <TabModule icon={<BiErrorCircle />} name={t('Receive Logs')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Receive Logs')}</h4>
                    <ReceiveHistory />
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

export default withAuth(Send);
