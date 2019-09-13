import 'isomorphic-fetch';

export default ({ url }, reqOpts, action) =>
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
        data,
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
