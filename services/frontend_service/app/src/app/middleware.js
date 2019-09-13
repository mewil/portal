import createSagaMiddleware from 'redux-saga';
import { BATCH } from 'redux-batched-actions';

export const createReduxSagaMiddleware = () => {
  const onError = (err) => {
    console.log(err);
  };
  const emitter = (emit) => (action) => {
    const { type, payload } = action;
    if (type === BATCH && Array.isArray(payload)) {
      payload.forEach(emit);
      return;
    }

    emit(action);
  };

  return createSagaMiddleware({
    onError,
    emitter,
  });
};
