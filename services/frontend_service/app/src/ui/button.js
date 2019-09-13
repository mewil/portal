import { debounce } from 'lodash';
import { h } from 'react-hyperscript-helpers';
import styled from 'styled-components';
import { Children } from 'react';
import PropTypes from 'prop-types';

export const Btn = styled.div`
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

const Inner = styled.div`
  display: flex;
  height: 100%;
  align-items: center;
  justify-content: ${({ justify }) => (justify ? 'space-between' : 'center')};
  user-select: none;
`;

export const Button = ({ children, ...props }) => {
  const childrenArray =
    typeof children === 'string'
      ? [h('span', children)]
      : Children.toArray(children);

  const onClick = debounce(
    (event) => {
      const domRect = event
        ? event.currentTarget.getBoundingClientRect()
        : null;
      props.onClick({ domRect });
    },
    1000,
    {
      leading: true,
      trailing: false,
    },
  );

  const onKeyDown = (ev) => {
    if (ev.key === 'Enter') onClick(ev);
  };

  return h(
    Btn,
    {
      ...props,
      onClick,
      onKeyDown,
    },
    [
      h(
        Inner,
        {
          justify: childrenArray.length > 1,
        },
        childrenArray,
      ),
    ],
  );
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
