import React, { useEffect, useState } from 'react';
import { Spinner } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { BiPlusCircle, BiSave } from 'react-icons/bi';
import SortableList from 'react-sortable-dnd-list';
import { v1 as uuidv1 } from 'uuid';
import useInfo from '../../../../data/useInfo';
import { repeaterUpdate } from '../../../../lib/menuUpdate';
import Loader from '../../../Loader';
import SortableFaq from './SortableFaq';

const FaqBuilder = () => {
  const [items, setItems] = useState([]);
  const [actionLoader, setActionLoader] = useState(false);

  const { data, loading } = useInfo();

  const { t } = useTranslation();

  const handleAdd = (item) => {
    const arr = [...items];
    arr.push(item);
    setItems(arr);
  };

  const handleDelete = (id) => {
    const arr = items.filter((itm) => itm.id !== id);
    setItems(arr);
  };

  const handleEdit = (id, field, value) => {
    const index = items.findIndex((itm) => itm.id === id);
    const arr = [...items];
    arr[index][field] = value;
    setItems(arr);
  };

  useEffect(() => {
    const jsonItems = data ? JSON.parse(data?.faq?.param1) : [];
    setItems(jsonItems);
  }, [data]);

  const handleSubmit = () => {
    const params = {
      param1: JSON.stringify(items)
    };
    repeaterUpdate('faq', params, setActionLoader);
  };

  if (loading) {
    return <Loader height="200px" />;
  }

  return (
    <>
      <SortableList
        className="list"
        itemComponent={SortableFaq}
        itemComponentProps={{ handleDelete, handleEdit }}
        value={items}
        index={0}
        onChange={(value) => setItems(value)}
      />
      <button
        type="button"
        className="bttn-mid btn-primary btn-new mt-10"
        onClick={() => handleAdd({
          id: uuidv1(), title: 'New Item', content: 'Lorem Ipsum'
        })}
      >
        <BiPlusCircle />
        {t('Add New FAQ')}
      </button>
      <button
        type="button"
        className="bttn-mid btn-ylo btn-new"
        onClick={() => handleSubmit()}
        disabled={actionLoader}
      >
        {actionLoader ? (
          <>
            <Spinner animation="border" role="status" size="sm" />
            {' '}
            {t('Save Changes')}
          </>
        ) : (
          <>
            <BiSave />
            {t('Save Changes')}
          </>
        )}
      </button>
    </>
  );
};
export default FaqBuilder;
