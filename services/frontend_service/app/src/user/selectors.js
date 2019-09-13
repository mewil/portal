import { get } from 'lodash';

const getGroup = (user, groupName) =>
  user && user.groups && user.groups.indexOf(groupName) !== -1;

export const getUserMetadata = (user) => ({
  isLoggedIn: user.isLoggedIn === true,
  isEmailVerified: user.isEmailVerified === true,
  isAdmin: getGroup(user, 'admin'),
});

export const getUser = (state) => get(state, 'user');
