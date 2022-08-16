/* eslint-disable jsx-a11y/label-has-associated-control */
import React, { useState } from 'react';
import { Modal } from 'react-bootstrap';
import { FiEdit, FiGrid, FiTrash } from 'react-icons/fi';
import fields from '../../ui/editor/ComponentsField';
import ImageUploadEdit from '../../ui/editor/ImageUploadEdit';

export default function SortableSection({
  dragging,
  dragged,
  children,
  handleDelete,
  handleEdit,
  ...rest
}) {
  const [active, setActive] = useState(false);
  const handleClose = () => setActive(false);

  const selectedField = fields.find((field) => field?.component === children?.component);

  return (
    <>
      <div
        {...rest}
        className={`menu-builder ${dragged ? 'is-dragging' : ''} ${active ? 'active' : ''}`}
      >
        <div className="builder-content">
          <FiGrid />
          <p>{children.label}</p>
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
              Edit Section
            </Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <div className="settings-box">
              {selectedField?.fields.map((input) => {
                if (input.type === 'image') {
                  return (
                    <ImageUploadEdit
                      input={input}
                      key={input.name}
                      defaultValue={children?.data?.[input.name]}
                      handleEdit={handleEdit}
                      childrenId={children?.id}
                      inputName={input?.name}
                    />
                  );
                } if (input.type === 'textarea') {
                  return (
                    <div className="single-profile" key={input.name}>
                      <label>{input.label}</label>
                      <textarea
                        name={input.name}
                        defaultValue={children?.data?.[input.name]}
                        onChange={(e) => handleEdit(children.id, input.name, e.target.value)}
                        className="input-box"
                        row={10}
                        style={{
                          width: '100%', height: '100px', paddingTop: '10px', lineHeight: '1.5'
                        }}
                        required
                      />
                    </div>
                  );
                }
                return (
                  <div className="single-profile" key={input.name}>
                    <label>{input.label}</label>
                    <input
                      type={input.type}
                      name={input.name}
                      defaultValue={children?.data?.[input.name]}
                      onChange={(e) => handleEdit(children.id, input.name, e.target.value)}
                    />
                  </div>
                );
              })}
            </div>
          </Modal.Body>
          <Modal.Footer>
            <button type="button" className="bttn-mid btn-grey" onClick={handleClose}>
              Close
            </button>
          </Modal.Footer>
        </form>
      </Modal>
    </>
  );
}
