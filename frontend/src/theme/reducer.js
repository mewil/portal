import { theme as defaultTheme } from './theme';

export const theme = (state = defaultTheme, action) => {
  const { type } = action;
  switch (type) {
    default:
      return state;
  }
};
