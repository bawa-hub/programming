/* eslint-disable max-len */
/* eslint-disable jsx-a11y/label-has-associated-control */
import cogoToast from 'cogo-toast';
import React, { useEffect, useState } from 'react';
import { Dropdown, Image, Table } from 'react-bootstrap';
import Spinner from 'react-bootstrap/Spinner';
import { useTranslation } from 'react-i18next';
import {
  BiCheckCircle, BiLeftArrowAlt, BiRightArrowAlt, BiXCircle
} from 'react-icons/bi';
import useCurrency from '../../data/useCurrency';
import useGateways, { useGatewayCurrencies } from '../../data/useGateways';
import useWallet from '../../data/useWallet';
import depositRequest from '../../lib/depositRequest';
import Loader from '../Loader';
import GatewayLogo from './GatewayLogo';

const DepositStep = ({ step, setStep, status }) => {
  const [selectedCurrency, setSelectedCurrency] = useState();
  const [currentBalance, setCurrentBalance] = useState(0);
  const [amount, setAmount] = useState();
  const [payment, setPayment] = useState();
  const [actionLoader, setActionLoader] = useState(false);
  const { data, loading } = useCurrency();
  const { data: walletData, loading: walletLoading } = useWallet();
  const { data: gatewayData, loading: gatewayLoading } = useGateways();
  const { data: currencyData, loading: currencyLoading } = useGatewayCurrencies();
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

  const handleDeposit = () => {
    if (payment) {
      depositRequest({
        payment_method: payment,
        amount: parseFloat(amount, 10),
        currency: selectedCurrency?.symbol
      }, setActionLoader);
    } else {
      cogoToast.error(t('Please select a payment method'), { position: 'bottom-center' });
    }
  };

  if (loading || walletLoading || gatewayLoading || currencyLoading) {
    return <Loader height="200px" />;
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
              defaultValue={amount}
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
                <td>{t('Amount')}</td>
                <td style={{ color: 'white', fontWeight: 'bold' }}>
                  {amount}
                  {' '}
                  {selectedCurrency?.symbol}
                </td>
              </tr>
              <tr>
                <td>{t('Total')}</td>
                <td style={{ color: 'white', fontWeight: 'bold' }}>
                  {amount}
                  {' '}
                  {selectedCurrency?.symbol}
                </td>
              </tr>
            </tbody>
          </Table>
          <div className="payment-method">
            <h4>{t('Payment Method')}</h4>
            {gatewayData?.map((gateway) => {
              const isCurrencyAvailable = currencyData[gateway?.value].supportedCurrencies.some((cur) => cur === selectedCurrency?.symbol);
              if (isCurrencyAvailable) {
                return (
                  <button
                    type="button"
                    onClick={() => setPayment(gateway.value)}
                    className={`gateway ${(payment === gateway.value) ? 'active' : ''}`}
                    key={gateway?.value}
                  >
                    <GatewayLogo name={gateway.value} />
                  </button>
                );
              }
              return <React.Fragment key={gateway?.value} />;
            })}
          </div>
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
            onClick={() => handleDeposit()}
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
                {t('Make Payment')}
              </>
            )}
          </button>
        </div>
      </>
    );
  } if (step === 3) {
    return (
      <div className="transaction-success">
        {status === 'success' ? (
          <BiCheckCircle color="green" size={70} />
        ) : (
          <BiXCircle color="red" size={70} />)}
        <h2>
          {t('Deposit')}
          {' '}
          {status === 'success' ? 'Successful' : 'Failed'}
        </h2>
        {status === 'success' ? (
          <p>
            {t('Your requested amount has been added to your desired wallet')}
            .
          </p>
        ) : (
          <p>
            {t('Your deposit request declined by the payment gateway')}
            .
          </p>
        )}
        <button
          type="button"
          onClick={() => setStep(1)}
          className="bttn-mid btn-ylo"
        >
          {t('Make Another Deposit')}
        </button>
      </div>
    );
  }
  return <></>;
};

export default DepositStep;
