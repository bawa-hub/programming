/* eslint-disable jsx-a11y/label-has-associated-control */
import Head from 'next/head';
import { useRouter } from 'next/router';
import React, { useState } from 'react';
import { Spinner, Table } from 'react-bootstrap';
import { confirmAlert } from 'react-confirm-alert';
import { useTranslation } from 'react-i18next';
import {
  BiAnalyse, BiCheck, BiErrorCircle, BiMailSend, BiScan, BiSync, BiTrash, BiWallet, BiX
} from 'react-icons/bi';
import Toggle from 'react-toggle';
import EditorHeader from '../../../components/admin/EditorHeader';
import BalanceManage from '../../../components/admin/user/BalanceManage';
import Loader from '../../../components/Loader';
import SidebarAdmin from '../../../components/SidebarAdmin';
import TabModule from '../../../components/tabs/TabModule';
import UserTab from '../../../components/tabs/UserTab';
import UserHeaderAdmin from '../../../components/UserHeaderAdmin';
import { useKycById } from '../../../data/useKyc';
import { useUserById } from '../../../data/useUsers';
import { useWalletByUserId } from '../../../data/useWallet';
import withAuthAdmin from '../../../hoc/withAuthAdmin';
import { kycAdminAction } from '../../../lib/kycUpdate';
import { merchantCreate, merchantDeleteFromUser } from '../../../lib/merchantUpdate';
import { profileUpdateAdmin, userDelete } from '../../../lib/profileUpdate';
import { sendMail } from '../../../lib/settingsUpdate';

const UserEdit = ({ userData, settings }) => {
  const router = useRouter();
  const { id, tab } = router.query;

  const [actionLoader, setActionLoader] = useState(false);

  const { data, loading } = useUserById(id);
  const { data: kyc, loading: kycLoading } = useKycById(id);
  const { data: userWallet, loading: walletLoading } = useWalletByUserId(id);

  const { t } = useTranslation();

  const handleUpdate = (e) => {
    e.preventDefault();
    const params = {
      name: e.target?.name?.value,
      email: e.target?.email?.value,
      phone: e.target?.phone?.value,
      password: e.target?.password?.value || undefined,
      address: e.target?.address?.value,
      active: e.target?.active?.checked,
      kyc: e.target?.kyc?.checked,
      role: e.target?.role?.value,
    };
    profileUpdateAdmin(id, params, setActionLoader);
  };

  const handleUpdateMerchant = (e) => {
    e.preventDefault();
    const params = {
      name: e.target?.name?.value,
      email: e.target?.email?.value,
      address: e.target?.address?.value,
      status: e.target?.status?.value,
      userId: id
    };
    merchantCreate(params, setActionLoader);
  };

  const handleMail = (e) => {
    e.preventDefault();
    const params = {
      userId: id,
      subject: e.target?.subject?.value,
      message: e.target?.message?.value,
    };
    sendMail(params, setActionLoader);
  };

  const handleKyc = (type, updateId, userId) => {
    kycAdminAction(type, updateId, userId, setActionLoader);
  };

  const deleteInit = () => {
    confirmAlert({
      title: `#${data?.id} - ${data?.name}`,
      message: t('Are you sure to remove this user?'),
      buttons: [
        {
          label: t('Yes'),
          onClick: () => userDelete(id, setActionLoader)
        },
        {
          label: t('No'),
        }
      ]
    });
  };

  const deleteInitMerc = () => {
    confirmAlert({
      title: `#${data?.id} - ${data?.name}`,
      message: t('Are you sure to remove this user as merchant?'),
      buttons: [
        {
          label: t('Yes'),
          onClick: () => merchantDeleteFromUser(id, data?.merchant?.id, setActionLoader)
        },
        {
          label: t('No'),
        }
      ]
    });
  };

  if (loading || kycLoading || walletLoading) {
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
          <EditorHeader name={`${t('User')} #${data?.id} (${data?.name})`} />
          <div className="row">
            <UserTab
              description={t('Manage your users from this panel')}
              defaultTab={parseInt(tab, 10)}
            >
              <TabModule icon={<BiAnalyse />} name={t('User Details')}>
                <div className="basic-card">
                  <h4 className="box-title">{t('User Details')}</h4>
                  <div className="settings-box">
                    <form onSubmit={handleUpdate}>
                      <div className="single-profile">
                        <label htmlFor="userName">{t('Full Name')}</label>
                        <input
                          id="userName"
                          name="name"
                          type="text"
                          placeholder={t('Full Name')}
                          defaultValue={data?.name}
                        />
                      </div>
                      <div className="single-profile">
                        <label htmlFor="userEmail">{t('Email')}</label>
                        <input
                          id="userEmail"
                          name="email"
                          type="email"
                          placeholder={t('Email')}
                          defaultValue={data?.email}
                        />
                      </div>
                      <div className="single-profile">
                        <label htmlFor="userPhone">{t('Phone')}</label>
                        <input
                          id="userPhone"
                          name="phone"
                          type="text"
                          placeholder={t('Phone Number')}
                          defaultValue={data?.phone}
                        />
                      </div>
                      <div className="single-profile">
                        <label htmlFor="userAdd">{t('Address')}</label>
                        <input
                          id="userAdd"
                          name="address"
                          type="text"
                          placeholder={t('Address')}
                          defaultValue={data?.address}
                        />
                      </div>
                      <div className="single-profile">
                        <label htmlFor="userPassword">{t('Password')}</label>
                        <input
                          id="userPassword"
                          name="password"
                          type="password"
                          placeholder={t('Leave blank to keep unchanged')}
                        />
                      </div>
                      <div className="single-profile">
                        <label htmlFor="userRole">{t('Role')}</label>
                        {data?.role === 2 ? (
                          <p>{t('This user is a merchant')}</p>
                        ) : (
                          <select name="role" id="userRole" defaultValue={data?.role}>
                            <option value="0">{t('Admin')}</option>
                            <option value="1">{t('User')}</option>
                          </select>
                        )}
                      </div>
                      <div className="single-profile">
                        <label>{t('Active')}</label>
                        <Toggle
                          defaultChecked={data?.active}
                          name="active"
                        />
                      </div>
                      <div className="single-profile">
                        <label>{t('KYC')}</label>
                        <Toggle
                          defaultChecked={data?.kyc}
                          name="kyc"
                        />
                      </div>
                      <button type="submit" className="bttn-mid btn-ylo" disabled={actionLoader}>
                        {actionLoader ? (
                          <>
                            <Spinner animation="border" role="status" size="sm" />
                            {' '}
                            {t('Update User')}
                          </>
                        ) : (
                          <>
                            <BiErrorCircle />
                            {t('Update User')}
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
                            {t('Remove User')}
                          </>
                        ) : (
                          <>
                            <BiTrash />
                            {t('Remove User')}
                          </>
                        )}
                      </button>
                    </form>
                  </div>
                </div>
              </TabModule>
              <TabModule icon={<BiWallet />} name={t('Balance')}>
                <div className="basic-card">
                  <h4 className="box-title">{t('Balance')}</h4>
                  <div className="settings-box">
                    <Table striped hover className="dark-color details-table" responsive>
                      <thead>
                        <tr>
                          <th scope="col">{t('Currency')}</th>
                          <th scope="col">{t('Amount')}</th>
                          <th scope="col">{t('Action')}</th>
                        </tr>
                      </thead>
                      <tbody>
                        {userWallet?.map((cur) => (
                          <BalanceManage
                            key={cur?.currency}
                            currency={cur}
                            userId={id}
                          />
                        ))}
                      </tbody>
                    </Table>
                  </div>
                </div>
              </TabModule>
              <TabModule icon={<BiScan />} name={t('KYC')}>
                <div className="basic-card">
                  <h4 className="box-title">{t('KYC')}</h4>
                  <div className="settings-box">
                    <Table striped hover className="dark-color details-table" responsive>
                      <tbody>
                        <tr>
                          <td className="head-td">Type</td>
                          <td>
                            {kyc?.type || 'N/A'}
                          </td>
                        </tr>
                        <tr>
                          <td className="head-td">Front</td>
                          <td>
                            {kyc?.front ? <img src={`${settings?.apiUrl?.param1}/public/${kyc?.front}`} alt="Front Side" /> : 'N/A'}
                          </td>
                        </tr>
                        <tr>
                          <td className="head-td">Back</td>
                          <td>
                            {kyc?.back ? <img src={`${settings?.apiUrl?.param1}/public/${kyc?.back}`} alt="Back Side" /> : 'N/A'}
                          </td>
                        </tr>
                        <tr>
                          <td className="head-td">Selfie</td>
                          <td>
                            {kyc?.selfie ? <img src={`${settings?.apiUrl?.param1}/public/${kyc?.selfie}`} alt="Selfie" /> : 'N/A'}
                          </td>
                        </tr>
                        <tr>
                          <td className="head-td">Status</td>
                          <td>
                            {kyc?.status || 'Not Yet Submitted'}
                          </td>
                        </tr>
                      </tbody>
                    </Table>
                    {kyc?.status === 'submitted' && (
                    <>
                      {actionLoader ? (
                        <Spinner animation="border" role="status" size="sm" />
                      ) : (
                        <>
                          <button
                            onClick={() => handleKyc('accept', kyc?.id, data?.id)}
                            type="submit"
                            className="bttn-mid btn-ylo mr-10 mt-20"
                            disabled={actionLoader}
                          >
                            <BiCheck />
                            Approve
                          </button>
                          <button
                            onClick={() => handleKyc('decline', kyc?.id, data?.id)}
                            type="submit"
                            className="bttn-mid btn-danger mt-20"
                            disabled={actionLoader}
                          >
                            <BiX />
                            Decline
                          </button>
                        </>
                      )}
                    </>
                    )}
                  </div>
                </div>
              </TabModule>
              <TabModule icon={<BiMailSend />} name={t('Send Email')}>
                <div className="basic-card">
                  <h4 className="box-title">{t('Send Email')}</h4>
                  <div className="settings-box">
                    <form onSubmit={handleMail}>
                      <div className="single-profile">
                        <label htmlFor="mailSub">{t('Subject')}</label>
                        <input
                          id="mailSub"
                          name="subject"
                          type="text"
                          placeholder={t('Subject')}
                          required
                        />
                      </div>
                      <div className="single-profile">
                        <label htmlFor="mailMsg">{t('Message')}</label>
                        <textarea
                          id="mailMsg"
                          name="message"
                          placeholder={t('Message')}
                          row={7}
                          style={{ height: 'auto' }}
                          required
                        />
                      </div>
                      <button type="submit" className="bttn-mid btn-ylo" disabled={actionLoader}>
                        {actionLoader ? (
                          <>
                            <Spinner animation="border" role="status" size="sm" />
                            {' '}
                            {t('Send Mail')}
                          </>
                        ) : (
                          <>
                            <BiErrorCircle />
                            {t('Send Mail')}
                          </>
                        )}
                      </button>
                    </form>
                  </div>
                </div>
              </TabModule>
              <TabModule icon={<BiSync />} name={t('Convert To Merchant')}>
                <div className="basic-card">
                  <h4 className="box-title">{t('Convert To Merchant')}</h4>
                  <div className="settings-box">
                    <form onSubmit={handleUpdateMerchant}>
                      {data?.role === 2 ? (
                        <button
                          type="button"
                          className="bttn-mid btn-danger ml-10"
                          onClick={() => deleteInitMerc()}
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
                      ) : (
                        <>
                          <div className="single-profile">
                            <label htmlFor="merName">{t('Merchant Name')}</label>
                            <input
                              id="merName"
                              name="name"
                              type="text"
                              placeholder={t('Merchant Name')}
                            />
                          </div>
                          <div className="single-profile">
                            <label htmlFor="merEmail">{t('Merchant Email')}</label>
                            <input
                              id="merEmail"
                              name="email"
                              type="email"
                              placeholder={t('Merchant Email')}
                            />
                          </div>
                          <div className="single-profile">
                            <label htmlFor="merAdd">{t('Merchant Address')}</label>
                            <input
                              id="merAdd"
                              name="address"
                              type="text"
                              placeholder={t('Address')}
                            />
                          </div>
                          <div className="single-profile">
                            <label htmlFor="merStat">{t('Status')}</label>
                            <select name="status" id="merStat">
                              <option value="pending">{t('Pending')}</option>
                              <option value="verified">{t('Verified')}</option>
                            </select>
                          </div>
                          <button type="submit" className="bttn-mid btn-ylo" disabled={actionLoader}>
                            {actionLoader ? (
                              <>
                                <Spinner animation="border" role="status" size="sm" />
                                {' '}
                                {t('Convert To Merchant')}
                              </>
                            ) : (
                              <>
                                <BiErrorCircle />
                                {t('Convert To Merchant')}
                              </>
                            )}
                          </button>
                        </>
                      )}
                    </form>
                  </div>
                </div>
              </TabModule>
            </UserTab>
          </div>
        </div>
      </div>
    </>
  );
};

export default withAuthAdmin(UserEdit);
