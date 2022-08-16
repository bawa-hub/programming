/* eslint-disable max-len */
import Head from 'next/head';
import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle, BiWallet } from 'react-icons/bi';
import ExchangeHistory from '../components/exchanges/ExchangeHistory';
import ExchangeStep from '../components/exchanges/ExchangeStep';
import Sidebar from '../components/Sidebar';
import TabModule from '../components/tabs/TabModule';
import UserTab from '../components/tabs/UserTab';
import TransactionSteps from '../components/TransactionSteps';
import UserHeader from '../components/UserHeader';
import withAuth from '../hoc/withAuth';

const Exchange = ({ userData, settings }) => {
  const [step, setStep] = useState(1);
  const { t } = useTranslation();
  return (
    <>
      <Head>
        <title>
          {t('Exchange')}
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
                title={t('Exchange Money')}
                description={t('Exchange money from one wallet to another')}
              >
                <TabModule icon={<BiWallet />} name={t('Exchange')}>
                  <div className="deposit-box basic-card">
                    <TransactionSteps step={step} />
                    <ExchangeStep step={step} setStep={setStep} settings={settings} />
                  </div>
                </TabModule>
                <TabModule icon={<BiErrorCircle />} name={t('Exchange Logs')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Exchange Logs')}</h4>
                    <ExchangeHistory />
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

export default withAuth(Exchange);
