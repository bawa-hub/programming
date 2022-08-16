import dayjs from 'dayjs';
import React, { useState } from 'react';
import { Table } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import Pagination from 'react-js-pagination';
import { useReceive } from '../../data/useTransfers';

const ReceiveHistory = () => {
  const [page, setPage] = useState(1);
  const { data } = useReceive(page, 10);
  const { t } = useTranslation();

  return (
    <>
      <Table striped hover className="dark-color" responsive>
        <thead>
          <tr>
            <th scope="col">{t('Date')}</th>
            <th scope="col">{t('From')}</th>
            <th scope="col">{t('TRX ID')}</th>
            <th scope="col">{t('Amount')}</th>
          </tr>
        </thead>
        <tbody>
          {data?.data?.map((history) => (
            <tr key={history?.id}>
              <td><span>{dayjs(history?.createdAt).format('DD/MM/YYYY')}</span></td>
              <td><strong className="status success">{history?.email}</strong></td>
              <td>
                <strong>
                  {history?.trxId}
                </strong>
              </td>
              <td>
                <strong className="cl-green">
                  +
                  {history?.amount}
                  {' '}
                  {history?.currency}
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

export default ReceiveHistory;
