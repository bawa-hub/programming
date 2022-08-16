import { useRouter } from 'next/router';
import React from 'react';
import { BiArrowBack } from 'react-icons/bi';

const EditorHeader = ({ name }) => {
  const router = useRouter();

  return (
    <div className="editor-header">
      <button type="button" onClick={() => router.back()}>
        <BiArrowBack />
      </button>
      <h2 className="mb-30">
        {name}
      </h2>
    </div>
  );
};

export default EditorHeader;
