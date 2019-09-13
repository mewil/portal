export const ADD_USER = 'posts/ADD_USER';
export const addUserAction = ({ user }) => ({
  type: ADD_USER,
  payload: { user },
});

export const FETCH_USER = 'posts/FETCH_USER';
export const fetchUserAction = ({ userId }) => ({
  type: FETCH_USER,
  payload: { userId },
});
