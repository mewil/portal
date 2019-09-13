const fs = require('fs');

function addDependencyToPackageJsonFiles(depends, version) {
  const pkgJsonMap = {};

  depends.forEach(({ importName, pkg }) => {
    const jso = pkg.contents;
    if (!jso.dependencies) {
      jso.dependencies = {};
    }

    jso.dependencies[importName] = version;

    if (!pkgJsonMap[pkg.file]) {
      pkgJsonMap[pkg.file] = {};
    }

    pkgJsonMap[pkg.file] = jso;
  });

  return pkgJsonMap;
}

function removeDependencyFromPackageJsonFiles(depends) {
  const pkgJsonMap = {};

  depends.forEach(({ importName, pkg }) => {
    const jso = pkg.contents;

    delete jso.dependencies[importName];

    if (!pkgJsonMap[pkg.file]) {
      pkgJsonMap[pkg.file] = {};
    }

    pkgJsonMap[pkg.file] = jso;
  });

  return pkgJsonMap;
}

function savePackageJsonFiles(pkgJsonMap) {
  Object.keys(pkgJsonMap).forEach((pkgFile) => {
    const contents = JSON.stringify(pkgJsonMap[pkgFile], null, 2);
    fs.writeFile(pkgFile, contents, (err) => {
      if (err) {
        console.error(err);
        return;
      }

      console.log(`${pkgFile} dependencies have been updated`);
    });
  });
}

module.exports = {
  addDependencyToPackageJsonFiles,
  removeDependencyFromPackageJsonFiles,
  savePackageJsonFiles,
};
