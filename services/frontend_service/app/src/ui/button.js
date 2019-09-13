import { h } from 'react-hyperscript-helpers';
import styled from 'styled-components';
import PropTypes from 'prop-types';

export const Btn = styled.div`
  cursor: pointer;
  box-sizing: border-box;
  display: inline-block;
  padding: 10px;
  min-width: 140px;
  background-color: ${({ danger, warning, success, hollow, theme }) => {
    if (hollow) {
      return 'transparent';
    }
    if (danger) return theme.danger;
    if (success) return theme.success;
    if (warning) return theme.warning;
    return theme.info;
  }};
  border: 1px solid;
  border-color: ${({ danger, warning, success, theme }) => {
    if (danger) return theme.danger;
    if (success) return theme.success;
    if (warning) return theme.warning;
    return theme.info;
  }};
  border-radius: ${({ rounded }) => (rounded ? '50px' : '3px')};
  font-size: 15px;
  text-align: center;
  color: ${({ danger, warning, success, hollow, theme }) => {
    if (hollow) {
      if (danger) return theme.danger;
      if (success) return theme.success;
      if (warning) return theme.warning;
      return theme.info;
    }
    return theme.white;
  }};
  ${({ square }) => {
    if (square) {
      return `
        width: 47px;
        height: 47px;
        min-width: 47px;
        min-height: 47px;
        padding: 0;
      `;
    }
  }};
`;

export const Button = ({ onClick, ...props }) => {
  const onKeyDown = (ev) => {
    if (ev.key === 'Enter') onClick(ev);
  };

  return h(Btn, {
    ...props,
    onClick,
    onKeyDown,
  });
};

Button.propTypes = {
  info: PropTypes.bool,
  danger: PropTypes.bool,
  warning: PropTypes.bool,
  success: PropTypes.bool,
  hollow: PropTypes.bool,
  rounded: PropTypes.bool,
  square: PropTypes.bool,
};
