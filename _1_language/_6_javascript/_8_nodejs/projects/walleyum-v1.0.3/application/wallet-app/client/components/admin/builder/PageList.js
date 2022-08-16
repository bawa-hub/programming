import { useRouter } from 'next/router';
import React from 'react';
import { Table } from 'react-bootstrap';
import { confirmAlert } from 'react-confirm-alert';
import { useTranslation } from 'react-i18next';
import { FiEdit, FiTrash } from 'react-icons/fi';
import usePages from '../../../data/usePages';
import { pageDelete } from '../../../lib/pageUpdate';
import Loader from '../../Loader';

const PageList = () => {
  const { data, loading } = usePages();
  const router = useRouter();
  const { t } = useTranslation();

  const deleteInit = (slug, name) => {
    confirmAlert({
      title: `${name}`,
      message: t('Are you sure to remove this page?'),
      buttons: [
        {
          label: t('Yes'),
          onClick: () => pageDelete(slug)
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
                <th scope="col">{t('Slug')}</th>
                <th scope="col">{t('Action')}</th>
              </tr>
            </thead>
            <tbody>
              {data?.map((page) => (
                <tr key={page?.id}>
                  <td>
                    <strong>
                      {page?.name}
                    </strong>
                  </td>
                  <td>
                    <strong>
                      {page?.slug}
                    </strong>
                  </td>
                  <td width="15%" align="center">
                    <div className="d-flex">
                      <button
                        type="button"
                        onClick={() => router.push(`/admin/builder/${page?.slug}`)}
                        className="action-btn"
                      >
                        <FiEdit />
                      </button>
                      {!(page?.slug === 'home') && (
                      <button
                        type="button"
                        onClick={() => deleteInit(page?.slug, page?.name)}
                        className="action-btn danger ml-10"
                      >
                        <FiTrash />
                      </button>
                      )}
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

export default PageList;
