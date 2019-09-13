import styled from 'styled-components';
import { h } from 'react-hyperscript-helpers';

const Container = styled.div`
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 1rem;
`;

const Text = styled.p`
  text-align: center;
`;

export const Footer = () =>
  h(Container, [h(Text, [`© ${new Date().getFullYear()} Michael Wilson`])]);
