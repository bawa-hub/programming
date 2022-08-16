import React from 'react';
import { useTranslation } from 'react-i18next';
import {
  BiCog, BiDollarCircle, BiHomeSmile, BiPaintRoll,
  BiPaperPlane, BiPlusCircle, BiStore, BiTransferAlt, BiUser, BiWallet
} from 'react-icons/bi';
import useSidebar from '../data/useSidebar';
import SidebarMenu from './SidebarMenu';

const SidebarAdmin = ({ settings }) => {
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
            <SidebarMenu href="/admin/dashboard">
              <BiHomeSmile />
              <span>{t('Dashboard')}</span>
            </SidebarMenu>
            <SidebarMenu href="/admin/deposits">
              <BiPlusCircle />
              <span>{t('Deposits')}</span>
            </SidebarMenu>
            <SidebarMenu href="/admin/transfers">
              <BiPaperPlane />
              <span>{t('Transfers')}</span>
            </SidebarMenu>
            <SidebarMenu href="/admin/withdraws">
              <BiWallet />
              <span>{t('Withdraws')}</span>
            </SidebarMenu>
            <SidebarMenu href="/admin/exchanges">
              <BiTransferAlt />
              <span>{t('Exchanges')}</span>
            </SidebarMenu>
            <SidebarMenu href="/admin/pays">
              <BiDollarCircle />
              <span>{t('Payments')}</span>
            </SidebarMenu>
            <SidebarMenu href="/admin/merchants">
              <BiStore />
              <span>{t('Merchants')}</span>
            </SidebarMenu>
            <SidebarMenu href="/admin/users">
              <BiUser />
              <span>{t('Users')}</span>
            </SidebarMenu>
            <SidebarMenu href="/admin/settings">
              <BiCog />
              <span>{t('Settings')}</span>
            </SidebarMenu>
            <SidebarMenu href="/admin/builder">
              <BiPaintRoll />
              <span>{t('Site Builder')}</span>
            </SidebarMenu>
          </ul>
          <div className="copyright">
            {settings?.site?.param1}
            {' '}
            {t('Admin Panel')}
          </div>
        </div>
      </div>
    </>
  );
};

export default SidebarAdmin;
