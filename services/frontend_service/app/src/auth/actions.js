export const FETCH_SIGNIN = 'posts/FETCH_SIGNIN';
export const fetchSignInAction = ({ email, password }) => ({
  type: FETCH_SIGNIN,
  payload: { email, password },
});

export const FETCH_SIGNUP = 'posts/FETCH_SIGNUP';
export const fetchSignUpAction = ({ email, username, password }) => ({
  type: FETCH_SIGNUP,
  payload: { email, username, password },
});

export const ADD_AUTH = 'posts/ADD_AUTH';
export const addAuthAction = ({ authToken, userId }) => ({
  type: ADD_AUTH,
  payload: { authToken, userId },
});
