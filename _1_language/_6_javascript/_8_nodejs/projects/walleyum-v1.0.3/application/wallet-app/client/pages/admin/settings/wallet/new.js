/* eslint-disable jsx-a11y/label-has-associated-control */
import Head from 'next/head';
import React, { useEffect, useState } from 'react';
import { Dropdown, Spinner } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle } from 'react-icons/bi';
import ImageUploading from 'react-images-uploading';
import Toggle from 'react-toggle';
import EditorHeader from '../../../../components/admin/EditorHeader';
import Loader from '../../../../components/Loader';
import SidebarAdmin from '../../../../components/SidebarAdmin';
import UserHeaderAdmin from '../../../../components/UserHeaderAdmin';
import { useCurrencyList } from '../../../../data/useCurrency';
import withAuthAdmin from '../../../../hoc/withAuthAdmin';
import { currencyAdd } from '../../../../lib/currencyRequest';

const WalletNew = ({ userData, settings }) => {
  const [actionLoader, setActionLoader] = useState(false);
  const [icon, setIcon] = useState();
  const [currency, setCurrency] = useState();
  const [crypto, setCrypto] = useState(false);
  const [active, setActive] = useState(true);

  const { data, loading } = useCurrencyList();

  const { t } = useTranslation();

  useEffect(() => {
    setCurrency(data?.[0]);
    setCrypto(data?.[0]?.crypto);
  }, [data]);

  useEffect(() => {
    setCrypto(currency?.crypto);
  }, [currency]);

  const handleUpdate = () => {
    const params = {
      name: currency?.name,
      symbol: currency?.symbol,
      crypto,
      active,
      rateUsd: 1,
      icon
    };
    currencyAdd(params, setActionLoader);
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
          <EditorHeader name="New Wallet" />
          <div className="row">
            <div className="col-12 col-xl-6">
              <div className="basic-card">
                <div className="settings-box">
                  <label htmlFor="curSelector">
                    {t('Currency')}
                  </label>
                  <Dropdown id="curSelector" className="mt-10 mb-20">
                    <Dropdown.Toggle className="bttn-small btn-emt dropdown-method" variant="link">
                      {currency?.symbol || t('Please Select')}
                    </Dropdown.Toggle>
                    <Dropdown.Menu>
                      {data?.map((wallet) => (
                        <Dropdown.Item
                          key={wallet?.symbol}
                          onClick={() => setCurrency(wallet)}
                        >
                          {wallet?.symbol}
                        </Dropdown.Item>
                      ))}
                    </Dropdown.Menu>
                  </Dropdown>
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
                    <label>{t('Crypto')}</label>
                    <Toggle
                      checked={crypto}
                      onChange={(e) => setCrypto(e.target.checked)}
                    />
                  </div>
                  <div className="single-profile">
                    <label>{t('Active')}</label>
                    <Toggle
                      checked={active}
                      onChange={(e) => setActive(e.target.checked)}
                    />
                  </div>
                  <button onClick={() => handleUpdate()} type="button" className="bttn-mid btn-ylo" disabled={actionLoader}>
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

export default withAuthAdmin(WalletNew);
