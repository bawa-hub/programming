import Head from 'next/head';
import React, { useState } from 'react';
import Alert from 'react-bootstrap/Alert';
import Spinner from 'react-bootstrap/Spinner';
import { useTranslation } from 'react-i18next';
import LanguageSwitch from '../components/LanguageSwitch';
import Loader from '../components/Loader';
import useSettings from '../data/useSettings';
import { forgotRequest } from '../lib/authRequest';

const ForgotPassword = () => {
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const { t } = useTranslation();

  const { data, loading: settingsLoading } = useSettings();

  const handleForget = (e) => {
    e.preventDefault();
    const { email } = e.target;
    const params = {
      email: email?.value,
    };
    forgotRequest(params, setLoading, setSuccess, email);
  };

  if (settingsLoading) {
    return <Loader height="100vh" />;
  }

  return (
    <>
      <Head>
        <title>
          {t('Forgot Password')}
          {' '}
          -
          {' '}
          {data?.site?.param1}
        </title>
        <link rel="icon" href={`${data?.apiUrl?.param1}/public/${data?.logo?.param2}`} />
      </Head>
      <div>
        <div className="container">
          <div className="row justify-content-center align-items-center vh-100">
            <div className="col-xl-4 col-lg-4 col-md-6 col-12">
              <div className="site-auth">
                <div className="lng-switch-cont">
                  <LanguageSwitch />
                </div>
                <div className="logo">
                  <a href="/"><img src={`${data?.apiUrl.param1}/public/${data?.logo?.param1}`} alt="sitename" /></a>
                </div>
                <h3>{t('Forget Password')}</h3>
                {success ? (
                  <Alert variant="success">
                    <Alert.Heading>{t('Reset Link Sent')}</Alert.Heading>
                    <p className="mb-0">
                      {t('A mail to reset your password has been sent your email address')}
                      .
                    </p>
                  </Alert>
                ) : (
                  <form onSubmit={handleForget}>
                    <input
                      className="box-input"
                      type="email"
                      name="email"
                      placeholder={t('Email Address')}
                      required
                    />
                    <button disabled={loading} type="submit" className="bttn-mid btn-ylo w-100">
                      {loading ? (
                        <Spinner animation="border" role="status" size="sm" />
                      ) : t('Reset Password')}
                    </button>
                  </form>
                )}
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default ForgotPassword;
