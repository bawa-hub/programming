import React, { useState } from 'react';
import { Modal } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { FiMinus, FiPlus } from 'react-icons/fi';
import { balanceManage } from '../../../lib/settingsUpdate';

const BalanceManage = ({ currency, userId }) => {
  const [activeAdd, setActiveAdd] = useState(false);
  const [activeMinus, setActiveMinus] = useState(false);
  const [actionLoader, setActionLoader] = useState(false);
  const { t } = useTranslation();

  const handleAdd = (e) => {
    e.preventDefault();
    const params = {
      type: 'add',
      currency: currency?.currency,
      userId,
      amount: parseFloat(e.target?.amount?.value, 10),
    };
    balanceManage(params, setActionLoader);
    setActiveAdd(false);
  };
  const handleSub = (e) => {
    e.preventDefault();
    const params = {
      type: 'subtract',
      currency: currency?.currency,
      userId,
      amount: parseFloat(e.target?.amount?.value, 10),
    };
    balanceManage(params, setActionLoader);
    setActiveMinus(false);
  };

  return (
    <>
      <tr key={currency?.id}>
        <td>{currency?.currency}</td>
        <td>
          {currency?.balance || 0.00}
        </td>
        <td width="15%" align="center">
          <div className="d-flex">
            <button
              type="button"
              className="action-btn"
              onClick={() => setActiveAdd(!activeAdd)}
            >
              <FiPlus />
            </button>
            <button
              type="button"
              className="action-btn danger ml-10"
              onClick={() => setActiveMinus(!activeMinus)}
            >
              <FiMinus />
            </button>
          </div>
        </td>
      </tr>
      <Modal
        show={activeAdd}
        onHide={() => setActiveAdd(!activeAdd)}
        size="lg"
        centered
        animation={false}
      >
        <form onSubmit={handleAdd}>
          <Modal.Header closeButton>
            <Modal.Title>
              {t('Add Balance')}
            </Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <div className="settings-box">
              <div className="single-profile">
                <label htmlFor="amount">
                  {t('Add Balance')}
                  {' '}
                  {currency?.symbol}
                </label>
                <input
                  id="amount"
                  name="amount"
                  type="number"
                  required
                  step={0.000000000001}
                />
              </div>
            </div>
          </Modal.Body>
          <Modal.Footer>
            <button type="button" className="bttn-mid btn-grey" onClick={() => setActiveAdd(!activeAdd)}>
              {t('Close')}
            </button>
            <button
              type="submit"
              disabled={actionLoader}
              className="bttn-mid btn-ylo"
            >
              {t('Submit')}
            </button>
          </Modal.Footer>
        </form>
      </Modal>
      <Modal
        show={activeMinus}
        onHide={() => setActiveMinus(!activeMinus)}
        size="lg"
        centered
        animation={false}
      >
        <form onSubmit={handleSub}>
          <Modal.Header closeButton>
            <Modal.Title>
              {t('Subtract Balance')}
            </Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <div className="settings-box">
              <div className="single-profile">
                <label htmlFor="amount">
                  {t('Subtract Balance')}
                  {' '}
                  {currency?.symbol}
                </label>
                <input
                  id="amount"
                  name="amount"
                  type="number"
                  required
                  step={0.000000000001}
                />
              </div>
            </div>
          </Modal.Body>
          <Modal.Footer>
            <button type="button" className="bttn-mid btn-grey" onClick={() => setActiveMinus(!activeMinus)}>
              {t('Close')}
            </button>
            <button
              type="submit"
              disabled={actionLoader}
              className="bttn-mid btn-ylo"
            >
              {t('Submit')}
            </button>
          </Modal.Footer>
        </form>
      </Modal>
    </>
  );
};

export default BalanceManage;
