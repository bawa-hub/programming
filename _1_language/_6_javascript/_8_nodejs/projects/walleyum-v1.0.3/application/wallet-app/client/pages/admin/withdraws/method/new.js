/* eslint-disable jsx-a11y/label-has-associated-control */
import Head from 'next/head';
import React, { useEffect, useState } from 'react';
import { Dropdown, Spinner } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { BiErrorCircle, BiPlusCircle, BiTrash } from 'react-icons/bi';
import Toggle from 'react-toggle';
import EditorHeader from '../../../../components/admin/EditorHeader';
import Loader from '../../../../components/Loader';
import SidebarAdmin from '../../../../components/SidebarAdmin';
import UserHeaderAdmin from '../../../../components/UserHeaderAdmin';
import useCurrency from '../../../../data/useCurrency';
import withAuthAdmin from '../../../../hoc/withAuthAdmin';
import { methodCreate } from '../../../../lib/methodUpdate';
import inputNumber from '../../../../utils/inputNumber';

const MethodNew = ({ userData, settings }) => {
  const [currency, setCurrency] = useState();
  const [fields, setFields] = useState([{ name: '', type: 'text', required: true }]);

  const [actionLoader, setActionLoader] = useState(false);

  const { data: curData, loading: curLoading } = useCurrency();

  const { t } = useTranslation();

  useEffect(() => {
    setCurrency(curData?.data?.[0].symbol);
  }, [curData]);

  const handleUpdate = (e) => {
    e.preventDefault();
    const params = {
      name: e.target?.name?.value,
      currency,
      minAmount: e.target?.minAmount?.value,
      maxAmount: e.target?.maxAmount?.value,
      fixedCharge: e.target?.fixedCharge?.value,
      percentageCharge: e.target?.perCharge?.value,
      active: e.target?.active?.checked,
      params: fields
    };
    methodCreate(params, setActionLoader);
  };

  const deleteField = (indexField) => {
    const arr = fields.filter((item, index) => index !== indexField);
    setFields(arr);
  };
  const addField = () => {
    const arr = [...fields];
    arr.push({ name: '', required: false, type: 'text' });
    setFields(arr);
  };
  const handleFieldChange = (index, type, value) => {
    const arr = fields;
    arr[index][type] = value;
    setFields(arr);
  };

  if (curLoading) {
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
          <EditorHeader name={t('Add New Method')} />
          <div className="row">
            <div className="col-12 col-xl-12">
              <div className="basic-card">
                <div className="settings-box">
                  <form onSubmit={handleUpdate}>
                    <div className="row">
                      <div className="col">
                        <h4 className="box-title">{t('Details')}</h4>
                      </div>
                    </div>
                    <div className="row">
                      <div className="col-12 col-xl-6">
                        <div className="single-profile">
                          <label htmlFor="methodName">{t('Name')}</label>
                          <input
                            id="methodName"
                            required
                            name="name"
                            type="text"
                            placeholder={t('Name')}
                          />
                        </div>
                        <div className="single-profile">
                          <label htmlFor="minAmount">
                            {t('Minimum Amount')}
                            {' '}
                            (
                            {currency}
                            )
                          </label>
                          <input
                            id="minAmount"
                            onKeyPress={inputNumber}
                            name="minAmount"
                            type="text"
                            placeholder={t('Min Amount')}
                            required
                          />
                        </div>
                        <div className="single-profile">
                          <label htmlFor="maxAmount">
                            {t('Maximum Amount')}
                            {' '}
                            (
                            {currency}
                            )
                          </label>
                          <input
                            id="maxAmount"
                            onKeyPress={inputNumber}
                            name="maxAmount"
                            type="text"
                            required
                            placeholder={t('Max Amount')}
                          />
                        </div>
                        <div className="single-profile">
                          <label>{t('Active')}</label>
                          <Toggle
                            name="active"
                          />
                        </div>
                      </div>
                      <div className="col-12 col-xl-6">
                        <label htmlFor="curSelector">
                          {t('Currency')}
                        </label>
                        <Dropdown id="curSelector" className="mt-10 mb-20">
                          <Dropdown.Toggle className="bttn-small btn-emt dropdown-method" variant="link">
                            {currency || t('Please Select')}
                          </Dropdown.Toggle>
                          <Dropdown.Menu>
                            {curData?.data?.map((wallet) => (
                              <Dropdown.Item
                                key={wallet?.symbol}
                                onClick={() => setCurrency(wallet?.symbol)}
                              >
                                {wallet?.symbol}
                              </Dropdown.Item>
                            ))}
                          </Dropdown.Menu>
                        </Dropdown>
                        <div className="single-profile">
                          <label htmlFor="fixedCharge">
                            {t('Fixed Charge')}
                            {' '}
                            (
                            {currency}
                            )
                          </label>
                          <input
                            id="fixedCharge"
                            onKeyPress={inputNumber}
                            name="fixedCharge"
                            required
                            type="text"
                            placeholder={t('Fixed Charge')}
                          />
                        </div>
                        <div className="single-profile">
                          <label htmlFor="perCharge">
                            {t('Percentage Charge')}
                            {' '}
                            (%)
                          </label>
                          <input
                            id="perCharge"
                            name="perCharge"
                            onKeyPress={inputNumber}
                            required
                            type="text"
                            placeholder={t('Percentage Charge')}
                          />
                        </div>
                      </div>
                    </div>
                    <div className="row">
                      <div className="col">
                        <h4 className="box-title">{t('Fields')}</h4>
                      </div>
                    </div>
                    <div className="row">
                      <div className="col">
                        {fields?.map((field, index) => (
                          <div className="method-field" key={String(index)}>
                            <div className="single-profile">
                              <label>{t('Field Name')}</label>
                              <input
                                type="text"
                                onChange={(e) => handleFieldChange(index, 'name', e.target.value)}
                                defaultValue={field?.name}
                              />
                            </div>
                            <div className="single-profile">
                              <label>{t('Field Type')}</label>
                              <select
                                defaultValue={field?.type}
                                onChange={(e) => handleFieldChange(index, 'type', e.target.value)}
                              >
                                <option value="text">{t('Text')}</option>
                                <option value="email">{t('Email')}</option>
                                <option value="Number">{t('Number')}</option>
                              </select>
                            </div>
                            <div className="single-profile toggler">
                              <label>{t('Required')}</label>
                              <Toggle
                                defaultChecked={field?.required}
                                onChange={(e) => handleFieldChange(index, 'required', e.target.checked)}
                              />
                            </div>
                            <div className="action">
                              <button
                                type="button"
                                onClick={() => deleteField(index)}
                                className="action-btn danger"
                              >
                                <BiTrash />
                              </button>
                            </div>
                          </div>
                        ))}
                        <button
                          type="button"
                          className="bttn-mid btn-primary btn-new"
                          onClick={() => addField()}
                        >
                          <BiPlusCircle />
                          {t('New Field')}
                        </button>
                      </div>
                    </div>
                    <button type="submit" className="bttn-mid btn-ylo" disabled={actionLoader}>
                      {actionLoader ? (
                        <>
                          <Spinner animation="border" role="status" size="sm" />
                          {' '}
                          {t('Update Method')}
                        </>
                      ) : (
                        <>
                          <BiErrorCircle />
                          {t('Update Method')}
                        </>
                      )}
                    </button>
                  </form>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default withAuthAdmin(MethodNew);
