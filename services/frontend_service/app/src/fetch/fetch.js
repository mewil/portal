import fetch from 'isomorphic-fetch';

import { forEach, camelCase, isPlainObject, isArray, get } from 'lodash';

const camelCaseObject = (obj) => {
  const result = {};
  forEach(obj, (v, k) => {
    result[camelCase(k)] =
      isPlainObject(v) || isArray(v) ? camelCaseObject(v) : v;
  });
  return result;
};

export const apiFetch = ({
  url,
  method = 'GET',
  body,
  form = false,
  authToken = '',
}) => {
  const headers = {};
  if (authToken !== '') {
    headers.Authorization = `Bearer ${authToken}`;
  }
  if (!form) {
    headers['Content-Type'] = 'application/json';
  }
  return fetch(url, {
    method,
    headers,
    credentials: 'same-origin',
    body:
      get(headers, 'Content-Type') === 'application/json'
        ? JSON.stringify(body)
        : body,
  }).then((res) =>
    res
      .json()
      .then((data) => ({
        headers: res.headers,
        status: res.status,
        statusText: res.statusText,
        url: res.url,
        data: camelCaseObject(get(data, 'data', {})),
      }))
      .catch((error) => ({
        headers: res.headers,
        status: res.status,
        statusText: res.statusText,
        url: res.url,
        error,
      })),
  );
};
