import React from 'react';
import { useTranslation } from 'react-i18next';
import { BiCheck } from 'react-icons/bi';

const TransactionSteps = ({ step, checkout }) => {
  const { t } = useTranslation();
  return (
    <>
      <div className="transaction-steps">
        {!checkout && (
        <div className={`single-step ${(step === 1) ? 'active' : null}`}>
          {step === 1 ? <span /> : <span><BiCheck /></span>}
          {t('Amount')}
        </div>
        )}
        <div className={`single-step ${(step === 2) ? 'active' : null}`}>
          {(step === 1 || step === 2) ? <span /> : <span><BiCheck /></span>}
          {t('Review')}
        </div>
        <div className={`single-step ${(step === 3) ? 'active' : null}`}>
          {(step === 3) ? <span><BiCheck /></span> : <span />}
          {t('Success')}
        </div>
      </div>
    </>
  );
};

export default TransactionSteps;
