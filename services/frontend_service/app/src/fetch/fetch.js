import 'isomorphic-fetch';

import { forEach, camelCase, isPlainObject, isArray } from 'lodash';

const camelCaseObject = (obj) => {
  const result = {};
  forEach(obj, (v, k) => {
    result[camelCase(k)] =
      isPlainObject(v) || isArray(v) ? camelCaseObject(v) : v;
  });
  return result;
};

export const apiFetch = ({ url }, reqOpts, action) =>
  fetch(url, reqOpts).then((res) =>
    res
      .json()
      .then((data) => ({
        headers: res.headers,
        status: res.status,
        statusText: res.statusText,
        url: res.url,
        reqOpts,
        action,
        data: camelCaseObject(data),
      }))
      .catch((error) => ({
        headers: res.headers,
        status: res.status,
        statusText: res.statusText,
        url: res.url,
        reqOpts,
        action,
        error,
      })),
  );
