/* eslint-disable jsx-a11y/label-has-associated-control */
import React, { useState } from 'react';
import { Spinner } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle } from 'react-icons/bi';
import { settingsUpdate } from '../../../lib/settingsUpdate';
import inputNumber from '../../../utils/inputNumber';

const ApiSettings = ({ settings }) => {
  const [actionLoader, setActionLoader] = useState(false);
  const { t } = useTranslation();

  const handleUpdate = (e) => {
    e.preventDefault();
    const params = {
      param1: e.target?.param1?.value,
    };
    settingsUpdate('freecurrencyapi', params, setActionLoader);
  };
  return (
    <>
      <div className="settings-box">
        <form onSubmit={handleUpdate}>
          <div className="single-profile">
            <label htmlFor="fcApi">{t('Free Currency API')}</label>
            <input
              id="fcApi"
              onKeyPress={inputNumber}
              name="param1"
              type="text"
              placeholder={t('Free Currency API Key')}
              defaultValue={settings?.freecurrencyapi?.param1}
            />
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

export default ApiSettings;
