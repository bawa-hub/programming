import Link from 'next/link';
import React, { useState } from 'react';
import { Image, Nav, Navbar } from 'react-bootstrap';
import Spinner from 'react-bootstrap/Spinner';
import { BiMenu, BiUserPlus } from 'react-icons/bi';
import { RiDashboardLine } from 'react-icons/ri';
import useCheckAuth from '../../data/useCheckAuth';

const SiteHeader = ({ logo, apiUrl, mainMenu }) => {
  const { data, loading } = useCheckAuth();
  const [show, setShow] = useState(false);
  return (
    <>
      <header className="site-header">
        <div className="container-fluid">
          <div className="row">
            <div className="col-xl-12">
              <Navbar expand="lg">
                <Link href="/">
                  <a className="site-logo navbar-brand">
                    <Image src={`${apiUrl?.param1}/public/${logo?.param1}`} alt="Logo" />
                  </a>
                </Link>
                <button
                  type="button"
                  onClick={() => setShow(!show)}
                  aria-label="Toggle navigation"
                  className="navbar-toggler collapsed"
                >
                  <span className="navbar-toggler-svg">
                    <BiMenu />
                  </span>
                </button>
                <div className={`navbar-collapse collapse ${show ? 'show' : 'hide'}`}>
                  <Nav className="ml-auto mainmenu">
                    {mainMenu?.map((item) => (
                      <Link href={item.url} key={item.id}>
                        <a className="nav-link">{item.name}</a>
                      </Link>
                    ))}
                    {loading ? (
                      <Link href="/">
                        <a className="bttn-small btn-ylo header-btn" style={{ minWidth: '140px', textAlign: 'center' }}>
                          <Spinner animation="border" role="status" size="sm">
                            <span className="visually-hidden">Loading...</span>
                          </Spinner>
                        </a>
                      </Link>
                    ) : (
                      <Link href={data ? '/dashboard' : '/login'}>
                        <a className="bttn-small btn-ylo header-btn">
                          {data ? <RiDashboardLine /> : <BiUserPlus />}
                          {data ? 'Dashboard' : 'Login/Register'}
                        </a>
                      </Link>
                    )}
                  </Nav>
                </div>
              </Navbar>
            </div>
          </div>
        </div>
      </header>
    </>
  );
};

export default SiteHeader;
