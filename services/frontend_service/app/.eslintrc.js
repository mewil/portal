module.exports = {
  parser: 'babel-eslint',
  extends: [
    'airbnb',
    'plugin:import/errors',
    'plugin:import/warnings',
    'prettier',
  ],
  env: {
    browser: true,
    mocha: true,
    node: true,
  },
  rules: {
    quotes: ['error', 'single'],
    'react/jsx-no-bind': 0,
    'react/no-multi-comp': 0,
    'react/sort-comp': 0,
    'react/prop-types': 0,
    'react/forbid-prop-types': [2, { forbid: ['any', 'array'] }],
    camelcase: 0,
    'import/no-extraneous-dependencies': 0,
    'import/prefer-default-export': 0,
    'import/first': 0,
    'import/newline-after-import': 0,
    'global-require': 0,
    'require-yield': 0,
    'no-underscore-dangle': 0,
    'no-use-before-define': 0,
    'consistent-return': 0,
    'no-console': 0,
    'func-names': 0,
    'object-shorthand': 0,
    'guard-for-in': 0,
    'new-cap': 0,
    'no-continue': 0,
    'no-prototype-builtins': 0,
    'arrow-parens': ['error', 'always'],
    'no-plusplus': [2, { allowForLoopAfterthoughts: true }],
    'no-bitwise': 0,
    'prefer-arrow-callback': 1,
    'no-constant-condition': 0,
    'no-restricted-syntax': 0,
    'no-multi-assign': 0,
    'space-before-function-paren': 0,
  },
  plugins: ['react', 'import'],
  parserOptions: {
    ecmaFeatures: {
      jsx: false,
    },
    sourceType: 'module',
    allowImportExportEverywhere: true,
  },
  globals: {
    expect: true,
    genRunner: true,
    expectGen: true,
    jest: true,
    spyOn: true,
    beforeAll: true,
    afterAll: true,
  },
};
