/* eslint-disable jsx-a11y/label-has-associated-control */
import Head from 'next/head';
import { useRouter } from 'next/router';
import React, { useState } from 'react';
import { Spinner, Table } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { BiCheck, BiX } from 'react-icons/bi';
import EditorHeader from '../../../components/admin/EditorHeader';
import Loader from '../../../components/Loader';
import SidebarAdmin from '../../../components/SidebarAdmin';
import UserHeaderAdmin from '../../../components/UserHeaderAdmin';
import { useWithdrawById } from '../../../data/useWithdraws';
import withAuthAdmin from '../../../hoc/withAuthAdmin';
import withdrawUpdate from '../../../lib/withdrawUpdate';

const WithdrawEdit = ({ userData, settings }) => {
  const router = useRouter();
  const { id } = router.query;

  const [actionLoader, setActionLoader] = useState(false);

  const { data, loading } = useWithdrawById(id);

  const { t } = useTranslation();

  const handleUpdate = (type) => {
    withdrawUpdate(id, type, setActionLoader);
  };

  if (loading) {
    return <Loader height="100vh" />;
  }
  return (
    <>
      <Head>
        <title>
          {t('Admin Panel')}
        </title>
      </Head>
      <UserHeaderAdmin />
      <SidebarAdmin userData={userData} settings={settings} />
      <div className="content-body">
        <div className="container-fluid">
          <EditorHeader name={`${t('Withdraw')} #${data?.id}`} />
          <div className="row">
            <div className="col-12 col-xl-6">
              <div className="basic-card">
                <h4 className="box-title">{t('Withdraw Details')}</h4>
                <div className="settings-box">
                  <Table striped hover className="dark-color details-table" responsive>
                    <tbody>
                      <tr>
                        <td className="head-td">{t('Status')}</td>
                        <td style={{ textTransform: 'capitalize' }}>
                          {data?.status}
                        </td>
                      </tr>
                      <tr>
                        <td className="head-td">{t('Amount')}</td>
                        <td>
                          {data?.amount}
                          {' '}
                          {data?.currency}
                        </td>
                      </tr>
                      <tr>
                        <td className="head-td">{t('Fee')}</td>
                        <td>
                          {data?.fee}
                          {' '}
                          {data?.currency}
                        </td>
                      </tr>
                      <tr>
                        <td className="head-td">{t('User Getting')}</td>
                        <td>
                          {data?.total}
                          {' '}
                          {data?.currency}
                        </td>
                      </tr>
                      <tr>
                        <td className="head-td">{t('User')}</td>
                        <td>
                          {data?.user?.email}
                          {' '}
                          (#
                          {data?.user?.id}
                          )
                        </td>
                      </tr>
                    </tbody>
                  </Table>
                  {data?.status === 'pending' && (
                  <>
                    {actionLoader ? (
                      <Spinner animation="border" role="status" size="sm" />
                    ) : (
                      <>
                        <button
                          onClick={() => handleUpdate('accept')}
                          type="submit"
                          className="bttn-mid btn-ylo mr-10 mt-20"
                          disabled={actionLoader}
                        >
                          <BiCheck />
                          {t('Approve')}
                        </button>
                        <button
                          onClick={() => handleUpdate('decline')}
                          type="submit"
                          className="bttn-mid btn-danger mt-20"
                          disabled={actionLoader}
                        >
                          <BiX />
                          {t('Decline')}
                        </button>
                      </>
                    )}
                  </>
                  )}
                </div>
              </div>
            </div>
            <div className="col-12 col-xl-6">
              <div className="basic-card">
                <h4 className="box-title">{t('Preffered Method')}</h4>
                <div className="settings-box">
                  <Table striped hover className="dark-color details-table" responsive>
                    <tbody>
                      <tr>
                        <td className="head-td">{t('Method')}</td>
                        <td style={{ textTransform: 'capitalize' }}>
                          {data?.method}
                        </td>
                      </tr>
                      {data?.params?.map((field) => (
                        <tr key={field?.name}>
                          <td className="head-td">{field?.name}</td>
                          <td>
                            {field?.value}
                          </td>
                        </tr>
                      ))}
                    </tbody>
                  </Table>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default withAuthAdmin(WithdrawEdit);
