import { get, pick, values } from 'lodash';

export const getPosts = (state) => get(state, 'posts', {});

export const getFeed = (state) => get(state, 'feed', {});

export const getFeedPostIds = (state) => get(getFeed(state), 'postIds', []);

export const getFeedNextPage = (state) => get(getFeed(state), 'nextPage', 0);

export const getFeedPosts = (state) => {
  const feedIds = getFeedPostIds(state);
  return values(pick(getPosts(state), feedIds)).sort(
    (a, b) => feedIds.indexOf(a) - feedIds.indexOf(b),
  );
};
