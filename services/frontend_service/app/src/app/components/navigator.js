import React from 'react';
import { h } from 'react-hyperscript-helpers';
import { connect } from 'react-redux';
import styled, { ThemeProvider } from 'styled-components';
import { withRouter } from 'react-router-dom';
import { IntlProvider } from 'react-intl';

import { getTheme } from '@portal/theme';

import { HeaderConn } from './header';
import { Footer } from './footer';

const Container = styled.div`
  margin-top: 80px;
`;

const Navigator = ({ theme, children }) =>
  h(IntlProvider, { locale: 'en' }, [
    h(ThemeProvider, { theme }, [
      h(HeaderConn),
      h(Container, [React.Children.toArray(children)]),
      h(Footer),
    ]),
  ]);

const mapStateToProps = (state) => ({
  theme: getTheme(state),
});

export const NavigatorConn = withRouter(connect(mapStateToProps)(Navigator));
