import React from 'react';
import { useTranslation } from 'react-i18next';
import {
  BiCog, BiDollarCircle, BiHomeSmile, BiPaperPlane, BiPlusCircle, BiStore, BiTransferAlt, BiWallet
} from 'react-icons/bi';
import useSidebar from '../data/useSidebar';
import SidebarMenu from './SidebarMenu';

const Sidebar = ({ userData, settings }) => {
  const sidebar = useSidebar();
  const { t } = useTranslation();
  return (
    <>
      <div
        className={`backdrop ${sidebar?.visible ? 's-visible' : 's-hidden'}`}
        onClick={() => sidebar?.toggle()}
        role="button"
        aria-hidden="true"
      />
      <div className={`sidebar ${sidebar?.visible ? 's-visible' : 's-hidden'}`}>
        <a href="/" className="brand-logo">
          <img src={`${settings?.apiUrl?.param1}/public/${settings?.logo?.param1}`} alt="" />
        </a>
        <button type="button" className="closeBtn" onClick={() => sidebar?.toggle()}>
          <span aria-hidden="true">×</span>
          <span className="sr-only">{t('Close')}</span>
        </button>
        <div className="menu">
          <ul>
            <SidebarMenu href="/dashboard">
              <BiHomeSmile />
              <span>{t('Dashboard')}</span>
            </SidebarMenu>
            <SidebarMenu href="/add-money">
              <BiPlusCircle />
              <span>{t('Add Money')}</span>
            </SidebarMenu>
            <SidebarMenu href="/send-money">
              <BiPaperPlane />
              <span>{t('Transfer')}</span>
            </SidebarMenu>
            <SidebarMenu href="/withdraw">
              <BiWallet />
              <span>{t('Withdraw')}</span>
            </SidebarMenu>
            <SidebarMenu href="/exchange">
              <BiTransferAlt />
              <span>{t('Exchange')}</span>
            </SidebarMenu>
            <SidebarMenu href="/make-payment">
              <BiDollarCircle />
              <span>{t('Payment')}</span>
            </SidebarMenu>
            {userData?.role === 2 && (
            <SidebarMenu href="/merchant">
              <BiStore />
              <span>{t('Merchant')}</span>
            </SidebarMenu>
            )}
            <SidebarMenu href="/settings">
              <BiCog />
              <span>{t('Settings')}</span>
            </SidebarMenu>
          </ul>
          <div className="copyright">
            &copy;
            {' '}
            {settings?.site?.param1}
          </div>
        </div>
      </div>
    </>
  );
};

export default Sidebar;
