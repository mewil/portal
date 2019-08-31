import React from 'react';
import { h } from 'react-hyperscript-helpers';
import { connect } from 'react-redux';
import { ThemeProvider } from 'styled-components';
import { withRouter } from 'react-router-dom';
import { IntlProvider } from 'react-intl';

import { getTheme } from '@portal/theme';

const Navigator = ({ theme, children }) =>
  h(IntlProvider, { locale: 'en' }, [
    h(ThemeProvider, { theme }, [h('div', [React.Children.toArray(children)])]),
  ]);

const mapStateToProps = (state) => ({
  theme: getTheme(state),
});

export const NavigatorConn = withRouter(connect(mapStateToProps)(Navigator));
