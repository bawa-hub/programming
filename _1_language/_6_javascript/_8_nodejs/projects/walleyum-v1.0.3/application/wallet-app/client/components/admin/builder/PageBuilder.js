/* eslint-disable max-len */
/* eslint-disable jsx-a11y/label-has-associated-control */
import React, { useEffect, useState } from 'react';
import { Spinner } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { BiSave } from 'react-icons/bi';
import SortableList from 'react-sortable-dnd-list';
import { usePageBySlug } from '../../../data/usePages';
import { pageUpdate } from '../../../lib/pageUpdate';
import Loader from '../../Loader';
import BlockAdd from '../../ui/editor/BlockAdd';
import ComponentsField from '../../ui/editor/ComponentsField';
import SortableSection from './SortableSection';

const PageBuilder = ({ slug }) => {
  const [name, setName] = useState();
  const [slugBuild, setSlugBuild] = useState();
  const [items, setItems] = useState([]);
  const [actionLoader, setActionLoader] = useState(false);
  const { t } = useTranslation();

  const { data, loading } = usePageBySlug(slug);

  const handleAdd = (item) => {
    const arr = [...items];
    arr.push(item);
    setItems(arr);
  };

  const handleDelete = (id) => {
    const arr = items.filter((itm) => itm.id !== id);
    setItems(arr);
  };

  const handleEdit = async (id, field, value) => {
    const index = items.findIndex((itm) => itm.id === id);
    const arr = [...items];
    arr[index].data[field] = value;
    setItems(arr);
  };

  const handleSubmit = () => {
    const slugRefined = slugBuild.replace(/[^\w\s]/gi, '');
    const params = {
      name,
      slug: slugRefined,
      content: items
    };
    pageUpdate(slug, params, setActionLoader);
  };

  useEffect(() => {
    const jsonItems = data ? data?.content : [];
    setName(data?.name);
    setSlugBuild(data?.slug);
    setItems(jsonItems);
  }, [loading]);

  if (loading) {
    return <Loader height="200px" />;
  }

  return (
    <>
      <div className="row">
        <div className="col-md-4">
          <div className="basic-card">
            <h4 className="box-title">{t('Page Settings')}</h4>
            <div className="settings-box">
              <form>
                <div className="single-profile">
                  <label htmlFor="pageTitle">{t('Page Title')}</label>
                  <input
                    defaultValue={name}
                    onChange={(e) => setName(e.target.value)}
                    id="pageTitle"
                    type="text"
                    placeholder={t('Title')}
                    required
                  />
                </div>
                <div className="single-profile">
                  <label htmlFor="pageSlug">{t('Slug')}</label>
                  <input
                    defaultValue={slugBuild}
                    onChange={(e) => setSlugBuild(e.target.value)}
                    id="pageSlug"
                    type="text"
                    placeholder={t('Slug')}
                    required
                    disabled={slugBuild === 'home'}
                  />
                </div>
              </form>
            </div>
          </div>
          <div className="basic-card mt-20">
            <h4 className="box-title">{t('Pick Components')}</h4>
            <div className="component-wrapper">
              {ComponentsField.map((field) => <BlockAdd key={field.component} field={field} handleAdd={handleAdd} />)}
            </div>
          </div>
        </div>
        <div className="col-md-8">
          <div className="basic-card">
            <h4 className="box-title">{t('Builder')}</h4>
            <SortableList
              className="list"
              itemComponent={SortableSection}
              itemComponentProps={{ handleDelete, handleEdit }}
              value={items}
              index={0}
              onChange={(value) => setItems(value)}
            />
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
          </div>
        </div>
      </div>
    </>
  );
};
export default PageBuilder;
