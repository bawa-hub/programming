/* eslint-disable jsx-a11y/label-has-associated-control */
import React, { useState } from 'react';
import { Modal } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { FiEdit, FiGrid, FiTrash } from 'react-icons/fi';
import IconUpload from '../IconUpload';

export default function SortableService({
  dragging,
  dragged,
  children,
  handleDelete,
  handleEdit,
  ...rest
}) {
  const [active, setActive] = useState(false);
  const handleClose = () => setActive(false);
  const { t } = useTranslation();
  return (
    <>
      <div
        {...rest}
        className={`menu-builder ${dragged ? 'is-dragging' : ''} ${active ? 'active' : ''}`}
      >
        <div className="builder-content">
          <FiGrid />
          <p>{children?.title}</p>
        </div>
        <div className="d-flex">
          <button
            type="button"
            onClick={() => setActive(true)}
            className="action-btn"
          >
            <FiEdit />
          </button>
          <button
            type="button"
            onClick={() => handleDelete(children.id)}
            className="action-btn danger ml-10"
          >
            <FiTrash />
          </button>
        </div>
      </div>
      <Modal
        show={active}
        onHide={handleClose}
        size="lg"
        centered
        animation={false}
      >
        <form>
          <Modal.Header closeButton>
            <Modal.Title>
              {t('Edit Service')}
            </Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <div className="settings-box">
              <IconUpload id={children?.id} value={children?.icon} handleEdit={handleEdit} />
              <div className="single-profile">
                <label htmlFor="serTitle">{t('Title')}</label>
                <input
                  id="serTitle"
                  name="title"
                  type="text"
                  required
                  defaultValue={children?.title}
                  onChange={(e) => handleEdit(children.id, 'title', e.target.value)}
                />
              </div>
              <div className="single-profile">
                <label htmlFor="serContent">{t('Content')}</label>
                <input
                  id="serContent"
                  name="content"
                  type="text"
                  required
                  defaultValue={children?.content}
                  onChange={(e) => handleEdit(children.id, 'content', e.target.value)}
                />
              </div>
            </div>
          </Modal.Body>
          <Modal.Footer>
            <button type="button" className="bttn-mid btn-grey" onClick={handleClose}>
              {t('Close')}
            </button>
          </Modal.Footer>
        </form>
      </Modal>
    </>
  );
}
