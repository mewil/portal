import { get, values } from 'lodash';

import { ADD_POSTS, ADD_FEED_PAGE } from './actions';

export const posts = (state = {}, action = {}) => {
  const { type, payload } = action;
  switch (type) {
    case ADD_POSTS: {
      const newPosts = values(get(payload, 'posts', {})).reduce(
        (results, p) => ({
          ...results,
          [p.postId]: p,
        }),
        {},
      );
      return {
        ...state,
        ...newPosts,
      };
    }

    default:
      return state;
  }
};

export const feed = (state = { postIds: [], nextPage: 0 }, action = {}) => {
  const { type, payload } = action;
  switch (type) {
    case ADD_FEED_PAGE: {
      const postIds = values(get(payload, 'posts', [])).map(
        ({ postId }) => postId,
      );
      const nextPage = get(payload, 'nextPage', 0);
      return {
        ...state,
        postIds: [...state.postIds, ...postIds],
        nextPage,
      };
    }
    default:
      return state;
  }
};
