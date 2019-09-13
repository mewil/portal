import { h } from 'react-hyperscript-helpers';
import { ConnectedRouter } from 'connected-react-router';
import { connect } from 'react-redux';

import { GlobalStyle } from '@portal/theme';

import { Routes } from './routes';
import { NavigatorConn } from './navigator';
import { history } from '../store';
import { bootAction } from '../actions';

const Router = ({ boot = () => {} }) => {
  boot();
  return h(
    ConnectedRouter,
    {
      history,
    },
    [h(NavigatorConn, [h(Routes), h(GlobalStyle)])],
  );
};

const mapDispatchToProps = (dispatch) => ({
  boot: () => dispatch(bootAction()),
});

export const RouterConn = connect(
  null,
  mapDispatchToProps,
)(Router);
