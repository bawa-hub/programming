import React, { useState } from 'react';
import { Table } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import Pagination from 'react-js-pagination';
import { useRequestsAdmin } from '../../../data/useRequests';
import statusColor from '../../../utils/statusColor';
import Loader from '../../Loader';

const RequestsList = ({ pending }) => {
  const [page, setPage] = useState(1);
  const [depId, setDepId] = useState('');
  const [status, setStatus] = useState(pending ? 'pending' : 'all');
  const [perPage, setPerPage] = useState(25);
  const { t } = useTranslation();

  const { data, loading } = useRequestsAdmin(page, perPage, status, depId);

  return (
    <>
      <div className="filter">
        <div className="filter-box">
          <p>
            {t('TRX ID')}
            {' '}
          </p>
          <input type="text" value={depId} onChange={(e) => setDepId(e.target.value)} />
        </div>
        {!pending && (
        <div className="filter-box">
          <p>
            {t('Status')}
            {' '}
          </p>
          <select value={status} name="status" onChange={(e) => setStatus(e.target.value)}>
            <option value="all">{t('All')}</option>
            <option value="pending">{t('Pending')}</option>
            <option value="success">{t('Success')}</option>
            <option value="failed">{t('Failed')}</option>
          </select>
        </div>
        )}
        <div className="filter-box">
          <p>
            {t('Per Page')}
            {' '}
          </p>
          <select value={perPage} name="status" onChange={(e) => setPerPage(e.target.value)}>
            <option value="20">25</option>
            <option value="50">50</option>
            <option value="100">100</option>
            <option value="150">150</option>
            <option value="200">200</option>
          </select>
        </div>
      </div>
      {loading ? (
        <Loader height="300px" />
      ) : (
        <>
          <Table striped hover className="dark-color" responsive>
            <thead>
              <tr>
                <th scope="col">{t('ID')}</th>
                <th scope="col">{t('Status')}</th>
                <th scope="col">{t('TRX ID')}</th>
                <th scope="col">{t('Amount')}</th>
                <th scope="col">{t('Merchant')}</th>
                <th scope="col">{t('Customer')}</th>
              </tr>
            </thead>
            <tbody>
              {data?.data?.map((history) => (
                <tr key={history?.id}>
                  <td>
                    <strong>
                      #
                      {history?.id}
                    </strong>
                  </td>
                  <td>
                    <strong
                      className={`status ${statusColor(history?.status)}`}
                      style={{ textTransform: 'capitalize' }}
                    >
                      {t(history?.status)}
                    </strong>
                  </td>
                  <td>
                    <strong>
                      {history?.trxId}
                    </strong>
                  </td>
                  <td>
                    <strong>
                      {history?.amount}
                      {' '}
                      {history?.currency}
                    </strong>
                  </td>
                  <td>
                    <strong>
                      {history?.merchant?.name}
                      {' '}
                      (
                      {history?.merchant?.merId}
                      )
                    </strong>
                  </td>
                  <td>
                    <strong>
                      {history?.customer}
                    </strong>
                  </td>
                </tr>
              ))}
            </tbody>
          </Table>
          <Pagination
            activePage={page}
            itemsCountPerPage={perPage}
            totalItemsCount={data?.count || perPage}
            pageRangeDisplayed={perPage}
            onChange={(pageNumber) => setPage(pageNumber)}
            itemClass="page-item"
            linkClass="page-link"
          />
        </>
      )}
    </>
  );
};

export default RequestsList;
