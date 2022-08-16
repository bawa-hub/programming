import React from 'react';
import useWallet from '../data/useWallet';
import Loader from './Loader';

const Wallet = () => {
  const { data, loading } = useWallet();
  if (loading) {
    return <Loader height="250px" />;
  }
  return (
    <>
      {data?.map((wallet) => (
        <div className="single-coin" key={wallet?.currency}>
          <div className="coin-title">
            {wallet?.icon && (
            <img src={wallet?.icon} alt="" />
            )}
            <h5 className="d-inline-block">{wallet?.currency}</h5>
          </div>
          <h4>
            {wallet?.balance.toFixed(6)}
            {' '}
            {wallet?.currency}
          </h4>
        </div>
      ))}
    </>
  );
};

export default Wallet;
