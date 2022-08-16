/* eslint-disable max-len */
/* eslint-disable jsx-a11y/label-has-associated-control */
import React, { useEffect, useState } from 'react';
import { Dropdown, Image, Table } from 'react-bootstrap';
import Spinner from 'react-bootstrap/Spinner';
import { useTranslation } from 'react-i18next';
import {
  BiCheckCircle, BiLeftArrowAlt, BiRightArrowAlt
} from 'react-icons/bi';
import useCurrency from '../../data/useCurrency';
import useMethods from '../../data/useMethods';
import useWallet from '../../data/useWallet';
import withdrawRequest from '../../lib/withdrawRequest';
import Loader from '../Loader';

const WithdrawStep = ({ step, setStep }) => {
  const [selectedCurrency, setSelectedCurrency] = useState();
  const [method, setMethod] = useState();
  const [currentBalance, setCurrentBalance] = useState(0);
  const [amount, setAmount] = useState('');
  const [actionLoader, setActionLoader] = useState(false);
  const { data, loading } = useCurrency();
  const { data: walletData, loading: walletLoading } = useWallet();
  const { data: methods, loading: methodsLoading } = useMethods();
  const { t } = useTranslation();

  useEffect(() => {
    setSelectedCurrency(data?.data[0]);
  }, [data]);

  useEffect(() => {
    const walletFind = walletData?.find((wallet) => wallet.currency === selectedCurrency?.symbol);
    setCurrentBalance(walletFind?.balance);
    setMethod();
  }, [selectedCurrency, walletData]);

  const handleNext = (e) => {
    e.preventDefault();
    setStep(step + 1);
  };

  const handleSubmit = () => {
    withdrawRequest({
      methodId: method?.id,
      amount: parseFloat(amount, 10),
      currency: selectedCurrency?.symbol,
    }, setActionLoader, setStep);
  };

  if (loading || walletLoading || methodsLoading) {
    return <Loader />;
  }

  if (step === 1) {
    return (
      <>
        <form onSubmit={handleNext}>
          <div className="currency-amount">
            <label htmlFor="currencySelector">{t('Wallet')}</label>
            <Dropdown id="currencySelector">
              <Dropdown.Toggle className="bttn-small btn-emt" variant="link">
                <Image src={selectedCurrency?.icon} rounded />
                {selectedCurrency?.symbol}
              </Dropdown.Toggle>

              <Dropdown.Menu>
                {data?.data?.map((currency) => (
                  <Dropdown.Item
                    key={currency.id}
                    onClick={() => setSelectedCurrency(currency)}
                  >
                    <Image src={currency.icon} rounded />
                    {currency.symbol}
                  </Dropdown.Item>
                ))}
              </Dropdown.Menu>
              <p className="available-balance">
                {t('Available Balance')}
                :
                <span>
                  {' '}
                  {currentBalance}
                  {' '}
                  {selectedCurrency?.symbol}
                </span>
              </p>
            </Dropdown>
          </div>
          <div className="currency-amount">
            <label>{t('Amount')}</label>
            <input
              onChange={(e) => setAmount(e.target.value)}
              value={amount}
              type="text"
              required
            />
          </div>
          <div className="currency-amount">
            <label htmlFor="methodSelector">{t('Withdraw Method')}</label>
            <Dropdown id="methodSelector">
              <Dropdown.Toggle className="bttn-small btn-emt" variant="link">
                {method?.name || t('Please Select')}
              </Dropdown.Toggle>

              <Dropdown.Menu>
                {methods?.data?.map((methodData) => {
                  if (!(methodData?.currency === selectedCurrency?.symbol)) {
                    return null;
                  }
                  return (
                    <Dropdown.Item
                      key={methodData?.id}
                      onClick={() => setMethod(methodData)}
                    >
                      {methodData?.name}
                    </Dropdown.Item>
                  );
                })}
              </Dropdown.Menu>
              {method && (
              <p className="available-balance">
                {t('Fee')}
                :
                {' '}
                <span>
                  {`${method?.percentageCharge}% + ${method?.fixedCharge} ${method?.currency}`}
                </span>
              </p>
              )}
            </Dropdown>
          </div>
          <div className="bttns mt-30">
            <button
              type="submit"
              className="bttn-mid btn-ylo"
              disabled={!(selectedCurrency && method)}
            >
              <BiRightArrowAlt />
              {t('Next')}
            </button>
          </div>
        </form>
      </>
    );
  } if (step === 2) {
    return (
      <>
        <div className="transaction-review">
          <h4>{t('Review Details')}</h4>
          <Table striped hover responsive className="dark-color">
            <tbody>
              <tr>
                <td>{t('Withdraw Method')}</td>
                <td style={{ color: 'white', fontWeight: 'bold' }}>
                  {method?.name}
                </td>
              </tr>
              <tr>
                <td>{t('Amount')}</td>
                <td style={{ color: 'white', fontWeight: 'bold' }}>
                  {amount}
                  {' '}
                  {selectedCurrency?.symbol}
                </td>
              </tr>
              <tr>
                <td>
                  {t('Fee')}
                  {' '}
                  (
                  {`${method?.percentageCharge}% + ${method?.fixedCharge} ${method?.currency}`}
                  )
                </td>
                <td style={{ color: 'white', fontWeight: 'bold' }}>
                  {amount * (method?.percentageCharge / 100) + method?.fixedCharge}
                  {' '}
                  {selectedCurrency?.symbol}
                </td>
              </tr>
              <tr>
                <td>{t('You Get')}</td>
                <td style={{ color: 'white', fontWeight: 'bold' }}>
                  {amount - (amount * (method?.percentageCharge / 100) + method?.fixedCharge)}
                  {' '}
                  {selectedCurrency?.symbol}
                </td>
              </tr>
            </tbody>
          </Table>
        </div>
        <div className="bttns mt-30">
          <button
            type="button"
            onClick={() => setStep(step - 1)}
            className="bttn-mid btn-grey mr-10"
          >
            <BiLeftArrowAlt />
            {t('Back')}
          </button>
          <button
            type="button"
            onClick={() => handleSubmit()}
            className="bttn-mid btn-ylo"
            disabled={actionLoader}
          >
            {actionLoader ? (
              <>
                <Spinner animation="border" role="status" size="sm" />
                {' '}
                {t('Processing')}
              </>
            ) : (
              <>
                <BiRightArrowAlt />
                {t('Withdraw')}
              </>
            )}
          </button>
        </div>
      </>
    );
  } if (step === 3) {
    return (
      <div className="transaction-success">
        <BiCheckCircle color="green" size={70} />
        <h2>
          {t('Withdraw Request Submitted')}
        </h2>
        <p>
          {t('We will review your withdraw request and send the fund to your desired account')}
          ,
          {' '}
          {t('please allow upto 24 hours for us to review')}
          .
        </p>
        <button
          type="button"
          onClick={() => setStep(1)}
          className="bttn-mid btn-ylo"
        >
          {t('Make Another Withdraw')}
        </button>
      </div>
    );
  }
  return <></>;
};

export default WithdrawStep;
