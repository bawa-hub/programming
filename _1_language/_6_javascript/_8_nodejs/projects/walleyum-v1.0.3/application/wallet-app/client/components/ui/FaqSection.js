/* eslint-disable max-len */
import React from 'react';
import Faq from 'react-faq-component';

const FaqSection = ({ data, info }) => (
  <>
    <section className="section-padding bg-black2" id="faq">
      <div className="container">
        <div className="row justify-content-center">
          <div className="col-xl-5 col-md-12 col-12">
            <div className="section-title">
              <h2>{data?.sectionTitle}</h2>
              <p>{data?.sectionSub}</p>
              {data?.buttonLabel && (
              <a href={data?.buttonUrl} className="bttn-mid btn-ylo">
                {data?.buttonLabel}
              </a>
              )}
            </div>
          </div>
          <div className="col-xl-7 col-md-12 col-12">
            <div className="faq-area">
              <Faq
                data={{ rows: JSON.parse(info?.faq?.param1 || []) }}
              />
            </div>
          </div>
        </div>
      </div>
    </section>
  </>
);

export default FaqSection;
