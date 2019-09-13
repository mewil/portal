import { h } from 'react-hyperscript-helpers';
import styled from 'styled-components';

const Heart = styled.div`
  background-color: ${({ active, theme }) =>
    active ? theme.danger : theme.white};
  height: 30px;
  width: 30px;
  transform: rotate(-45deg);
  transition: background-color 0.4s;

  ::before,
  ::after {
    content: '';
    background-color: ${({ active, theme }) =>
      active ? theme.danger : theme.white};
    height: 30px;
    width: 30px;
    border-radius: 50%;
    position: absolute;
    transition: background-color 0.4s;
  }
  ::before {
    top: -15px;
  }
  ::after {
    left: 15px;
  }
`;

export const LikeButton = ({ active }) => h(Heart, { active });
