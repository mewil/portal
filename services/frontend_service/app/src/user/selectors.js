import { get } from 'lodash';

export const getUser = (state, userId) => get(state, `users.${userId}`, null);
