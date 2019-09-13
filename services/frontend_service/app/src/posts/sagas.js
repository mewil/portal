import { takeEvery } from 'redux-saga/effects';

import { FETCH_FEED, FETCH_CREATE_POST, FETCH_LIKE_POST } from './actions';
import { onFetchFeed, onFetchCreatePost, onFetchLikePost } from './effects';

export function* fetchFeedSaga() {
  yield takeEvery(FETCH_FEED, onFetchFeed);
}

export function* fetchCreatePostSaga() {
  yield takeEvery(FETCH_CREATE_POST, onFetchCreatePost);
}

export function* fetchLikePostSaga() {
  yield takeEvery(FETCH_LIKE_POST, onFetchLikePost);
}
