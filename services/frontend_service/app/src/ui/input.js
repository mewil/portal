import { h } from 'react-hyperscript-helpers';
import styled from 'styled-components';
import { get } from 'lodash';

import { TextFaded } from './typography';

export const InputContainer = styled.div`
  display: flex;
  flex-direction: column;
  position: relative;
`;

const Row = styled.div`
  display: flex;
  align-items: center;
`;

const InputEl = styled.input`
  padding: 12px;
  margin: 6px 0;
  font-family: ${({ theme }) => theme.primaryFont};
  font-size: 15px;
  line-height: 24px;
  color: ${({ theme }) => theme.darkGray};
  border-radius: ${({ rounded }) => (rounded ? '50px' : '3px')};
  box-sizing: border-box;
  border: solid 1px
    ${({ theme, success, error }) => {
      if (error) return theme.danger;
      if (success) return theme.success;
      return theme.lightGray;
    }};
  outline: none;
  &:focus {
    border-color: ${({ theme }) => theme.info};
  }
  &::placeholder {
    color: ${({ theme }) => theme.mediumGray};
  }
  &:focus::placeholder {
    color: transparent;
  }
  &:disabled {
    color: ${({ theme }) => theme.mediumGray};
    background-color: ${({ theme }) => theme.lightGray};
  }
  height: 48px;
  ${({ type }) => {
    if (type === 'password')
      return `
      letter-spacing: 3px;
      font-weight: 700;
    `;
  }};
`;

export const Label = styled(TextFaded.withComponent('label'))`
  margin: 6px 2px;
  margin-bottom: 0;
`;

const RequiredLabel = styled(TextFaded)`
  font-size: 10px;
  float: right;
  ${({ theme, error }) => (error ? `color: ${theme.danger};` : '')};
`;

const ErrorEl = styled(Label)`
  position: absolute;
  top: 54px;
  z-index: 1;
  width: -webkit-fill-available;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  margin-bottom: 6px;
  padding: 12px 0;
  background-color: ${({ theme }) => theme.white};
  box-shadow: 0 2px 2px 0 rgba(133, 132, 155, 0.16),
    0 1px 1px 0 rgba(133, 132, 155, 0.1);
`;

export const Input = ({
  label,
  required,
  innerRef,
  focused,
  setFocus,
  containerStyles = {},
  icon,
  ...props
}) => {
  const { error } = props;
  const showErrorDropdown = get(error, 'length', 0) && !focused;

  return h(
    InputContainer,
    {
      style: containerStyles,
    },
    [
      required
        ? h(Label, [
            required
              ? h(RequiredLabel, { error: showErrorDropdown }, 'Required')
              : null,
          ])
        : null,
      h(Row, [
        h(InputEl, {
          innerRef,
          onBlur: () => setFocus(false),
          onFocus: () => setFocus(true),
          ...props,
        }),
      ]),
      showErrorDropdown ? h(ErrorEl, [h('span', error)]) : null,
    ],
  );
};
