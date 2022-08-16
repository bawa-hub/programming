import Link from 'next/link';
import React from 'react';
import { Table } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { FiEdit } from 'react-icons/fi';
import { useGatewayAdmin } from '../../../data/useGateways';
import GatewayLogo from '../../deposit/GatewayLogo';
import Loader from '../../Loader';

const MethodsList = () => {
  const { data, loading } = useGatewayAdmin();
  const { t } = useTranslation();

  return (
    <>
      {loading ? (
        <Loader height="300px" />
      ) : (
        <>
          <Table striped hover className="dark-color" responsive>
            <thead>
              <tr>
                <th scope="col">{t('Logo')}</th>
                <th scope="col">{t('Name')}</th>
                <th scope="col">{t('Status')}</th>
                <th scope="col">{t('Action')}</th>
              </tr>
            </thead>
            <tbody>
              {data?.map((gateway) => (
                <tr key={gateway?.id}>
                  <td className="gateway" width="15%">
                    <GatewayLogo name={gateway?.value} />
                  </td>
                  <td>
                    <strong>
                      {gateway?.name}
                    </strong>
                  </td>
                  <td>
                    <strong>
                      {gateway?.active ? t('Active') : t('Inactive')}
                    </strong>
                  </td>
                  <td width="15%" align="center">
                    <Link href={`/admin/deposits/gateway/${gateway?.value}`}>
                      <a className="action-btn">
                        <FiEdit />
                      </a>
                    </Link>
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

export default MethodsList;
