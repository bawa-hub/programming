/* eslint-disable jsx-a11y/label-has-associated-control */
import Head from 'next/head';
import { useRouter } from 'next/router';
import React, { useState } from 'react';
import { Spinner, Table } from 'react-bootstrap';
import { confirmAlert } from 'react-confirm-alert';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle, BiTrash } from 'react-icons/bi';
import EditorHeader from '../../../components/admin/EditorHeader';
import Loader from '../../../components/Loader';
import SidebarAdmin from '../../../components/SidebarAdmin';
import UserHeaderAdmin from '../../../components/UserHeaderAdmin';
import { useMerchantById } from '../../../data/useMerchants';
import withAuthAdmin from '../../../hoc/withAuthAdmin';
import merchantUpdate, { merchantDelete } from '../../../lib/merchantUpdate';

const MerchantEdit = ({ userData, settings }) => {
  const router = useRouter();
  const { id } = router.query;

  const [actionLoader, setActionLoader] = useState(false);

  const { data, loading } = useMerchantById(id);

  const { t } = useTranslation();

  const handleUpdate = (e) => {
    e.preventDefault();
    const params = {
      name: e.target?.name?.value,
      email: e.target?.email?.value,
      address: e.target?.address?.value,
      status: e.target?.status?.value,
    };
    merchantUpdate(id, params, setActionLoader);
  };

  const deleteInit = () => {
    confirmAlert({
      title: `${data?.merId} - ${data?.name}`,
      message: 'Are you sure to remove this merchant? Remember the merchant will be demoted to user.',
      buttons: [
        {
          label: t('Yes'),
          onClick: () => merchantDelete(id, setActionLoader)
        },
        {
          label: t('No'),
        }
      ]
    });
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
          <EditorHeader name={`${t('Merchant')} #${data?.id}`} />
          <div className="row">
            <div className="col-12 col-xl-6">
              <div className="basic-card">
                <h4 className="box-title">{t('Merchant Details')}</h4>
                <div className="settings-box">
                  <form onSubmit={handleUpdate}>
                    <div className="single-profile">
                      <label htmlFor="merId">{t('Merchant ID')}</label>
                      <input
                        id="merId"
                        disabled
                        name="merId"
                        type="text"
                        placeholder={t('Merchant ID')}
                        defaultValue={data?.merId}
                      />
                    </div>
                    <div className="single-profile">
                      <label htmlFor="merName">{t('Merchant Name')}</label>
                      <input
                        id="merName"
                        name="name"
                        type="text"
                        placeholder={t('Merchant Name')}
                        defaultValue={data?.name}
                      />
                    </div>
                    <div className="single-profile">
                      <label htmlFor="merEmail">{t('Merchant Email')}</label>
                      <input
                        id="merEmail"
                        name="email"
                        type="email"
                        placeholder={t('Merchant Email')}
                        defaultValue={data?.email}
                      />
                    </div>
                    <div className="single-profile">
                      <label htmlFor="merAdd">{t('Merchant Address')}</label>
                      <input
                        id="merAdd"
                        name="address"
                        type="text"
                        placeholder={t('Merchant Address')}
                        defaultValue={data?.address}
                      />
                    </div>
                    <div className="single-profile">
                      <label htmlFor="merStatus">{t('Status')}</label>
                      <select name="status" id="merStatus" defaultValue={data?.status}>
                        <option value="pending">{t('Pending')}</option>
                        <option value="verified">{t('Verified')}</option>
                      </select>
                    </div>
                    <button type="submit" className="bttn-mid btn-ylo" disabled={actionLoader}>
                      {actionLoader ? (
                        <>
                          <Spinner animation="border" role="status" size="sm" />
                          {' '}
                          {t('Update Merchant')}
                        </>
                      ) : (
                        <>
                          <BiErrorCircle />
                          {t('Update Merchant')}
                        </>
                      )}
                    </button>
                    <button
                      type="button"
                      className="bttn-mid btn-danger ml-10"
                      onClick={() => deleteInit()}
                      disabled={actionLoader}
                    >
                      {actionLoader ? (
                        <>
                          <Spinner animation="border" role="status" size="sm" />
                          {' '}
                          {t('Remove Merchant')}
                        </>
                      ) : (
                        <>
                          <BiTrash />
                          {t('Remove Merchant')}
                        </>
                      )}
                    </button>
                  </form>
                </div>
              </div>
            </div>
            <div className="col-12 col-xl-6">
              <div className="basic-card mb-20">
                <h4 className="box-title">{t('Proof')}</h4>
                <div className="settings-box">
                  <Table striped hover className="dark-color details-table" responsive>
                    <tbody>
                      <tr>
                        <td className="head-td">{t('File')}</td>
                        <td>
                          {data?.proof ? (
                            <a href={data?.proof}>{t('Download/View')}</a>
                          ) : 'N/A'}
                        </td>
                      </tr>
                    </tbody>
                  </Table>
                </div>
              </div>
              <div className="basic-card">
                <h4 className="box-title">{t('User')}</h4>
                <div className="settings-box">
                  <Table striped hover className="dark-color details-table" responsive>
                    <tbody>
                      <tr>
                        <td className="head-td">{t('ID')}</td>
                        <td>
                          #
                          {data?.user?.id}
                        </td>
                      </tr>
                      <tr>
                        <td className="head-td">{t('Name')}</td>
                        <td>
                          {data?.user?.name}
                        </td>
                      </tr>
                      <tr>
                        <td className="head-td">{t('Email')}</td>
                        <td>
                          {data?.user?.email}
                        </td>
                      </tr>
                      <tr>
                        <td className="head-td">{t('Address')}</td>
                        <td>
                          {data?.user?.address || 'N/A'}
                        </td>
                      </tr>
                      <tr>
                        <td className="head-td">{t('Phone')}</td>
                        <td>
                          {data?.user?.phone || 'N/A'}
                        </td>
                      </tr>
                      <tr>
                        <td className="head-td">{t('Activation')}</td>
                        <td>
                          {data?.user?.active ? t('Active') : t('Inactive')}
                        </td>
                      </tr>
                      <tr>
                        <td className="head-td">{t('KYC')}</td>
                        <td>
                          {data?.user?.kyc ? t('Verified') : t('Pending Verification')}
                        </td>
                      </tr>
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

export default withAuthAdmin(MerchantEdit);
