import React from 'react';
import Spinner from 'react-bootstrap/Spinner';

const Loader = ({ height }) => (
  <div style={{
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    height
  }}
  >
    <Spinner animation="border" role="status" size="lg" />
  </div>
);

export default Loader;
