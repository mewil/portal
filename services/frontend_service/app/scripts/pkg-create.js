/*
 * This script creates `package.json` and `index.js` files for new packages in
 * a new directory under `src/`.
 */
const path = require('path');
const fs = require('fs');

const pkgFolder = './src/';

run(process.argv.slice(2)[0]);

function run(name) {
  const pkgDir = path.join(pkgFolder, name);
  if (!fs.existsSync(pkgDir)) {
    fs.mkdirSync(pkgDir);
  }

  const pkgJson = path.join(pkgDir, 'package.json');
  fs.writeFile(pkgJson, createPkgJsonContent(name), (err) => {
    if (err) throw err;
  });

  const pkgIndex = path.join(pkgDir, 'index.js');
  fs.closeSync(fs.openSync(pkgIndex, 'w'));

  console.log(`Created ${name} package with package.json and index.js`);
}

function createPkgJsonContent(name) {
  const json = {
    name: `@portal/${name}`,
    version: '1.0.0',
    description: '',
    main: 'index.js',
    author: '',
    license: 'GPL-3.0',
    dependencies: {},
  };

  return JSON.stringify(json, null, 2);
}
