import React, { useEffect, useState } from 'react';
import { Dropdown } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { BiFlag } from 'react-icons/bi';

const LanguageSwitch = () => {
  const { i18n } = useTranslation();
  const [selectedLang, setSelectedLang] = useState();

  const languages = [
    {
      code: 'en',
      label: 'English'
    },
    {
      code: 'fr',
      label: 'Français'
    },
    {
      code: 'es',
      label: 'Español'
    }
  ];

  useEffect(() => {
    const lang = localStorage.getItem('lang') || 'en';
    const findLang = languages.find((lng) => lng.code === lang);
    setSelectedLang(findLang);
  }, []);

  const setLang = (lng) => {
    setSelectedLang(lng);
    localStorage.setItem('lang', lng?.code);
    localStorage.setItem('langLabel', lng?.label);
    i18n.changeLanguage(lng?.code);
  };

  return (
    <div>
      <Dropdown>
        <Dropdown.Toggle className="bttn-small btn-blue btn-lang" variant="link">
          <BiFlag />
          {selectedLang?.label}
        </Dropdown.Toggle>
        <Dropdown.Menu>
          {languages.map((lng) => (
            <Dropdown.Item key={lng?.code} onClick={() => setLang(lng)}>
              {lng?.label}
            </Dropdown.Item>
          ))}
        </Dropdown.Menu>
      </Dropdown>
    </div>
  );
};

export default LanguageSwitch;
