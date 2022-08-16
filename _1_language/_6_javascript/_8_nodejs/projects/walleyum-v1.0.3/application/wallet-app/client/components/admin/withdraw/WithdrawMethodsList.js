import Link from 'next/link';
import React from 'react';
import { Table } from 'react-bootstrap';
import { confirmAlert } from 'react-confirm-alert';
import { useTranslation } from 'react-i18next';
import { FiEdit, FiTrash } from 'react-icons/fi';
import { useMethodsAdmin } from '../../../data/useMethods';
import { methodDelete } from '../../../lib/methodUpdate';
import Loader from '../../Loader';

const WithdrawMethodsList = () => {
  const { data, loading } = useMethodsAdmin();
  const { t } = useTranslation();

  const deleteInit = (id, name) => {
    confirmAlert({
      title: `${name}`,
      message: t('Are you sure to remove this method?'),
      buttons: [
        {
          label: t('Yes'),
          onClick: () => methodDelete(id)
        },
        {
          label: t('No'),
        }
      ]
    });
  };

  return (
    <>
      {loading ? (
        <Loader height="300px" />
      ) : (
        <>
          <Table striped hover className="dark-color" responsive>
            <thead>
              <tr>
                <th scope="col">{t('Name')}</th>
                <th scope="col">{t('Currency')}</th>
                <th scope="col">{t('Limit')}</th>
                <th scope="col">{t('Charge')}</th>
                <th scope="col">{t('Status')}</th>
                <th scope="col">{t('Action')}</th>
              </tr>
            </thead>
            <tbody>
              {data?.data?.map((gateway) => (
                <tr key={gateway?.id}>
                  <td>
                    <strong>
                      {gateway?.name}
                    </strong>
                  </td>
                  <td>
                    <strong>
                      {gateway?.currency}
                    </strong>
                  </td>
                  <td>
                    <strong>
                      {gateway?.minAmount}
                      {' '}
                      {gateway?.currency}
                      {' - '}
                      {gateway?.maxAmount}
                      {' '}
                      {gateway?.currency}
                    </strong>
                  </td>
                  <td>
                    <strong>
                      {gateway?.fixedCharge}
                      {' '}
                      {gateway?.currency}
                      {' + '}
                      {gateway?.percentageCharge}
                      %
                    </strong>
                  </td>
                  <td>
                    <strong>
                      {gateway?.active ? t('Active') : t('Inactive')}
                    </strong>
                  </td>
                  <td width="15%" align="center">
                    <div className="d-flex">
                      <Link href={`/admin/withdraws/method/${gateway?.id}`}>
                        <a className="action-btn">
                          <FiEdit />
                        </a>
                      </Link>
                      <button
                        type="button"
                        onClick={() => deleteInit(gateway?.id, `${gateway?.name} - ${gateway?.currency}`)}
                        className="action-btn danger ml-10"
                      >
                        <FiTrash />
                      </button>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </Table>
        </>
      )}
    </>
  );
};

export default WithdrawMethodsList;
