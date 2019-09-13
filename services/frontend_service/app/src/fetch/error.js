import { get } from 'lodash';

export const filterStatus = (response) => {
  const status = get(response, 'status', 0);
  return status >= 200 && status < 300;
};

export const responseHasError = (response) => !filterStatus(response);
