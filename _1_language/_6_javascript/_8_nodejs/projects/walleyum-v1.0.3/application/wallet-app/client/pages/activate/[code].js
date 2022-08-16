import { useRouter } from 'next/router';
import React, { useEffect } from 'react';
import { activationRequest } from '../../lib/authRequest';

const Activate = () => {
  const router = useRouter();
  const { code } = router.query;

  useEffect(() => {
    if (code) {
      activationRequest({ code });
    }
  }, [code]);

  return (
    <div>
      <p>
        Processing...
      </p>
    </div>
  );
};

export default Activate;
