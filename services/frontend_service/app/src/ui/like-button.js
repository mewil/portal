import { h } from 'react-hyperscript-helpers';
import styled from 'styled-components';

const Heart = styled.div`
  background-color: ${({ active, theme }) =>
    active ? theme.danger : theme.lightGray};
  height: 16px;
  width: 16px;
  transform: rotate(-45deg);
  transition: background-color 0.4s;

  ::before,
  ::after {
    content: '';
    background-color: ${({ active, theme }) =>
      active ? theme.danger : theme.lightGray};
    height: 16px;
    width: 16px;
    border-radius: 50%;
    position: absolute;
    transition: background-color 0.4s;
  }
  ::before {
    top: -8px;
  }
  ::after {
    left: 8px;
  }
`;

export const LikeButton = ({ theme, active }) => h(Heart, { theme, active });
