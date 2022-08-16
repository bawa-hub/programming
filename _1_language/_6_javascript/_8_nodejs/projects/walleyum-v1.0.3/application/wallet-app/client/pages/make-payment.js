import Head from 'next/head';
import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle, BiWallet } from 'react-icons/bi';
import PaymentHistory from '../components/payment/PaymentHistory';
import PaymentStep from '../components/payment/PaymentStep';
import Sidebar from '../components/Sidebar';
import TabModule from '../components/tabs/TabModule';
import UserTab from '../components/tabs/UserTab';
import TransactionSteps from '../components/TransactionSteps';
import UserHeader from '../components/UserHeader';
import withAuth from '../hoc/withAuth';

const MakePayment = ({ userData, settings }) => {
  const [step, setStep] = useState(1);
  const { t } = useTranslation();
  return (
    <>
      <Head>
        <title>
          {t('Payment')}
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
                title={t('Make Payment')}
                description={t('Make payment to different merchants')}
              >
                <TabModule icon={<BiWallet />} name={t('Make Payment')}>
                  <div className="deposit-box basic-card">
                    <TransactionSteps step={step} />
                    <PaymentStep step={step} setStep={setStep} />
                  </div>
                </TabModule>
                <TabModule icon={<BiErrorCircle />} name={t('Payment Logs')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Payment Logs')}</h4>
                    <PaymentHistory />
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

export default withAuth(MakePayment);
