import Link from 'next/link';
import React from 'react';

const Footer = ({ site, footerMenu }) => (
  <>
    <footer className="footer">
      <div className="container">
        <div className="row">
          <div className="col-xl-6 col-sm-12 order-xl-1 order-2">
            <div className="copyright">
              <p>
                &copy;
                {' '}
                {site?.param1}
                {' '}
                {new Date().getFullYear()}
                . All Rights Reserved.
              </p>
            </div>
          </div>
          <div className="col-xl-6 col-sm-12 order-xl-2 order-1">
            <div className="footer-menu">
              <ul>
                {footerMenu?.map((item) => (
                  <li key={item.id}>
                    <Link href={item.url}>
                      <a>{item.name}</a>
                    </Link>
                  </li>
                ))}
              </ul>
            </div>
          </div>
        </div>
      </div>
    </footer>
  </>
);

export default Footer;
