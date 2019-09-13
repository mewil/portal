import { h } from 'react-hyperscript-helpers';
import { useCallback } from 'react';
import { useDropzone } from 'react-dropzone';
import { connect } from 'react-redux';
import styled from 'styled-components';

import { getTheme } from '@portal/theme';

import { Button } from './button';

const Container = styled.div`
  display: flex;
  justify-content: center;
`;

export const Dropzone = ({ theme, onDropFile = () => {} }) => {
  const onDrop = useCallback((files) => {
    onDropFile(files[0]);
  }, []);
  const { getRootProps } = useDropzone({ onDrop });

  return h(Container, { ...getRootProps() }, [
    h(
      Button,
      {
        hollow: true,
        color: theme.primary,
        style: {
          width: '400px',
          height: '400px',
        },
      },
      ['Drag A File Here'],
    ),
  ]);
};

const mapStateToProps = (state) => ({
  theme: getTheme(state),
});

export const DropzoneConn = connect(mapStateToProps)(Dropzone);
