/* eslint-disable jsx-a11y/label-has-associated-control */
import React from 'react';
import {
  BiEnvelope, BiMap, BiPaperPlane, BiPhone
} from 'react-icons/bi';

const ContactSection = ({ data }) => (
  <>
    <section className="section-padding" id="contact">
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
          <div className="col-xl-4 col-md-4 col-sm-6 col-12">
            <div className="single-contact-box">
              <BiEnvelope />
              <h4>Contact Emails</h4>
              <p>{data?.email}</p>
            </div>
          </div>
          <div className="col-xl-4 col-md-4 col-sm-6 col-12">
            <div className="single-contact-box">
              <BiPhone />
              <h4>Contact Number</h4>
              <p>{data?.phone}</p>
            </div>
          </div>
          <div className="col-xl-4 col-md-4 col-sm-6 col-12">
            <div className="single-contact-box">
              <BiMap />
              <h4>Contact Address</h4>
              <p>{data?.address}</p>
            </div>
          </div>
        </div>
        <div className="row justify-content-center">
          <div className="col-xl-12 col-lg-12 col-md-12 col-sm-12">
            <div className="site-form mt-30">
              <form action="#">
                <div className="row">
                  <div className="col-xl-6 col-lg-6 col-sm-12">
                    <input type="text" placeholder="First Name" required />
                  </div>
                  <div className="col-xl-6 col-lg-6 col-sm-12">
                    <input type="text" placeholder="Last Name" required />
                  </div>
                  <div className="col-xl-6 col-lg-6 col-sm-12">
                    <input type="text" placeholder="Phone" required />
                  </div>
                  <div className="col-xl-6 col-lg-6 col-sm-12">
                    <input type="email" placeholder="Email" required />
                  </div>
                  <div className="col-xl-12 col-lg-12 col-sm-12">
                    <textarea name="msg" rows="4" placeholder="Your message" />
                  </div>
                  <div className="col-xl-6 col-lg-6 col-sm-12">
                    <button type="submit" className="bttn-mid btn-ylo">
                      <BiPaperPlane />
                      Send message
                    </button>
                  </div>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </section>
  </>
);

export default ContactSection;
