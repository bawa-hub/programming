import Head from 'next/head';
import { useRouter } from 'next/router';
import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import CheckoutStep from '../components/payment/CheckoutStep';
import TransactionSteps from '../components/TransactionSteps';
import withAuth from '../hoc/withAuth';

const Checkout = () => {
  const [step, setStep] = useState(2);
  const router = useRouter();
  const { trx } = router.query;
  const { t } = useTranslation();
  return (
    <>
      <Head>
        <title>
          {t('Checkout')}
        </title>
      </Head>
      <div className="checkout-cont">
        <div className="container">
          <div className="row">
            <div className="col">
              <div className="deposit-box basic-card">
                <TransactionSteps step={step} checkout />
                <CheckoutStep step={step} setStep={setStep} trx={trx} />
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default withAuth(Checkout);
