const path = require('path');
const { get } = require('lodash');
const babylon = require('babylon');
const astWalk = require('babylon-walk');

const { read, readdir, task, stat, cleanGlob, flatten } = require('./util');
const {
  SUCCESS,
  DEPENDENCY_MISSING_IN_PACKAGE_JSON,
  NOT_PORTAL_IMPORT,
  UNNECESSARY_PACKAGE_DEPENDENCY,
} = require('./constants');

const defaultPkg = {
  file: '',
  contents: '',
};
const defaultOutput = {
  type: '',
  importName: '',
  file: '',
  pkg: defaultPkg,
};

function* run(dirs) {
  const topLevelDirs = yield dirs.map((dir) =>
    task(getTopLevelDir, cleanGlob(dir)),
  );
  const files = yield flatten(topLevelDirs)
    .filter(Boolean)
    .map(({ file, pkg }) => task(walk, file, pkg));
  const jsFiles = flatten(files).filter(({ file }) => /\.js$/.test(file));
  const trees = yield jsFiles.map((fileObj) => task(parse, fileObj));
  const flatTrees = flatten(trees);

  // Lint rules
  const dependsErrors = flatTrees.map((props) =>
    ensureDependencyInPackageJson(props),
  );

  const state = new Set();
  flatTrees.forEach((props) =>
    ensureOnlyNecessaryDependencyInPackageJson({ ...props, state }),
  );
  const unnecessaryErrors = detectUnnecessaryPackagesInPackageJson(state);

  const errors = [...flatten(dependsErrors), ...unnecessaryErrors];
  return errors;
}

/*
 * Get list of packages and their pacakge.json files
 */
function* getTopLevelDir(dir) {
  const files = yield readdir(dir);
  const stats = yield files.map((file) => getStats(dir, file));

  const results = yield stats
    .filter(({ statFile }) => statFile.isDirectory())
    .map(({ statFile, file }) => {
      const filePath = path.join(dir, file);
      const pkgFilePath = path.join(filePath, 'package.json');

      return getStats(filePath, 'package.json')
        .then((res) => {
          if (res.type !== SUCCESS) {
            return Promise.reject(res);
          }
        })
        .then(() => read(pkgFilePath))
        .then((contents) => ({
          type: SUCCESS,
          statFile,
          file: filePath,
          pkg: {
            file: pkgFilePath,
            contents: JSON.parse(contents.toString()),
          },
        }))
        .catch(console.log);
    });

  return results;
}

function* parse({ file, pkg }) {
  const contents = yield read(file);
  const ast = babylon.parse(contents.toString(), {
    sourceType: 'module',
    plugins: [
      'flow',
      'exportExtensions',
      'classProperties',
      'jsx',
      'objectRestSpread',
    ],
    allowImportExportEverywhere: true,
  });

  return { file, pkg, ast };
}

function ensureOnlyNecessaryDependencyInPackageJson({ ast, pkg, state }) {
  /* eslint-disable no-param-reassign */
  const visitors = {
    ImportDeclaration: (node, curState) => {
      const importName = node.source.value;
      if (!/^@portal/.test(importName)) return;
      state.add(importName);
      if (!state.hasOwnProperty(pkg.file)) {
        curState[pkg.file] = { pkg, actual: new Set() };
      }

      curState[pkg.file].actual.add(importName);
    },
    CallExpression: (node, curState) => {
      if (node.callee.name !== 'require') {
        return;
      }

      const importName = node.arguments[0].value;
      if (!importName) return;

      if (!/^@portal/.test(importName)) return;
      if (!state.hasOwnProperty(pkg.file)) {
        curState[pkg.file] = { pkg, actual: new Set() };
      }

      curState[pkg.file].actual.add(importName);
    },
  };

  astWalk.recursive(ast, visitors, state);

  return state;
}

function detectUnnecessaryPackagesInPackageJson(state) {
  return Object.values(state).reduce((acc, check) => {
    const ret = [];
    const { pkg } = check;
    const expected = Object.keys(get(pkg, 'contents.dependencies', {}));

    for (let i = 0; i < expected.length; i++) {
      const expectedDepend = expected[i];
      if (check.actual.has(expectedDepend)) {
        continue;
      }

      const err = {
        type: UNNECESSARY_PACKAGE_DEPENDENCY,
        importName: expectedDepend,
        file: '',
        pkg,
      };

      ret.push(err);
    }

    return [...acc, ...ret];
  }, []);
}

function ensureDependencyInPackageJson({ ast, file, pkg }) {
  const visitors = {
    ImportDeclaration: (node, state) => {
      const importName = node.source.value;
      const result = checkDependency(importName, file, pkg);
      if (result.type === DEPENDENCY_MISSING_IN_PACKAGE_JSON) {
        state.push(result);
      }
    },
    CallExpression: (node, state) => {
      if (node.callee.name !== 'require') {
        return;
      }

      const importName = node.arguments[0].value;
      if (!importName) return;

      const result = checkDependency(importName, file, pkg);
      if (result.type === DEPENDENCY_MISSING_IN_PACKAGE_JSON) {
        state.push(result);
      }
    },
  };

  const errors = [];
  astWalk.recursive(ast, visitors, errors);
  return errors;
}

function checkDependency(importName, file, pkg = defaultPkg) {
  if (!/^@portal/.test(importName))
    return { ...defaultOutput, type: NOT_PORTAL_IMPORT };

  const err = {
    type: DEPENDENCY_MISSING_IN_PACKAGE_JSON,
    importName,
    file,
    pkg,
  };

  const depends = get(pkg, 'contents.dependencies', {});

  if (!depends.hasOwnProperty(importName)) {
    return err;
  }

  return { ...defaultOutput, type: SUCCESS };
}

function* walk(dir, pkg = defaultPkg) {
  const files = yield readdir(dir);
  const stats = yield files.map((file) => getStats(dir, file));

  const recurFiles = yield stats.map(({ statFile, file }) => {
    const filePath = path.join(dir, file);

    if (statFile.isDirectory()) {
      return task(walk, filePath, pkg);
    }

    return { file: filePath, pkg };
  });

  return flatten(recurFiles);
}

function getStats(dir, file) {
  const filePath = path.join(dir, file);
  return stat(filePath).then((statFile) => ({ statFile, file, type: SUCCESS }));
}

module.exports = {
  checkDepends: run,
};
