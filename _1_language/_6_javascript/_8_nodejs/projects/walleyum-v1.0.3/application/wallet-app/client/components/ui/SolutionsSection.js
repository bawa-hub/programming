/* eslint-disable max-len */
import React from 'react';

const SolutionsSection = ({ data, info }) => {
  const solutionData = JSON.parse(info?.solutions?.param1 || []);
  return (
    <>
      <section className="section-padding">
        <div className="container">
          <div className="row justify-content-center">
            <div className="col-xl-6 centered">
              <div className="section-title">
                <h2>{data?.sectionTitle}</h2>
                <p>{data?.sectionSub}</p>
              </div>
            </div>
          </div>
          <div className="row">
            {solutionData?.map((solution) => (
              <div key={solution?.id} className="col-xl-3 col-lg-3 col-md-6 col-sm-6 col-12">
                <div className="single-solutions">
                  <div className="solutions-img">
                    <img src={`/api/${solution?.icon}`} alt="Sol Icon" />
                  </div>
                  <div className="solutions-title">
                    <h3>{solution?.title}</h3>
                    <p>{solution?.content}</p>
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

export default SolutionsSection;
