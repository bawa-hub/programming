/* eslint-disable jsx-a11y/label-has-associated-control */
import React, { useEffect, useState } from 'react';
import { Dropdown, Image, Spinner } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle } from 'react-icons/bi';
import useCurrency from '../../data/useCurrency';
import requestPayment from '../../lib/requestPayment';

const RequestPayment = () => {
  const [selectedCurrency, setSelectedCurrency] = useState();
  const [actionLoader, setActionLoader] = useState(false);
  const { data } = useCurrency();
  const { t } = useTranslation();

  useEffect(() => {
    setSelectedCurrency(data?.data[0]);
  }, [data]);

  const handleSubmit = (e) => {
    e.preventDefault();
    const params = {
      email: e.target?.email?.value,
      amount: parseFloat(e.target?.amount?.value, 10),
      currency: e.target?.currency?.value,
    };
    requestPayment(params, setActionLoader);
  };

  return (
    <div className="basic-card">
      <h3 className="box-title">{t('Request Payment')}</h3>
      <div className="settings-box">
        <form onSubmit={handleSubmit}>
          <div className="single-profile">
            <label htmlFor="emailField">{t('Email')}</label>
            <input id="emailField" name="email" type="email" placeholder={t('Customer Email')} />
          </div>
          <div className="single-profile">
            <label htmlFor="amountField">{t('Amount')}</label>
            <input id="amountField" name="amount" type="text" placeholder={t('Amount')} />
          </div>
          <div className="currency-amount mb-20">
            <label htmlFor="currencySelector">{t('Currency')}</label>
            <Dropdown id="currencySelector">
              <Dropdown.Toggle className="bttn-small btn-emt dropdown-method mt-10" variant="link">
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
            </Dropdown>
          </div>
          <button type="submit" className="bttn-mid btn-ylo" disabled={actionLoader}>
            {actionLoader ? (
              <>
                <Spinner animation="border" role="status" size="sm" />
                {' '}
                {t('Request')}
              </>
            ) : (
              <>
                <BiErrorCircle />
                {t('Request')}
              </>
            )}
          </button>
        </form>
      </div>
    </div>
  );
};

export default RequestPayment;
