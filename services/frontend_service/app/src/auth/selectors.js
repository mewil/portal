import { get } from 'lodash';

export const getAuthUserId = (state) => get(state, 'auth.userId', null);

export const getAuthToken = (state) => get(state, 'auth.authToken', null);
