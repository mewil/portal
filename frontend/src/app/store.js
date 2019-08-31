import { createStore, applyMiddleware, compose, combineReducers } from 'redux';
import { spawn, all } from 'redux-saga/effects';
import { enableBatching, batchDispatchMiddleware } from 'redux-batched-actions';
import { routerMiddleware } from 'react-router-redux';
import { END } from 'redux-saga';

import { createReduxSagaMiddleware } from './middleware';
import { packages } from './packages';
console.log(packages);

const sagaMiddleware = createReduxSagaMiddleware();

const prepSagas = (sagas = {}, options = []) =>
  Object.values(sagas).map((saga) => spawn(saga, ...options));

const sagaCreator = (sagas) =>
  function* rootSaga(...options) {
    yield all(prepSagas(sagas, options));
  };

const middleware = [sagaMiddleware, batchDispatchMiddleware, routerMiddleware];

const enhancer = compose(applyMiddleware(...middleware));

const rootReducers = enableBatching(combineReducers(packages.reducers));

const rootSagas = sagaCreator(packages.sagas);

export const configureStore = () => {
  const store = createStore(rootReducers, enhancer);
  store.close = () => store.dispatch(END);
  sagaMiddleware.run(rootSagas);
  return store;
};
