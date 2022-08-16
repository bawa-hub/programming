/* eslint-disable jsx-a11y/label-has-associated-control */
import React, { useState } from 'react';
import { Spinner } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle } from 'react-icons/bi';
import storeUpdate from '../../lib/storeUpdate';

const StoreSettings = ({ userData }) => {
  const [actionLoader, setActionLoader] = useState(false);
  const { t } = useTranslation();

  const handleUpdate = (e) => {
    e.preventDefault();
    const params = {
      name: e.target?.name?.value,
      address: e.target?.address?.value,
    };
    storeUpdate(params, setActionLoader);
  };

  return (
    <div className="basic-card">
      <h3 className="box-title">{t('Store Settings')}</h3>
      <div className="settings-box">
        <form onSubmit={handleUpdate}>
          <div className="single-profile">
            <label htmlFor="merId">{t('Merchant ID')}</label>
            <input id="merId" name="merId" type="text" disabled placeholder={t('Merchant ID')} defaultValue={userData?.merchant?.merId} />
          </div>
          <div className="single-profile">
            <label htmlFor="name">{t('Name')}</label>
            <input id="name" name="name" type="text" placeholder={t('Store Name')} defaultValue={userData?.merchant?.name} />
          </div>
          <div className="single-profile">
            <label htmlFor="emailField">{t('Email')}</label>
            <input
              id="emailField"
              name="email"
              disabled
              type="email"
              placeholder={t('Store Email Address')}
              defaultValue={userData?.merchant?.email}
            />
          </div>
          <div className="single-profile">
            <label htmlFor="addressField">{t('Address')}</label>
            <input
              id="addressField"
              name="address"
              type="text"
              placeholder={t('Store Address')}
              defaultValue={userData?.merchant?.address}
            />
          </div>
          <button type="submit" className="bttn-mid btn-ylo" disabled={actionLoader}>
            {actionLoader ? (
              <>
                <Spinner animation="border" role="status" size="sm" />
                {' '}
                {t('Update Store')}
              </>
            ) : (
              <>
                <BiErrorCircle />
                {t('Update Store')}
              </>
            )}
          </button>
        </form>
      </div>
    </div>
  );
};

export default StoreSettings;
