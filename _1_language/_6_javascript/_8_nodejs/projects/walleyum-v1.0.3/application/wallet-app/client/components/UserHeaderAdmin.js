import Link from 'next/link';
import React from 'react';
import { Dropdown } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import {
  BiCog, BiLogOut, BiMenu, BiUser
} from 'react-icons/bi';
import useSidebar from '../data/useSidebar';
import { signOutAdminRequest } from '../lib/authRequest';
import LanguageSwitch from './LanguageSwitch';

const Header = () => {
  const sidebar = useSidebar();
  const { t } = useTranslation();
  return (
    <>
      <header className="user-header">
        <div className="container-fluid">
          <div className="row">
            <div className="col-xl-12">
              <nav className="navbar navbar-expand-lg navbar-light px-0 justify-content-between">
                <button type="button" className="nav-toggle" onClick={() => sidebar?.toggle()}>
                  <BiMenu />
                  <span>{t('Toggle')}</span>
                </button>
                <div className="d-flex align-items-center" />
                <div className="dashboard-log">
                  <div className="d-flex align-items-center">
                    <div className="lng-switch-cont mr-10">
                      <LanguageSwitch />
                    </div>
                    <Dropdown>
                      <Dropdown.Toggle className="bttn-small btn-blue" variant="link">
                        <BiUser />
                        {t('Admin')}
                      </Dropdown.Toggle>
                      <Dropdown.Menu>
                        <Link href="/admin/settings">
                          <a className="dropdown-item">
                            <BiCog />
                            {t('Settings')}
                          </a>
                        </Link>
                        <Dropdown.Item onClick={() => signOutAdminRequest()}>
                          <BiLogOut />
                          {t('Logout')}
                        </Dropdown.Item>
                      </Dropdown.Menu>
                    </Dropdown>
                  </div>
                </div>
              </nav>
            </div>
          </div>
        </div>
      </header>
    </>
  );
};

export default Header;
