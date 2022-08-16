/* eslint-disable max-len */
import React from 'react';

const ServicesSection = ({ data, info }) => {
  const serviceData = JSON.parse(info?.services?.param1 || []);
  return (
    <>
      <section className="features" id="features">
        <div className="container">
          <div className="row justify-content-center">
            <div className="col-xl-6 centered">
              <div className="section-title">
                <h2>{data?.sectionTitle}</h2>
                <p>{data?.sectionSub}</p>
              </div>
            </div>
          </div>
          <div className="row justify-content-center">
            {serviceData?.map((service) => (
              <div className="col-xl-3 col-lg-3 col-sm-6" key={service?.id}>
                <div className="single-features-box">
                  <div className="img">
                    <img src={`/api/${service?.icon}`} alt="Ser Icon" />
                  </div>
                  <div className="content">
                    <h3>{service?.title}</h3>
                    <p>{service?.content}</p>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>
    </>
  );
};

export default ServicesSection;
