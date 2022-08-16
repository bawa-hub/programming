import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { AiOutlinePlus } from 'react-icons/ai';
import useLinkeds from '../../data/useLinkeds';
import Loader from '../Loader';
import LinkAccount from './LinkAccount';
import LinkedCard from './LinkedCard';

const LinkedAcc = () => {
  const { data, loading } = useLinkeds();
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  const { t } = useTranslation();

  if (loading) {
    return <Loader height="200px" />;
  }
  return (
    <>
      <div className="basic-card">
        <h4 className="box-title">{t('Linked Accounts')}</h4>
        <div className="settings-box">
          {data?.data?.map((account) => <LinkedCard account={account} key={account?.id} />)}
          <button type="button" className="bttn-mid btn-ylo" onClick={handleShow}>
            <AiOutlinePlus />
            {t('Link New Account')}
          </button>
        </div>
      </div>
      <LinkAccount show={show} handleClose={handleClose} />
    </>
  );
};

export default LinkedAcc;
