import React, { useState } from 'react';
import {
  Col, Row
} from 'react-bootstrap';

const UserTab = ({
  title, description, defaultTab, children
}) => {
  const [tab, setTab] = useState(defaultTab || 0);

  return (
    <Row>
      <Col sm={3}>
        <h3 className="mb-30">{title}</h3>
        <p className="mb-30">{description}</p>
        {Array.isArray(children) ? (
          <div className="nav-pills">
            {children?.map((child, index) => (
              <button
                key={child?.props?.name}
                className={`bttn-mid mb-10 btn-grad nav-link text-left w-100 ${tab === index ? 'active' : ''}`}
                onClick={() => setTab(index)}
                type="button"
              >
                {child?.props?.icon}
                {child?.props?.name}
              </button>
            ))}
          </div>
        ) : (
          <div className="nav-pills">
            <button
              key={children?.props?.name}
              className={`bttn-mid mb-10 btn-grad nav-link text-left w-100 ${tab === 0 ? 'active' : ''}`}
              onClick={() => setTab(0)}
              type="button"
            >
              {children?.props?.icon}
              {children?.props?.name}
            </button>
          </div>
        )}
      </Col>
      <Col sm={9}>
        {Array.isArray(children) ? (
          <div>
            {children?.map((child, index) => (
              <div key={child?.props?.name} className={`tab-wrapper ${tab === index ? 'active' : 'not-active'}`}>
                {child?.props?.children}
              </div>
            ))}
          </div>
        ) : (
          <div>
            <div key={children?.props?.name} className={`tab-wrapper ${tab === 0 ? 'active' : 'not-active'}`}>
              {children?.props?.children}
            </div>
          </div>
        )}
      </Col>
    </Row>
  );
};

export default UserTab;
