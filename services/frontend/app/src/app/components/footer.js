import styled from 'styled-components';
import { h, div } from 'react-hyperscript-helpers';

const Container = styled.div`
  position: relative;
  display: flex;
  padding: 1rem;
  align-content: space-between;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;
  flex-direction: column;
`;

export const Footer = () =>
  h(Container, [div([`© ${new Date().getFullYear()} Michael Wilson`])]);
