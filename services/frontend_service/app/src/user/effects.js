import { select, call, put } from 'redux-saga/effects';
import { get } from 'lodash';

import { apiFetch, responseHasError } from '@portal/fetch';
import { getAuthToken } from '@portal/auth';

import { addUserAction } from './actions';

export function* onFetchUser({ payload: { userId } }) {
  const url = `/v1/user/${userId}`;
  const authToken = yield select(getAuthToken);
  const result = yield call(apiFetch, { url, authToken });

  if (responseHasError(result)) {
    return;
  }

  const user = get(result, 'data.user', {});

  yield put(addUserAction({ user }));
}
