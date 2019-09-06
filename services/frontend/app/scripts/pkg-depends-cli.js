const { get } = require('lodash');

const opkg = require('../package.json');

const { task, error } = require('./util');
const { checkDepends } = require('./pkg-depends-check');
const {
  addDependencyToPackageJsonFiles,
  removeDependencyFromPackageJsonFiles,
  savePackageJsonFiles,
} = require('./pkg-depends-update');
const {
  UNNECESSARY_PACKAGE_DEPENDENCY,
  DEPENDENCY_MISSING_IN_PACKAGE_JSON,
} = require('./constants');

const VERSION = '1.0.0';
const { workspaces } = opkg;
const option = get(process.argv, '[2]', '');
const shouldFix = option === '--fix';

task(run, workspaces, shouldFix, VERSION);

function* run(dirs, fix, version) {
  console.log(
    'Checking that all packages have included all @portal dependencies and no unnecessary @portal dependencies inside their `package.json`',
  );
  const errors = yield task(checkDepends, dirs);

  if (errors.length === 0) {
    console.log(
      'All `package.json` files have the required dependencies and nothing else',
    );
    return;
  }

  const dependsErrors = errors.filter(
    (err) => err.type === DEPENDENCY_MISSING_IN_PACKAGE_JSON,
  );
  const unnecessaryErrors = errors.filter(
    (err) => err.type === UNNECESSARY_PACKAGE_DEPENDENCY,
  );

  if (!fix) {
    errors.forEach((err) => console.log(err));

    error(`Detected (${errors.length}) errors total`);

    if (dependsErrors.length > 0) {
      error(
        `Detected (${dependsErrors.length}) dependency issues relating to dependencies missing in package.json`,
      );
    }

    if (unnecessaryErrors.length > 0) {
      const pkgFiles = {};
      unnecessaryErrors.forEach((err) => {
        const key = err.pkg.file;
        if (!pkgFiles.hasOwnProperty(key)) pkgFiles[key] = 0;
        pkgFiles[key] += 1;
      });

      Object.keys(pkgFiles).forEach((pkgFile) =>
        error(`${pkgFile} has (${pkgFiles[pkgFile]}) dependency issues`),
      );
    }

    console.log('Run `yarn pkg-update` to automatically fix dependency issues');
    process.exit(1);
  }

  const pkgJsonMap = {
    ...addDependencyToPackageJsonFiles(dependsErrors, version),
    ...removeDependencyFromPackageJsonFiles(unnecessaryErrors),
  };

  savePackageJsonFiles(pkgJsonMap);
}
