import Head from 'next/head';
import { useRouter } from 'next/router';
import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import AddPage from '../../../components/admin/builder/AddPage';
import PageBuilder from '../../../components/admin/builder/PageBuilder';
import SidebarAdmin from '../../../components/SidebarAdmin';
import UserHeaderAdmin from '../../../components/UserHeaderAdmin';
import withAuthAdmin from '../../../hoc/withAuthAdmin';

const EditPage = ({ userData, settings }) => {
  const router = useRouter();
  const { slug } = router.query;
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
          <PageBuilder slug={slug} />
        </div>
      </div>
    </>
  );
};

export default withAuthAdmin(EditPage);
