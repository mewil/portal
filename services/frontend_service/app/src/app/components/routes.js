import { h } from 'react-hyperscript-helpers';
import { Route, Switch } from 'react-router';

import { routes } from '../constants';
import { HomePageConn } from '@portal/home';
import { NewPostPageConn } from '@portal/posts';

export const Routes = () =>
  h(Switch, [
    h(Route, {
      exact: true,
      path: routes.HOME,
      component: HomePageConn,
    }),
    h(Route, {
      exact: true,
      path: routes.NEW,
      component: NewPostPageConn,
    }),
    h(Route, { component: HomePageConn }),
  ]);
