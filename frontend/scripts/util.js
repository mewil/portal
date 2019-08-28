const fs = require('fs');

function task(generator, ...opts) {
  const iterator = generator(...opts);
  return new Promise((resolve) => {
    recursivelyNext();
    function recursivelyNext(...data) {
      const yielded = iterator.next(...data);
      if (yielded.done) {
        resolve(yielded.value);
        return;
      }
      if (Array.isArray(yielded.value)) {
        Promise.all(yielded.value)
          .then(recursivelyNext)
          .catch(console.log);

        return;
      }
      if (!isPromise(yielded.value)) {
        return;
      }
      yielded.value.then(recursivelyNext).catch(console.log);
    }
  });
}

function flatten(arr) {
  return [].concat(...arr);
}

function isPromise(val) {
  return val && typeof val.then === 'function';
}

function readdir(dir) {
  return new Promise((resolve, reject) => {
    fs.readdir(dir, (err, files) => {
      if (err) {
        return reject(err);
      }
      return resolve(files);
    });
  });
}

function stat(dir) {
  return new Promise((resolve, reject) => {
    fs.stat(dir, (err, stats) => {
      if (err) {
        return reject(new Error(`${err.path} not found`));
      }
      return resolve(stats);
    });
  });
}

function read(file) {
  return new Promise((resolve, reject) => {
    fs.readFile(file, (err, data) => {
      if (err) {
        return reject(err);
      }
      return resolve(data);
    });
  });
}

function write(file, data) {
  return new Promise((resolve, reject) => {
    fs.writeFile(file, data, (err) => {
      if (err) {
        return reject(err);
      }
      return resolve(`${file} was created!`);
    });
  });
}

function cleanGlob(dir) {
  return dir.replace('*', '');
}

function getArgs(args) {
  return args.reduce((output, arg) => {
    const argSplit = arg.split('=');
    const argName = argSplit[0].replace('--', '');
    const argValue = argSplit[1];
    return Object.assign({}, output, {
      [argName]: argValue,
    });
  }, {});
}

module.exports = {
  read,
  write,
  stat,
  readdir,
  task,
  cleanGlob,
  flatten,
  getArgs,
};
