import React from 'react';
import {
  BiDetail, BiHeading, BiText
} from 'react-icons/bi';

const fields = [
  {
    label: 'Heading Block',
    component: 'HeadingBlock',
    icon: <BiHeading />,
    fields: [
      {
        label: 'Content',
        name: 'content',
        type: 'text',
      },
    ]
  },
  {
    label: 'Text Block',
    component: 'TextBlock',
    icon: <BiText />,
    fields: [
      {
        label: 'Content',
        name: 'content',
        type: 'textarea',
      },
    ]
  },
  {
    label: 'Page Title',
    component: 'PageTitle',
    icon: <BiText />,
    fields: [
      {
        label: 'Heading',
        name: 'content',
        type: 'text',
      },
    ]
  },
  {
    label: 'Contact Section',
    component: 'ContactSection',
    icon: <BiDetail />,
    fields: [
      {
        label: 'Section Title',
        name: 'sectionTitle',
        type: 'text',
      },
      {
        label: 'Section Subtitle',
        name: 'sectionSub',
        type: 'text',
      },
      {
        label: 'Email',
        name: 'email',
        type: 'email',
      },
      {
        label: 'Phone',
        name: 'phone',
        type: 'text',
      },
      {
        label: 'Address',
        name: 'address',
        type: 'text',
      },
    ]
  },
  {
    label: 'About Section',
    component: 'AboutSection',
    icon: <BiDetail />,
    fields: [
      {
        label: 'Section Title',
        name: 'sectionTitle',
        type: 'text',
      },
      {
        label: 'Content',
        name: 'content',
        type: 'textarea',
      },
      {
        label: 'Image',
        name: 'image',
        type: 'image',
      },
    ]
  },
  {
    label: 'Banner Section',
    component: 'BannerSection',
    icon: <BiDetail />,
    fields: [
      {
        label: 'Section Title',
        name: 'sectionTitle',
        type: 'text',
      },
      {
        label: 'Content',
        name: 'content',
        type: 'textarea',
      },
      {
        label: 'Button Label',
        name: 'buttonLabel',
        type: 'text',
      },
      {
        label: 'Button URL',
        name: 'buttonUrl',
        type: 'text',
      },
      {
        label: 'Cover Image',
        name: 'cover',
        type: 'image',
      },
      {
        label: 'Side Image',
        name: 'side',
        type: 'image',
      },
    ]
  },
  {
    label: 'Services Section',
    component: 'ServicesSection',
    icon: <BiDetail />,
    fields: [
      {
        label: 'Section Title',
        name: 'sectionTitle',
        type: 'text',
      },
      {
        label: 'Section Subtitle',
        name: 'sectionSub',
        type: 'text',
      },
    ]
  },
  {
    label: 'Solutions Section',
    component: 'SolutionsSection',
    icon: <BiDetail />,
    fields: [
      {
        label: 'Section Title',
        name: 'sectionTitle',
        type: 'text',
      },
      {
        label: 'Section Subtitle',
        name: 'sectionSub',
        type: 'text',
      },
    ]
  },
  {
    label: 'Faq Section',
    component: 'FaqSection',
    icon: <BiDetail />,
    fields: [
      {
        label: 'Section Title',
        name: 'sectionTitle',
        type: 'text',
      },
      {
        label: 'Section Subtitle',
        name: 'sectionSub',
        type: 'text',
      },
      {
        label: 'Button Label',
        name: 'buttonLabel',
        type: 'text',
      },
      {
        label: 'Button URL',
        name: 'buttonUrl',
        type: 'text',
      },
    ]
  },
];

export default fields;
