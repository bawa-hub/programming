import React from 'react';

const PageTitle = ({ data }) => (
  <div className="page-title">
    <h2>{data?.content}</h2>
  </div>
);

export default PageTitle;
