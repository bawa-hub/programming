import { useRouter } from 'next/router';
import React from 'react';
import { Table } from 'react-bootstrap';
import { confirmAlert } from 'react-confirm-alert';
import { useTranslation } from 'react-i18next';
import { FiEdit, FiTrash } from 'react-icons/fi';
import useCurrencies from '../../../data/useCurrency';
import currencyDelete from '../../../lib/currencyRequest';
import Loader from '../../Loader';

const WalletList = ({ settings }) => {
  const { data, loading } = useCurrencies();
  const router = useRouter();

  const { t } = useTranslation();

  const deleteInit = (curId, curSymbol, curName) => {
    confirmAlert({
      title: `${curName} - ${curSymbol}`,
      message: t('Are you sure to remove this currency?'),
      buttons: [
        {
          label: t('Yes'),
          onClick: () => currencyDelete(curId)
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
                <th scope="col">{t('Symbol')}</th>
                <th scope="col">{t('Crypto')}</th>
                <th scope="col">
                  {t('Rate')}
                  {' '}
                  (USD)
                </th>
                <th scope="col">{t('Rate From API')}</th>
                <th scope="col">{t('Action')}</th>
              </tr>
            </thead>
            <tbody>
              {data?.data?.map((cur) => (
                <tr key={cur?.id}>
                  <td>
                    {cur?.icon && (
                    <img width="30px" src={`${settings.apiUrl.param1}/public/${cur?.icon}`} alt="Icon" />
                    )}
                    <strong>
                      {cur?.name}
                    </strong>
                  </td>
                  <td>
                    <strong>
                      {cur?.symbol}
                    </strong>
                  </td>
                  <td>
                    <strong>
                      {cur?.crypto ? t('Yes') : t('No')}
                    </strong>
                  </td>
                  <td>
                    <strong>
                      {cur?.rateUsd}
                      {' USD'}
                    </strong>
                  </td>
                  <td>
                    <strong>
                      {cur?.ratefromApi ? t('Yes') : t('No')}
                    </strong>
                  </td>
                  <td width="15%" align="center">
                    <div className="d-flex">
                      <button
                        type="button"
                        onClick={() => router.push(`/admin/settings/wallet/${cur?.id}`)}
                        className="action-btn"
                      >
                        <FiEdit />
                      </button>
                      <button
                        type="button"
                        onClick={() => deleteInit(cur?.id, cur?.symbol, cur?.name)}
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

export default WalletList;
