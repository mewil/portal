import { call, put } from 'redux-saga/effects';
import { batchActions } from 'redux-batched-actions';
import { get } from 'lodash';

import { replace } from 'connected-react-router';

import { apiFetch, responseHasError } from '@portal/fetch';
import { addUserAction } from '@portal/user';

import { addAuthAction } from './actions';

export function* onFetchSignIn({ payload: { email, password } }) {
  const url = '/v1/auth/signin';
  const result = yield call(apiFetch, {
    url,
    method: 'POST',
    body: {
      email,
      password,
    },
  });

  if (responseHasError(result)) {
    return;
  }

  const user = get(result, 'data.user', {});
  const userId = get(user, 'userId', null);
  const authToken = get(result, 'data.token', null);

  yield put(batchActions([addAuthAction({ authToken, userId }), replace('/')]));
}

export function* onFetchSignUp({ payload: { email, username, password } }) {
  const url = '/v1/auth/signup';
  const result = yield call(apiFetch, {
    url,
    method: 'POST',
    body: {
      email,
      username,
      password,
    },
  });

  if (responseHasError(result)) {
    return;
  }

  const user = get(result, 'data.user', {});
  const userId = get(user, 'userId', null);
  const authToken = get(result, 'data.token', null);

  yield put(
    batchActions([
      addAuthAction({ authToken, userId }),
      addUserAction({ user }),
      replace('/'),
    ]),
  );
}
