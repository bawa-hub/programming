/* eslint-disable jsx-a11y/label-has-associated-control */
import React, { useState } from 'react';
import Spinner from 'react-bootstrap/Spinner';
import { useTranslation } from 'react-i18next';
import { imageUpload } from '../../../lib/settingsUpdate';

const IconUpload = ({ id, value, handleEdit }) => {
  const [actionLoader, setActionLoader] = useState(false);
  const handleUpload = async (e) => {
    const data = await imageUpload(e.target.files[0], setActionLoader);
    handleEdit(id, 'icon', data.path);
  };
  const { t } = useTranslation();
  return (
    <div className="single-profile">
      <label>{t('Icon')}</label>
      <div className="image-field">
        {value && (
        <div className="image-preview-icon">
          <img src={`/api/${value}`} alt="Icon" />
        </div>
        )}
        <input type="text" defaultValue={value} placeholder={t('Icon URL')} disabled />
        <label className="bttn-mid btn-ylo" htmlFor="inputFile">
          {actionLoader ? <Spinner animation="border" role="status" size="sm" /> : t('Upload')}
        </label>
        <input id="inputFile" className="inputFile" type="file" onChange={handleUpload} />
      </div>
    </div>
  );
};

export default IconUpload;
