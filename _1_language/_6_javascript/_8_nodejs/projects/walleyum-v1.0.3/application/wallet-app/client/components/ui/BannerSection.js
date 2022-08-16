/* eslint-disable max-len */
import Link from 'next/link';
import React from 'react';

const BannerSection = ({ data }) => (
  <>
    <section className="banner blue-overlay-2" style={{ backgroundImage: `url(${`/api/${data?.cover}`})` }}>
      <div className="container">
        <div className="row">
          <div className="col-xl-7 col-md-7 col-12 d-flex align-items-center">
            <div className="banner-content">
              <h2>{data?.sectionTitle}</h2>
              <p>{data?.content}</p>
              {data?.buttonLabel && (
              <Link href={data?.buttonUrl || '/'}>
                <a className="bttn-mid btn-ylo mr-10">
                  {data?.buttonLabel}
                </a>
              </Link>
              )}
            </div>
          </div>
          <div className="col-xl-5 col-md-5 col-12">
            <div className="banner-content img">
              <img src={`/api/${data?.side}`} alt="Side" />
            </div>
          </div>
        </div>
      </div>
    </section>
  </>
);

export default BannerSection;
