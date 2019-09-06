import { createElement } from 'react';
import { render } from 'react-dom';
import { h } from 'react-hyperscript-helpers';
import { Provider } from 'react-redux';

import { configureStore } from './store';
import { RouterConn } from './components';

require('es6-promise').polyfill();

const App = () => {
  const store = configureStore();
  return h(Provider, { store }, [h(RouterConn)]);
};

render(createElement(App), document.getElementById('app'));
