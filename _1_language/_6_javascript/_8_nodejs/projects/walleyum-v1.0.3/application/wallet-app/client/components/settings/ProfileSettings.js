/* eslint-disable jsx-a11y/label-has-associated-control */
import React, { useState } from 'react';
import { Spinner } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle } from 'react-icons/bi';
import profileUpdate from '../../lib/profileUpdate';

const ProfileSettings = ({ userData }) => {
  const [actionLoader, setActionLoader] = useState(false);

  const { t } = useTranslation();

  const handleUpdate = (e) => {
    e.preventDefault();
    const params = {
      name: e.target?.fullName?.value,
      phone: e.target?.phone?.value,
      address: e.target?.address?.value,
      password: e.target?.password?.value || undefined,
    };
    profileUpdate(params, setActionLoader);
  };

  return (
    <div className="basic-card">
      <h3 className="box-title">{t('Account Settings')}</h3>
      <div className="settings-box">
        <form onSubmit={handleUpdate}>
          <div className="single-profile">
            <label htmlFor="fullName">{t('Full Name')}</label>
            <input id="fullName" name="fullName" type="text" placeholder={t('Your Full Name')} defaultValue={userData?.name} />
          </div>
          <div className="single-profile">
            <label htmlFor="phone">{t('Phone')}</label>
            <input id="phone" name="phone" type="text" placeholder={t('Your Phone Number')} defaultValue={userData?.phone} />
          </div>
          <div className="single-profile">
            <label htmlFor="address">{t('Address')}</label>
            <input id="address" name="address" type="text" placeholder={t('Your Address')} defaultValue={userData?.address} />
          </div>
          <div className="single-profile">
            <label htmlFor="email">{t('Email')}</label>
            <input id="email" name="email" type="email" placeholder={t('Your Email Address')} defaultValue={userData?.email} disabled />
          </div>
          <div className="single-profile">
            <label htmlFor="password">{t('Password')}</label>
            <input
              id="password"
              name="password"
              type="password"
              placeholder={t('Leave blank to keep unchanged')}
              autoComplete="wallet-password-change"
            />
          </div>
          <button type="submit" className="bttn-mid btn-ylo" disabled={actionLoader}>
            {actionLoader ? (
              <>
                <Spinner animation="border" role="status" size="sm" />
                {' '}
                {t('Update Profile')}
              </>
            ) : (
              <>
                <BiErrorCircle />
                {t('Update Profile')}
              </>
            )}
          </button>
        </form>
      </div>
    </div>
  );
};

export default ProfileSettings;
