/* eslint-disable jsx-a11y/label-has-associated-control */
import Head from 'next/head';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';
import { Spinner } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle } from 'react-icons/bi';
import ImageUploading from 'react-images-uploading';
import Toggle from 'react-toggle';
import EditorHeader from '../../../../components/admin/EditorHeader';
import Loader from '../../../../components/Loader';
import SidebarAdmin from '../../../../components/SidebarAdmin';
import UserHeaderAdmin from '../../../../components/UserHeaderAdmin';
import { useCurrencyById } from '../../../../data/useCurrency';
import withAuthAdmin from '../../../../hoc/withAuthAdmin';
import { currencyUpdate } from '../../../../lib/currencyRequest';

const WalletEdit = ({ userData, settings }) => {
  const router = useRouter();
  const { id } = router.query;

  const [actionLoader, setActionLoader] = useState(false);
  const [icon, setIcon] = useState();
  const [active, setActive] = useState(true);

  const { data, loading } = useCurrencyById(id);

  const { t } = useTranslation();

  useEffect(() => {
    setActive(data?.active);
  }, [data]);

  const handleUpdate = () => {
    const params = {
      active,
      icon
    };
    currencyUpdate(id, params, setActionLoader);
  };

  if (loading) {
    return <Loader height="100vh" />;
  }
  return (
    <>
      <Head>
        <title>
          {t('Admin Panel')}
        </title>
      </Head>
      <UserHeaderAdmin />
      <SidebarAdmin userData={userData} settings={settings} />
      <div className="content-body">
        <div className="container-fluid">
          <EditorHeader name={`${t('Currency')} #${id}`} />
          <div className="row">
            <div className="col-12 col-xl-6">
              <div className="basic-card">
                <div className="settings-box">
                  <div className="single-profile">
                    <label>{t('Icon')}</label>
                    <ImageUploading
                      value={icon}
                      onChange={setIcon}
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
                            className="kyc-upload-button new-btn"
                            style={isDragging ? { background: '#0d3b66' } : undefined}
                            onClick={onImageUpload}
                            {...dragProps}
                          >
                            {imageList.length <= 0 ? `${t('Click or Drop Here')} (512x512px)` : (
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
                    <label>{t('Active')}</label>
                    <Toggle
                      checked={active}
                      onChange={(e) => setActive(e.target.checked)}
                      name="active"
                    />
                  </div>
                  <button
                    onClick={() => handleUpdate()}
                    type="submit"
                    className="bttn-mid btn-ylo"
                    disabled={actionLoader}
                  >
                    {actionLoader ? (
                      <>
                        <Spinner animation="border" role="status" size="sm" />
                        {' '}
                        {t('Add Wallet')}
                      </>
                    ) : (
                      <>
                        <BiErrorCircle />
                        {t('Add Wallet')}
                      </>
                    )}
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default withAuthAdmin(WalletEdit);
