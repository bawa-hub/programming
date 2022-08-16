import Head from 'next/head';
import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { BiListCheck, BiPlusCircle } from 'react-icons/bi';
import AddPage from '../../../components/admin/builder/AddPage';
import FooterMenuBuilder from '../../../components/admin/builder/FooterMenuBuilder';
import LogoFavicon from '../../../components/admin/builder/LogoFavicon';
import MenuBuilder from '../../../components/admin/builder/MenuBuilder';
import PageList from '../../../components/admin/builder/PageList';
import FaqBuilder from '../../../components/admin/repeater/faq/FaqBuilder';
import ServicesBuilder from '../../../components/admin/repeater/services/ServicesBuilder';
import SolutionsBuilder from '../../../components/admin/repeater/solutions/SolutionsBuilder';
import SidebarAdmin from '../../../components/SidebarAdmin';
import TabModule from '../../../components/tabs/TabModule';
import UserTab from '../../../components/tabs/UserTab';
import UserHeaderAdmin from '../../../components/UserHeaderAdmin';
import withAuthAdmin from '../../../hoc/withAuthAdmin';

const Builder = ({ userData, settings }) => {
  const [active, setActive] = useState(false);
  const { t } = useTranslation();
  return (
    <>
      <Head>
        <title>
          {t('Admin Panel')}
        </title>
      </Head>
      <UserHeaderAdmin />
      <SidebarAdmin userData={userData} settings={settings} />
      <AddPage active={active} setActive={setActive} />
      <div className="content-body">
        <div className="container-fluid">
          <div className="row">
            <div className="col">
              <UserTab
                title={t('Site Builder')}
                description={t('Manage UI of your site from this panel')}
              >
                <TabModule icon={<BiListCheck />} name={t('Pages')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Pages')}</h4>
                    <PageList />
                    <button
                      type="button"
                      className="bttn-mid btn-primary btn-new mt-10"
                      onClick={() => setActive(true)}
                    >
                      <BiPlusCircle />
                      {t('Add New Page')}
                    </button>
                  </div>
                </TabModule>
                <TabModule icon={<BiListCheck />} name={t('Logo & Favicon')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Logo & Favicon')}</h4>
                    <LogoFavicon />
                  </div>
                </TabModule>
                <TabModule icon={<BiListCheck />} name={t('Main Menu')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Main Menu')}</h4>
                    <MenuBuilder />
                  </div>
                </TabModule>
                <TabModule icon={<BiListCheck />} name={t('Footer Menu')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Footer Menu')}</h4>
                    <FooterMenuBuilder />
                  </div>
                </TabModule>
                <TabModule icon={<BiListCheck />} name={t('Services')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Services')}</h4>
                    <ServicesBuilder />
                  </div>
                </TabModule>
                <TabModule icon={<BiListCheck />} name={t('Solutions')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('Solutions')}</h4>
                    <SolutionsBuilder />
                  </div>
                </TabModule>
                <TabModule icon={<BiListCheck />} name={t('FAQ')}>
                  <div className="basic-card">
                    <h4 className="box-title">{t('FAQ')}</h4>
                    <FaqBuilder />
                  </div>
                </TabModule>
              </UserTab>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default withAuthAdmin(Builder);
