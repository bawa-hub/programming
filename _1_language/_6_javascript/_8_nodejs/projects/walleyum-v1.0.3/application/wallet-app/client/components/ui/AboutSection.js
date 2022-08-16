/* eslint-disable max-len */
import React from 'react';

const AboutSection = ({ data }) => (
  <>
    <section className="section-padding">
      <div className="container">
        <div className="row align-items-center">
          <div className="col-xl-6 col-lg-6 col-sm-12">
            <div className="about-img">
              <img src={`/api/${data?.image}`} alt="About" />
            </div>
          </div>
          <div className="col-xl-6 col-lg-6 col-sm-12">
            <div className="about-content">
              <div className="section-title">
                <h2>{data?.sectionTitle}</h2>
              </div>
              <p>{data?.content}</p>
            </div>
          </div>
        </div>
      </div>
    </section>
  </>
);

export default AboutSection;
