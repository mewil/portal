import { get } from 'lodash';

import { ADD_USER } from './actions';

export const users = (state = {}, action = {}) => {
  const { type, payload } = action;
  switch (type) {
    case ADD_USER: {
      const user = get(payload, 'user');
      const userId = get(user, 'userId');
      return {
        ...state,
        [userId]: user,
      };
    }
    default:
      return state;
  }
};
