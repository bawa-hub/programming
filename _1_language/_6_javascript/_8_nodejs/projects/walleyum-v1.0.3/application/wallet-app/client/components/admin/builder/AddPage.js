/* eslint-disable jsx-a11y/label-has-associated-control */
import React from 'react';
import { Modal } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import pageAdd from '../../../lib/pageUpdate';

const AddPage = ({ active, setActive }) => {
  const handleClose = () => setActive(false);
  const { t } = useTranslation();

  const handleSubmit = (e) => {
    e.preventDefault();
    const slug = e.target?.slug?.value.replace(/[^\w\s]/gi, '');
    const params = {
      name: e.target?.name?.value,
      slug,
      type: 'landing',
      content: []
    };
    pageAdd(params);
    handleClose();
  };

  return (
    <Modal
      show={active}
      onHide={handleClose}
      size="lg"
      centered
      animation={false}
    >
      <form onSubmit={handleSubmit}>
        <Modal.Header closeButton>
          <Modal.Title>
            {t('Add New Page')}
          </Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <div className="settings-box">
            <div className="single-profile">
              <label htmlFor="pageName">{t('Name')}</label>
              <input
                id="pageName"
                name="name"
                type="text"
                required
              />
            </div>
            <div className="single-profile">
              <label htmlFor="pageSlug">{t('Slug')}</label>
              <input
                id="pageSlug"
                name="slug"
                type="text"
                required
              />
            </div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <button type="button" className="bttn-mid btn-grey" onClick={handleClose}>
            {t('Close')}
          </button>
          <button type="submit" className="bttn-mid btn-ylo">
            {t('Submit')}
          </button>
        </Modal.Footer>
      </form>
    </Modal>
  );
};

export default AddPage;
