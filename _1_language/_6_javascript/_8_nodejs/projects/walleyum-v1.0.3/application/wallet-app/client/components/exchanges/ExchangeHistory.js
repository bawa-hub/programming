import dayjs from 'dayjs';
import React, { useState } from 'react';
import { Table } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import Pagination from 'react-js-pagination';
import useExchanges from '../../data/useExchanges';
import statusColor from '../../utils/statusColor';

const ExchangeHistory = () => {
  const [page, setPage] = useState(1);
  const { data } = useExchanges(page, 10);
  const { t } = useTranslation();

  return (
    <>
      <Table striped hover className="dark-color" responsive>
        <thead>
          <tr>
            <th scope="col">{t('Date')}</th>
            <th scope="col">{t('Status')}</th>
            <th scope="col">{t('Description')}</th>
            <th scope="col">{t('From')}</th>
            <th scope="col">{t('To')}</th>
            <th scope="col">{t('Fee')}</th>
            <th scope="col">{t('You Get')}</th>
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
                  {t('Exchange')}
                  {' '}
                  #
                  {history?.id}
                </strong>
              </td>
              <td>
                <strong className="cl-red">
                  -
                  {history?.amountFrom.toFixed(4)}
                  {' '}
                  {history?.from}
                </strong>
              </td>
              <td>
                <strong>
                  {history?.amountTo.toFixed(4)}
                  {' '}
                  {history?.to}
                </strong>
              </td>
              <td>
                {history?.fee.toFixed(4)}
                {' '}
                {history?.to}
              </td>
              <td>
                <strong className="cl-green">
                  +
                  {history?.total.toFixed(4)}
                  {' '}
                  {history?.to}
                </strong>
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

export default ExchangeHistory;
