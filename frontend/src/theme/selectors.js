import { get } from 'lodash';

export const getTheme = (state) => get(state, 'theme', {});
