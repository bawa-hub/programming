/* eslint-disable jsx-a11y/label-has-associated-control */
import React from 'react';

const HeadingBlock = ({ data }) => (
  <div className="container mt-20 mb-10">
    <h3>{data?.content}</h3>
  </div>
);

export default HeadingBlock;
