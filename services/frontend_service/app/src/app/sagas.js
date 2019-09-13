import { takeEvery } from 'redux-saga/effects';

import { BOOT } from './actions';
import { onBoot } from './effects';

export function* bootSaga() {
  yield takeEvery(BOOT, onBoot);
}
