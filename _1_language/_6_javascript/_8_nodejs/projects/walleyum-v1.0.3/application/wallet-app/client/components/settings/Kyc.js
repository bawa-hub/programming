/* eslint-disable jsx-a11y/label-has-associated-control */
import cogoToast from 'cogo-toast';
import React, { useState } from 'react';
import { Dropdown, Spinner } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle } from 'react-icons/bi';
import ImageUploading from 'react-images-uploading';
import useKyc from '../../data/useKyc';
import kycUpdate, { kycResubmit } from '../../lib/kycUpdate';
import Loader from '../Loader';

const Kyc = () => {
  const [type, setType] = useState('nid');
  const [front, setFront] = useState();
  const [back, setBack] = useState();
  const [selfie, setSelfie] = useState();
  const [actionLoader, setActionLoader] = useState(false);
  const { t } = useTranslation();

  const { data, loading } = useKyc();

  if (loading) {
    return <Loader height="300px" />;
  }

  const typePicker = (typeDoc) => {
    if (typeDoc === 'nid') {
      return t('National Identity Card');
    } if (typeDoc === 'passport') {
      return t('Passport');
    } if (typeDoc === 'driving') {
      return t('Driving License');
    }
    return null;
  };

  const handleSubmit = () => {
    if (!front) {
      return cogoToast.error(t('Please upload front side'), { position: 'bottom-center' });
    } if (!selfie) {
      return cogoToast.error(t('Please upload your selfie'), { position: 'bottom-center' });
    } if (!(type === 'passport') && !back) {
      return cogoToast.error(t('Please upload back side'), { position: 'bottom-center' });
    }
    kycUpdate({
      type, front, back, selfie
    }, setActionLoader);
    return null;
  };

  const handleResubmit = () => {
    kycResubmit(setActionLoader);
    return null;
  };

  return (
    <div className="basic-card">
      <h4 className="box-title">{t('KYC')}</h4>
      <div className="settings-box">
        <p>
          {t('Status')}
          :
          {' '}
          <strong
            className={`status ${data?.status === 'declined' ? 'danger' : 'success'}`}
            style={{ textTransform: 'capitalize' }}
          >
            {data?.status || t('Awaiting Submission')}
          </strong>
        </p>
        {!(data?.status) ? (
          <>
            <div className="single-profile">
              <label htmlFor="documentSelector">{t('Document Type')}</label>
              <Dropdown id="documentSelector">
                <Dropdown.Toggle className="bttn-small btn-emt dropdown-kyc" variant="link">
                  {typePicker(type)}
                </Dropdown.Toggle>
                <Dropdown.Menu>
                  <Dropdown.Item onClick={() => setType('nid')}>
                    {t('National Identity Card')}
                  </Dropdown.Item>
                  <Dropdown.Item onClick={() => setType('passport')}>
                    {t('Passport')}
                  </Dropdown.Item>
                  <Dropdown.Item onClick={() => setType('driving')}>
                    {t('Driving License')}
                  </Dropdown.Item>
                </Dropdown.Menu>
              </Dropdown>
            </div>
            <div className="single-profile">
              <label>{t('Front Side')}</label>
              <ImageUploading
                value={front}
                onChange={setFront}
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
                      className="kyc-upload-button"
                      style={isDragging ? { background: '#0d3b66' } : undefined}
                      onClick={onImageUpload}
                      {...dragProps}
                    >
                      {imageList.length <= 0 ? t('Click or Drop Here (Front Side)') : (
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
            {!(type === 'passport') ? (
              <div className="single-profile">
                <label>{t('Back Side')}</label>
                <ImageUploading
                  value={back}
                  onChange={setBack}
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
                        className="kyc-upload-button"
                        style={isDragging ? { background: '#0d3b66' } : undefined}
                        onClick={onImageUpload}
                        {...dragProps}
                      >
                        {imageList.length <= 0 ? t('Click or Drop Here (Back Side)') : (
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
            ) : (<></>)}
            <div className="single-profile">
              <label>{t('Selfie')}</label>
              <ImageUploading
                value={selfie}
                onChange={setSelfie}
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
                      {imageList.length <= 0 ? t('Click or Drop Here (Selfie)') : (
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
            <button onClick={() => handleSubmit()} type="button" className="bttn-mid btn-ylo" disabled={actionLoader}>
              {actionLoader ? (
                <>
                  <Spinner animation="border" role="status" size="sm" />
                  {' '}
                  {t('Submit KYC')}
                </>
              ) : (
                <>
                  <BiErrorCircle />
                  {t('Submit KYC')}
                </>
              )}
            </button>
          </>
        ) : (
          <>
            {data?.status === 'declined' ? (
              <>
                <p>
                  {t('Your submitted document is not valid')}
                  .
                  {' '}
                  {t('Please resubmit all of your documents')}
                  .
                </p>
                <button onClick={() => handleResubmit()} type="button" className="bttn-mid btn-ylo" disabled={actionLoader}>
                  {actionLoader ? (
                    <>
                      <Spinner animation="border" role="status" size="sm" />
                      {' '}
                      {t('Resubmit KYC')}
                    </>
                  ) : (
                    <>
                      <BiErrorCircle />
                      {t('Resubmit KYC')}
                    </>
                  )}
                </button>
              </>
            ) : (
              <p>
                {t('Thanks for submitting your KYC')}
                .
                {' '}
                {t('We will look into your documents and approve your application')}
                .
              </p>
            )}
          </>
        )}
      </div>
    </div>
  );
};

export default Kyc;
