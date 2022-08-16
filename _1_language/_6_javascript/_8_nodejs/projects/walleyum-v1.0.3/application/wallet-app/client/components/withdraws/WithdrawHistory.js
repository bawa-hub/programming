import dayjs from 'dayjs';
import React, { useState } from 'react';
import { Table } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import Pagination from 'react-js-pagination';
import useWithdraws from '../../data/useWithdraws';
import statusColor from '../../utils/statusColor';

const WithdrawHistory = () => {
  const [page, setPage] = useState(1);
  const { data } = useWithdraws(page, 10);
  const { t } = useTranslation();

  return (
    <>
      <Table striped hover className="dark-color" responsive>
        <thead>
          <tr>
            <th scope="col">{t('Date')}</th>
            <th scope="col">{t('Status')}</th>
            <th scope="col">{t('Description')}</th>
            <th scope="col">{t('Deducted')}</th>
            <th scope="col">{t('Fee')}</th>
            <th scope="col">{t('You Get')}</th>
            <th scope="col">{t('Method')}</th>
          </tr>
        </thead>
        <tbody>
          {data?.data?.map((history) => (
            <tr key={history?.id}>
              <td><span>{dayjs(history?.createdAt).format('DD/MM/YYYY')}</span></td>
              <td>
                <strong
                  style={{ textTransform: 'capitalize' }}
                  className={`status ${statusColor(history?.status)}`}
                >
                  {t(history?.status)}
                </strong>
              </td>
              <td>
                <strong>
                  {t('Withdraw Request')}
                  {' '}
                  #
                  {history?.id}
                </strong>
              </td>
              <td>
                <strong className="cl-red">
                  -
                  {history?.amount}
                  {' '}
                  {history?.currency}
                </strong>
              </td>
              <td>
                <strong>
                  {history?.fee}
                  {' '}
                  {history?.currency}
                </strong>
              </td>
              <td>
                <strong className="cl-green">
                  +
                  {history?.total}
                  {' '}
                  {history?.currency}
                </strong>
              </td>
              <td>
                {history?.method}
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
      <Pagination
        activePage={page}
        itemsCountPerPage={10}
        totalItemsCount={data?.count || 10}
        pageRangeDisplayed={10}
        onChange={(pageNumber) => setPage(pageNumber)}
        itemClass="page-item"
        linkClass="page-link"
      />
    </>
  );
};

export default WithdrawHistory;
