/* eslint-disable jsx-a11y/label-has-associated-control */
import React, { useState } from 'react';
import { Dropdown, Spinner } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle } from 'react-icons/bi';
import useCurrency from '../../../data/useCurrency';
import { settingsUpdate } from '../../../lib/settingsUpdate';
import inputNumber from '../../../utils/inputNumber';

const RefferalSettings = ({ settings }) => {
  const [actionLoader, setActionLoader] = useState(false);
  const [currency, setCurrency] = useState(settings?.refferal?.param2);

  const { t } = useTranslation();

  const { data } = useCurrency();

  const handleUpdate = (e) => {
    e.preventDefault();
    const params = {
      param1: e.target?.param1?.value,
      param2: currency
    };
    settingsUpdate('refferal', params, setActionLoader);
  };
  return (
    <>
      <div className="settings-box">
        <form onSubmit={handleUpdate}>
          <div className="single-profile">
            <label htmlFor="rewardAmount">{t('Reward Amount')}</label>
            <input
              id="rewardAmount"
              onKeyPress={inputNumber}
              name="param1"
              type="text"
              placeholder={t('Reward Amount')}
              defaultValue={settings?.refferal?.param1}
            />
          </div>
          <div className="single-profile">
            <label htmlFor="curSelector">
              {t('Currency')}
            </label>
            <Dropdown id="curSelector" className="mt-10 mb-20">
              <Dropdown.Toggle className="bttn-small btn-emt dropdown-method" variant="link">
                {currency || t('Please Select')}
              </Dropdown.Toggle>
              <Dropdown.Menu>
                {data?.data?.map((wallet) => (
                  <Dropdown.Item
                    key={wallet?.symbol}
                    onClick={() => setCurrency(wallet?.symbol)}
                  >
                    {wallet?.symbol}
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
                {t('Update Settings')}
              </>
            ) : (
              <>
                <BiErrorCircle />
                {t('Update Settings')}
              </>
            )}
          </button>
        </form>
      </div>
    </>
  );
};

export default RefferalSettings;
