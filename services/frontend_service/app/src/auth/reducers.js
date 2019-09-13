import { get } from 'lodash';

import { ADD_AUTH } from './actions';

export const auth = (
  state = { authToken: null, userId: null },
  action = {},
) => {
  const { type, payload } = action;
  switch (type) {
    case ADD_AUTH: {
      const authToken = get(payload, 'authToken', null);
      const userId = get(payload, 'userId', null);
      return {
        ...state,
        authToken,
        userId,
      };
    }
    default:
      return state;
  }
};
