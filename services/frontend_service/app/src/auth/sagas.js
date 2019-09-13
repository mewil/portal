import { takeEvery } from 'redux-saga/effects';

import { FETCH_SIGNIN, FETCH_SIGNUP } from './actions';
import { onFetchSignIn, onFetchSignUp } from './effects';

export function* fetchSignInSaga() {
  yield takeEvery(FETCH_SIGNIN, onFetchSignIn);
}

export function* fetchSignUpSaga() {
  yield takeEvery(FETCH_SIGNUP, onFetchSignUp);
}
