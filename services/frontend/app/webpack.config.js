const webpack = require('webpack');
const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin');

const lifecycleEvent = process.env.npm_lifecycle_event;

const devConfig = {
  entry: ['babel-polyfill', './src/app/app.js'],
  output: {
    publicPath: '/',
    path: path.resolve('./build'),
    filename: 'js/app.js',
  },
  mode: 'development',
  devtool: 'source-map',
  resolve: {
    modules: ['web_modules', 'node_modules', 'app', 'static'],
    extensions: ['.js'],
  },
  module: {
    rules: [
      {
        test: /\.js?$/,
        enforce: 'pre',
        use: ['eslint-loader'],
        exclude: /node_modules/,
      },
      {
        test: /\.js?$/,
        exclude: /(node_modules|bower_components)/,
        use: ['babel-loader'],
      },
      {
        test: /\.(eot|ttf|woff|woff2|otf)$/,
        use: 'url-loader',
      },
      {
        test: /\.(png|jpg|jpeg|gif|woff|svg)$/,
        use: 'file-loader',
      },
    ],
  },
  plugins: [
    // inject styles and javascript into index.html
    new HtmlWebpackPlugin({
      title: 'Webpack Build',
      template: './src/app/index.html',
    }),
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': '"development"',
    }),
  ],
  devServer: {
    historyApiFallback: true,
    contentBase: './build',
    proxy: {
      '/api': {
        target: 'http://localhost:9090',
        xfwd: true,
        changeOrigin: true,
      },
    },
  },
  performance: {
    hints: process.env.NODE_ENV === 'production' ? 'warning' : false,
  },
};

const buildConfig = {
  entry: ['babel-polyfill', './src/app/app.js'],
  output: {
    publicPath: '/',
    path: path.resolve('./build'),
    filename: 'js/app.js',
  },
  mode: 'production',
  devtool: 'source-map',
  resolve: {
    extensions: ['.js'],
  },
  optimization: {
    minimize: true,
  },
  module: {
    rules: [
      {
        test: /\.js?$/,
        enforce: 'pre',
        use: ['eslint-loader'],
        exclude: /node_modules/,
      },
      {
        test: /\.js?$/,
        exclude: /(node_modules|bower_components)/,
        use: ['babel-loader'],
      },
      { test: /\.(png|jpg|jpeg|gif|woff|svg|otf)$/, use: 'file-loader' },
    ],
  },
  plugins: [
    new HtmlWebpackPlugin({
      title: 'Webpack Build',
      template: './src/app/index.html',
    }),
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': '"production"',
    }),
    new CleanWebpackPlugin([
      'build/fonts',
      'build/js',
      'build/styles',
      'build/index.html',
    ]),
    new CopyWebpackPlugin([
      { context: './app/favicon/', from: '**/*', to: './favicon/' },
      { context: './app/fonts/', from: '**/*', to: './fonts/' },
    ]),
    new webpack.HotModuleReplacementPlugin(),
    new webpack.NoEmitOnErrorsPlugin(),
  ],
  devServer: {
    historyApiFallback: true,
    contentBase: './build',
    proxy: {
      '/api': {
        target: 'http://localhost:9090',
        xfwd: true,
        changeOrigin: true,
      },
    },
  },
  performance: {
    hints: process.env.NODE_ENV === 'production' ? 'warning' : false,
  },
};

switch (lifecycleEvent) {
  case 'build':
    module.exports = buildConfig;
    break;
  default:
    module.exports = devConfig;
    break;
}
