/* eslint-disable jsx-a11y/label-has-associated-control */
import React, { useState } from 'react';
import { Spinner } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle } from 'react-icons/bi';
import ImageUploading from 'react-images-uploading';
import { logoUpdate } from '../../../lib/settingsUpdate';

const LogoFavicon = () => {
  const [logo, setLogo] = useState();
  const [favicon, setFavicon] = useState();
  const [actionLoader, setActionLoader] = useState(false);
  const { t } = useTranslation();
  const handleUpdate = (e) => {
    e.preventDefault();
    const params = {
      logo, favicon
    };
    logoUpdate(params, setActionLoader);
  };
  return (
    <>
      <div className="settings-box">
        <form onSubmit={handleUpdate}>
          <div className="single-profile">
            <label>{t('Logo')}</label>
            <ImageUploading
              value={logo}
              onChange={setLogo}
              acceptType={['jpg', 'png', 'jpeg']}
            >
              {({
                imageList,
                onImageUpload,
                isDragging,
                dragProps,
              }) => (
                // write your building UI
                <div className="upload__image-wrapper">
                  <button
                    type="button"
                    className="kyc-upload-button-selfie"
                    style={isDragging ? { background: '#0d3b66' } : undefined}
                    onClick={onImageUpload}
                    {...dragProps}
                  >
                    {imageList.length <= 0 ? t('Click or Drop Here (Logo)') : (
                      imageList.map((image) => (
                        <div key={image.dataURL} className="image-item">
                          <img src={image.dataURL} alt="" width="100%" />
                        </div>
                      ))
                    )}
                  </button>
                </div>
              )}
            </ImageUploading>
          </div>
          <div className="single-profile">
            <label>{t('Favicon')}</label>
            <ImageUploading
              value={favicon}
              onChange={setFavicon}
              acceptType={['jpg', 'png', 'jpeg']}
            >
              {({
                imageList,
                onImageUpload,
                isDragging,
                dragProps,
              }) => (
                // write your building UI
                <div className="upload__image-wrapper">
                  <button
                    type="button"
                    className="kyc-upload-button-selfie"
                    style={isDragging ? { background: '#0d3b66' } : undefined}
                    onClick={onImageUpload}
                    {...dragProps}
                  >
                    {imageList.length <= 0 ? t('Click or Drop Here (Favicon)') : (
                      imageList.map((image) => (
                        <div key={image.dataURL} className="image-item">
                          <img src={image.dataURL} alt="" width="100%" />
                        </div>
                      ))
                    )}
                  </button>
                </div>
              )}
            </ImageUploading>
          </div>
          <button type="submit" className="bttn-mid btn-ylo" disabled={actionLoader}>
            {actionLoader ? (
              <>
                <Spinner animation="border" role="status" size="sm" />
                {' '}
                {t('Update')}
              </>
            ) : (
              <>
                <BiErrorCircle />
                {t('Update')}
              </>
            )}
          </button>
        </form>
      </div>
    </>
  );
};

export default LogoFavicon;
