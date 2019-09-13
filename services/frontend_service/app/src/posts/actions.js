export const FETCH_FEED = 'posts/FETCH_FEED';
export const fetchFeedAction = () => ({
  type: FETCH_FEED,
});

export const FETCH_CREATE_POST = 'posts/FETCH_CREATE_POST';
export const fetchCreatePostAction = ({ caption, imageUrl, imageType }) => ({
  type: FETCH_CREATE_POST,
  payload: { caption, imageUrl, imageType },
});

export const ADD_POSTS = 'posts/ADD_POSTS';
export const addPostsAction = ({ posts }) => ({
  type: ADD_POSTS,
  payload: { posts },
});

export const ADD_FEED_PAGE = 'posts/ADD_FEED_PAGE';
export const addFeedPageAction = ({ posts, nextPage }) => ({
  type: ADD_FEED_PAGE,
  payload: { posts, nextPage },
});

export const FETCH_LIKE_POST = 'posts/FETCH_LIKE_POST';
export const fetchLikePostAction = ({ postId }) => ({
  type: FETCH_LIKE_POST,
  payload: { postId },
});
