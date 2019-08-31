import styled from 'styled-components';
import { Helmet } from 'react-helmet';
import { NavLink } from 'react-router-dom';
import { connect } from 'react-redux';
import { h, div } from 'react-hyperscript-helpers';

// eslint-disable-next-line no-unused-vars
import { theme as globalTheme, devices, getTheme } from '@portal/theme';

import { routes } from '../constants';

const Wrapper = styled.div`
  margin: 0 px 16%;
  background-color: white;
  ${devices.small`
        margin: 0px 8%;
    `};
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  padding-top: 15px
  padding-bottom: 15px
  z-index: 100;
  display: flex;
  height: 80px;
  align-items: center;
  justify-content: flex-start;
`;

const NavContainer = styled.div`
  display: flex;
  align-items: center;
  justify-content: flex-end;
`;

const StyledNavLink = styled(NavLink)`
  color: ${({ theme }) => theme.primary};
  font-size: 16px;
  padding: 2px 20px;
  margin: 10px 0 10px 15px;
  border-radius: 5px;
  text-decoration: none;
  transition: all 0.3s;
  text-transform: uppercase;
  &:hover {
    font-weight: bold;
  }
  &:first-child {
    margin: 0;
    margin-left: 15px;
  }
`;

const Header = () =>
  div(
    h(Helmet, [h('title', 'Portal')]),
    h(Wrapper, [
      h(NavContainer, { right: false, isOpen: false }, [
        h(StyledNavLink, { to: routes.HOME }, 'Home'),
      ]),
    ]),
  );

const mapStateToProps = (state) => ({
  theme: getTheme(state),
});

export const HeaderConn = connect(mapStateToProps)(Header);
