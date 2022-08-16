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
import useWallet from '../../data/useWallet';
import paymentRequest from '../../lib/paymentRequest';
import Loader from '../Loader';

const PaymentStep = ({ step, setStep }) => {
  const [selectedCurrency, setSelectedCurrency] = useState();
  const [currentBalance, setCurrentBalance] = useState(0);
  const [amount, setAmount] = useState('');
  const [merchant, setMerchant] = useState('');
  const [actionLoader, setActionLoader] = useState(false);
  const { data, loading } = useCurrency();
  const { data: walletData, loading: walletLoading } = useWallet();
  const { t } = useTranslation();

  useEffect(() => {
    setSelectedCurrency(data?.data[0]);
  }, [data]);

  useEffect(() => {
    const walletFind = walletData?.find((wallet) => wallet.currency === selectedCurrency?.symbol);
    setCurrentBalance(walletFind?.balance);
  }, [selectedCurrency, walletData]);

  const handleNext = (e) => {
    e.preventDefault();
    setStep(step + 1);
  };

  const handleSubmit = () => {
    paymentRequest({
      merchantId: merchant,
      amount: parseFloat(amount, 10),
      currency: selectedCurrency?.symbol,
    }, setActionLoader, setStep);
  };

  if (loading || walletLoading) {
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
            <label>{t('Merchant ID')}</label>
            <input
              onChange={(e) => setMerchant(e.target.value)}
              value={merchant}
              type="text"
              required
            />
          </div>

          <div className="bttns mt-30">
            <button
              type="submit"
              className="bttn-mid btn-ylo"
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
                <td>{t('Merchant ID')}</td>
                <td style={{ color: 'white', fontWeight: 'bold' }}>
                  {merchant}
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
                {t('Pay')}
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
          {t('Payment Successful')}
        </h2>
        <p>
          {amount}
          {' '}
          {selectedCurrency?.symbol}
          {' '}
          {t('sent to')}
          {' '}
          {merchant}
        </p>
        <button
          type="button"
          onClick={() => setStep(1)}
          className="bttn-mid btn-ylo"
        >
          {t('Make Another Payment')}
        </button>
      </div>
    );
  }
  return <></>;
};

export default PaymentStep;
