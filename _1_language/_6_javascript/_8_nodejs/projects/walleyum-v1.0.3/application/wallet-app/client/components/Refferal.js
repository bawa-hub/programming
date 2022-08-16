import cogoToast from 'cogo-toast';
import copy from 'copy-to-clipboard';
import React, { useRef } from 'react';
import { useTranslation } from 'react-i18next';
import { BiLink } from 'react-icons/bi';

const Refferal = ({ userData, settings }) => {
  const { t } = useTranslation();
  const inputRef = useRef(null);

  const copyToClipboard = () => {
    copy(`${settings?.appUrl?.param1}/register?refer=${userData?.id}`);
    cogoToast.success(t('Refferal Link Copied!'), { position: 'bottom-center' });
  };

  return (
    <>
      <div className="refferal-box">
        <h4 className="box-title">{t('Refferal Link')}</h4>
        <p>{t('Share this refferal link to your friends and earn money')}</p>
        <div className="refferal-form">
          <input
            type="text"
            disabled
            ref={inputRef}
            defaultValue={`${settings?.appUrl?.param1}/register?refer=${userData?.id}`}
            className="input-text input-box"
          />
          <button
            type="button"
            onClick={() => copyToClipboard()}
            className="input-box btn-blue"
          >
            <BiLink />
          </button>
        </div>
        <p className="mb-0 mt-10">
          {userData?.referCount}
          {' '}
          {t('user joined from your refferal')}
        </p>
      </div>
    </>
  );
};

export default Refferal;
