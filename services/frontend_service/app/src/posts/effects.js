import { select, call, put } from 'redux-saga/effects';
import { batchActions } from 'redux-batched-actions';
import { get } from 'lodash';

import { apiFetch, responseHasError } from '@portal/fetch';
import { getAuthToken } from '@portal/auth';

import { addPostsAction, addFeedPageAction } from './actions';
import { getFeedNextPage } from './selectors';
import { replace } from 'connected-react-router';

export function* onFetchFeed() {
  const page = yield select(getFeedNextPage);

  const url = `/v1/post/?page=${page || 0}`;
  const authToken = yield select(getAuthToken);
  const result = yield call(apiFetch, { url, authToken });

  if (responseHasError(result)) return;

  const posts = get(result, 'data.feed', []);
  const nextPage = get(result, 'data.nextPage', 0);
  yield put(
    batchActions([
      addPostsAction({ posts }),
      addFeedPageAction({ posts, nextPage }),
    ]),
  );
}

export function* onFetchCreatePost({
  payload: { caption, imageUrl, imageType },
}) {
  const file = yield fetch(imageUrl)
    .then((r) => r.blob())
    .then((blobFile) => new File([blobFile], imageUrl, { type: imageType }));

  const formData = new FormData();
  formData.append('media', file);
  formData.append('caption', caption);

  const url = '/v1/post/';
  const authToken = yield select(getAuthToken);
  const result = yield call(apiFetch, {
    url,
    method: 'POST',
    form: true,
    authToken,
    body: formData,
  });

  if (responseHasError(result)) return;

  const post = get(result, 'data.post', {});
  yield put(batchActions([addPostsAction({ posts: [post] }), replace('/')]));
}

export function* onFetchLikePost({ payload: { postId } }) {
  const url = `/v1/post/${postId}/like`;
  const authToken = yield select(getAuthToken);
  const result = yield call(apiFetch, {
    url,
    method: 'POST',
    authToken,
  });

  if (responseHasError(result)) return;

  const post = get(result, 'data.post', {});
  yield put(addPostsAction({ posts: [post] }));
}
