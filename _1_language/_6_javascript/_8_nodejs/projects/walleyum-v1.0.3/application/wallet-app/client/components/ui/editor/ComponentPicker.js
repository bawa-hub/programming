import React from 'react';
import AboutSection from '../AboutSection';
import BannerSection from '../BannerSection';
import ContactSection from '../ContactSection';
import FaqSection from '../FaqSection';
import HeadingBlock from '../HeadingBlock';
import PageTitle from '../PageTitle';
import ServicesSection from '../ServicesSection';
import SolutionsSection from '../SolutionsSection';
import TextBlock from '../TextBlock';

const ComponentPicker = ({ type, data, info }) => {
  switch (type) {
    case 'PageTitle':
      return <PageTitle data={data} />;
    case 'TextBlock':
      return <TextBlock data={data} />;
    case 'HeadingBlock':
      return <HeadingBlock data={data} />;
    case 'ContactSection':
      return <ContactSection data={data} />;
    case 'AboutSection':
      return <AboutSection data={data} />;
    case 'BannerSection':
      return <BannerSection data={data} />;
    case 'ServicesSection':
      return <ServicesSection data={data} info={info} />;
    case 'SolutionsSection':
      return <SolutionsSection data={data} info={info} />;
    case 'FaqSection':
      return <FaqSection data={data} info={info} />;
    default:
      return <></>;
  }
};

export default ComponentPicker;
