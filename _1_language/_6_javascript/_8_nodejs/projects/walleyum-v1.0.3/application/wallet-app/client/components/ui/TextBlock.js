/* eslint-disable jsx-a11y/label-has-associated-control */
import React from 'react';

const TextBlock = ({ data }) => (
  <div className="container mt-20 mb-10">
    <p>{data?.content}</p>
  </div>
);

export default TextBlock;
