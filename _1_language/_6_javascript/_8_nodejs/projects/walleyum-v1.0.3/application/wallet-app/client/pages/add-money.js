/* eslint-disable jsx-a11y/label-has-associated-control */
import Head from 'next/head';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle, BiWallet } from 'react-icons/bi';
import DepositHistory from '../components/deposit/DepositHistory';
import DepositStep from '../components/deposit/DepositStep';
import Sidebar from '../components/Sidebar';
import TabModule from '../components/tabs/TabModule';
import UserTab from '../components/tabs/UserTab';
import TransactionSteps from '../components/TransactionSteps';
import UserHeader from '../components/UserHeader';
import withAuth from '../hoc/withAuth';

const AddMoney = ({ userData, settings }) => {
  const [step, setStep] = useState(1);
  const router = useRouter();
  const { status } = router.query;

  const { t } = useTranslation();

  useEffect(() => {
    if (status === 'success' || status === 'failed') {
      setStep(3);
    }
  }, []);

  return (
    <>
      <Head>
        <title>
          {t('Deposit')}
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
                title={t('Add Money')}
                description={t('Add money to your wallet by using different payment methods')}
              >
                <TabModule icon={<BiWallet />} name={t('Deposit')}>
                  <div className="deposit-box basic-card">
                    <TransactionSteps step={step} />
                    <DepositStep step={step} setStep={setStep} status={status} />
                  </div>
                </TabModule>
                <TabModule icon={<BiErrorCircle />} name={t('History')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Deposit Logs')}</h4>
                    <DepositHistory />
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

export default withAuth(AddMoney);
