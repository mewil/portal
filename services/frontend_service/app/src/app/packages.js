import { get, compact } from 'lodash';

const use = (modules = []) => ({
  reducers: combine(modules, 'reducers'),
  actions: combine(modules, 'actions'),
  sagas: combine(modules, 'sagas'),
  selectors: combine(modules, 'selectors'),
});

const combine = (modules = [], name = '') => {
  if (!name) return {};

  const propFromModules = compact(
    modules.map((module) => get(module, name, null)),
  );

  return propFromModules.reduce(
    (merged, property) =>
      !property
        ? merged
        : Object.keys(property).reduce(
            (accumulator, key) => ({
              ...accumulator,
              [key]: accumulator[key] || property[key],
            }),
            merged,
          ),
    {},
  );
};

export const packages = use([
  require('@portal/app'),
  require('@portal/auth'),
  require('@portal/home'),
  require('@portal/posts'),
  require('@portal/theme'),
  require('@portal/user'),
]);
