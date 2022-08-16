import Link from 'next/link';
import React from 'react';
import { Dropdown } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import {
  BiCog, BiLinkAlt, BiLogOut, BiMenu, BiUser
} from 'react-icons/bi';
import useSidebar from '../data/useSidebar';
import { signOutRequest } from '../lib/authRequest';
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
                        {t('My Profile')}
                      </Dropdown.Toggle>
                      <Dropdown.Menu>
                        <Link href="/settings">
                          <a className="dropdown-item">
                            <BiCog />
                            {t('Settings')}
                          </a>
                        </Link>
                        <Link href="/settings?tab=2">
                          <a className="dropdown-item">
                            <BiLinkAlt />
                            {t('Linked Accounts')}
                          </a>
                        </Link>
                        <Dropdown.Item onClick={() => signOutRequest()}>
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
