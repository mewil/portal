import { createStore, applyMiddleware, compose, combineReducers } from 'redux';
import { createBrowserHistory } from 'history';
import { connectRouter, routerMiddleware } from 'connected-react-router';
import { spawn, all } from 'redux-saga/effects';
import { enableBatching, batchDispatchMiddleware } from 'redux-batched-actions';
import { END } from 'redux-saga';

import { createReduxSagaMiddleware } from './middleware';
import { packages } from './packages';
console.log(packages);

export const history = createBrowserHistory();

const sagaMiddleware = createReduxSagaMiddleware();

const prepSagas = (sagas = {}, options = []) =>
  Object.values(sagas).map((saga) => spawn(saga, ...options));

const sagaCreator = (sagas) =>
  function* rootSaga(...options) {
    yield all(prepSagas(sagas, options));
  };

const middleware = [
  sagaMiddleware,
  batchDispatchMiddleware,
  routerMiddleware(history),
];

const enhancer = compose(applyMiddleware(...middleware));

const rootReducers = enableBatching(
  combineReducers({ ...packages.reducers, router: connectRouter(history) }),
);

const rootSagas = sagaCreator(packages.sagas);

export const configureStore = () => {
  const store = createStore(rootReducers, enhancer);
  store.close = () => store.dispatch(END);
  sagaMiddleware.run(rootSagas);
  return store;
};
