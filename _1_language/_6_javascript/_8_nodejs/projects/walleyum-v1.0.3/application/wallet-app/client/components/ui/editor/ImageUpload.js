/* eslint-disable jsx-a11y/label-has-associated-control */
import React, { useRef, useState } from 'react';
import Spinner from 'react-bootstrap/Spinner';
import { imageUpload } from '../../../lib/settingsUpdate';

const ImageUpload = ({ input }) => {
  const { name, label } = input;
  const inputEl = useRef(null);
  const [actionLoader, setActionLoader] = useState(false);
  const handleUpload = async (e) => {
    const data = await imageUpload(e.target.files[0], setActionLoader);
    inputEl.current.value = data.path;
  };
  return (
    <div className="single-profile" key={name}>
      <label>{label}</label>
      <div className="image-field">
        <input
          type="text"
          name={name}
          placeholder="Image URL"
          disabled
          ref={inputEl}
        />
        <label className="bttn-mid btn-ylo" htmlFor={`inputFile-${name}`}>
          {actionLoader ? <Spinner animation="border" role="status" size="sm" /> : 'Upload'}
        </label>
        <input id={`inputFile-${name}`} className="inputFile" type="file" onChange={handleUpload} />
      </div>
    </div>
  );
};

export default ImageUpload;
