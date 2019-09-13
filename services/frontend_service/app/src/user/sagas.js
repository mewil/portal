import { takeEvery } from 'redux-saga/effects';

import { FETCH_USER } from './actions';
import { onFetchUser } from './effects';

export function* fetchUserSaga() {
  yield takeEvery(FETCH_USER, onFetchUser);
}
