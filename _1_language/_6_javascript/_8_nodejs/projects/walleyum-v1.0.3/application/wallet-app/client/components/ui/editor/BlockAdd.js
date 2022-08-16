/* eslint-disable jsx-a11y/label-has-associated-control */
import React, { useState } from 'react';
import { Modal } from 'react-bootstrap';
import { v1 as uuidv1 } from 'uuid';
import ImageUpload from './ImageUpload';

const BlockAdd = ({ field, handleAdd }) => {
  const { label } = field;
  const [active, setActive] = useState(false);

  const handleSubmit = (e) => {
    e.preventDefault();
    const data = field.fields.reduce(
      (obj, item) => Object.assign(obj, { [item.name]: e.target?.[item.name].value }), {}
    );
    handleAdd({
      id: uuidv1(),
      label,
      component: field.component,
      data
    });
    setActive(false);
  };

  return (
    <>
      <button className="block-add" type="button" onClick={() => setActive(true)}>
        {field.icon}
        {' '}
        {label}
      </button>
      <Modal
        show={active}
        onHide={() => setActive(false)}
        size="lg"
        centered
        animation={false}
      >
        <form onSubmit={handleSubmit}>
          <Modal.Header closeButton>
            <Modal.Title>
              Add
              {' '}
              {label}
            </Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <div className="settings-box">
              {field.fields.map((input) => {
                if (input.type === 'image') {
                  return <ImageUpload input={input} key={input.name} />;
                } if (input.type === 'textarea') {
                  return (
                    <div className="single-profile" key={input.name}>
                      <label>{input.label}</label>
                      <textarea
                        name={input.name}
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
                    <input type={input.type} name={input.name} />
                  </div>
                );
              })}
            </div>
          </Modal.Body>
          <Modal.Footer>
            <button type="button" className="bttn-mid btn-grey" onClick={() => setActive(false)}>
              Close
            </button>
            <button type="submit" className="bttn-mid btn-ylo">
              Add
            </button>
          </Modal.Footer>
        </form>
      </Modal>
    </>
  );
};

export default BlockAdd;
